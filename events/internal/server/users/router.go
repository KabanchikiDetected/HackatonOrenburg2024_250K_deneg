package users

import (
	"context"
	"net/http"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
	"github.com/KabanchikiDetected/hackaton/events/internal/server/utils"
	auth "github.com/KabanchikiDetected/hackaton/events/pkg/users"
	"github.com/go-pkgz/routegroup"
)

type UsersService interface {
	AddEventToUser(ctx context.Context, id string, eventID string) error
	UserEvents(ctx context.Context, id string) (domain.EventsToStudent, error)
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

func (h *UserRouter) userEvents(w http.ResponseWriter, r *http.Request) {
	payload, err := auth.FromContext(r.Context())
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	events, err := h.service.UserEvents(r.Context(), payload.ID)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = utils.Encode(w, r, events)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserRouter) addEventToUser(w http.ResponseWriter, r *http.Request) {
	payload, err := auth.FromContext(r.Context())

	if err != nil {
		utils.HandleError(err, w)
		return
	}

	eventID := r.PathValue("event_id")

	err = h.service.AddEventToUser(r.Context(), payload.ID, eventID)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	userEvents, err := h.service.UserEvents(r.Context(), payload.ID)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	err = utils.Encode(w, r, userEvents)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (r *UserRouter) init() {
	key := utils.GetKey()
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("GET /my/events", r.userEvents)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("POST /my/events/{event_id}", r.addEventToUser)
}
