package users

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/KabanchikiDetected/hackaton/students/internal/domain"
	"github.com/KabanchikiDetected/hackaton/students/internal/errors"
	"github.com/KabanchikiDetected/hackaton/students/internal/server/schemas"
	"github.com/KabanchikiDetected/hackaton/students/internal/server/utils"
	auth "github.com/KabanchikiDetected/hackaton/students/pkg/users"
	"github.com/go-pkgz/routegroup"
)

type UsersService interface {
	Users(ctx context.Context) ([]domain.User, error)
	User(ctx context.Context, id string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
	UpdateUser(ctx context.Context, id string, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
	InsertImage(ctx context.Context, id string, image string) error
}

type UserRouter struct {
	mux     *routegroup.Bundle
	service UsersService
}

func Register(service UsersService, mux *routegroup.Bundle) *UserRouter {
	router := &UserRouter{
		mux:     mux,
		service: service,
	}
	router.init()
	return router
}

func (h *UserRouter) user(w http.ResponseWriter, r *http.Request) {
	payload, err := auth.FromContext(r.Context())
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	user, err := h.service.User(r.Context(), payload.ID)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	err = utils.Encode(w, r, user)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserRouter) userByID(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIdFromPath(w, r)
	user, err := h.service.User(r.Context(), id)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = utils.Encode(w, r, user)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserRouter) users(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.Users(r.Context())
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	err = utils.Encode(w, r, users)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserRouter) createUser(w http.ResponseWriter, r *http.Request) {
	payload, err := auth.FromContext(r.Context())

	if err != nil {
		utils.HandleError(err, w)
		return
	}

	var user schemas.UserSchema
	err = utils.Decode(w, r, &user)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	userDomain, err := user.ToDomain()
	userDomain.ID = payload.ID
	userDomain.Role = domain.Role(payload.Role)

	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = h.service.CreateUser(r.Context(), userDomain)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = utils.Encode(w, r, userDomain)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserRouter) updateUser(w http.ResponseWriter, r *http.Request) {
	payload, err := auth.FromContext(r.Context())
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	var user schemas.UserSchema
	err = utils.Decode(w, r, &user)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	userDomain, err := user.ToDomain()
	userDomain.Role = domain.Role(payload.Role)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = h.service.UpdateUser(r.Context(), payload.ID, userDomain)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	userDomain.ID = payload.ID

	err = utils.Encode(w, r, userDomain)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserRouter) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIdFromPath(w, r)

	payload, err := auth.FromContext(r.Context())
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	if strings.ToLower(payload.Role) != "deputy" {
		utils.HandleError(errors.Forbidden, w)
		return
	}

	err = h.service.DeleteUser(r.Context(), id)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	utils.SendResponceMessage(w, "user deleted")

	w.WriteHeader(http.StatusOK)
}

func (h *UserRouter) insertImage(w http.ResponseWriter, r *http.Request) {
	payload, err := auth.FromContext(r.Context())

	if err != nil {
		utils.HandleError(err, w)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		utils.SendErrorMessage(w, err.Error())
		return
	}
	defer file.Close()

	var mediaDir string = "media"

	if _, err := os.Stat(mediaDir); os.IsNotExist(err) {
		err = os.Mkdir(mediaDir, 0777)
		if err != nil {
			fmt.Println(err)
		}
	}
	filePath := fmt.Sprintf("%s/%s", mediaDir, header.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println(err)
	}
	err = h.service.InsertImage(r.Context(), payload.ID, filePath)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	utils.SendResponceMessage(w, "Image inserted")
	w.WriteHeader(http.StatusOK)
}

func (r *UserRouter) init() {
	key := utils.GetKey()
	r.mux.HandleFunc("GET /{id}", r.userByID)
	r.mux.HandleFunc("GET /", r.users)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("GET /me", r.user)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("POST /", r.createUser)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("POST /me/image", r.insertImage)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("PUT /me", r.updateUser)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("DELETE /{id}", r.deleteUser)
}
