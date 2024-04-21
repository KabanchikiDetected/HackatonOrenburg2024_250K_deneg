package users

import (
	"context"
	"net/http"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
	"github.com/KabanchikiDetected/hackaton/events/internal/server/schemas"
	"github.com/KabanchikiDetected/hackaton/events/internal/server/utils"
	auth "github.com/KabanchikiDetected/hackaton/events/pkg/users"
	"github.com/go-pkgz/routegroup"
)

type UsersService interface {
	AddEventToUser(ctx context.Context, id string, eventID string) error
	UserEvents(ctx context.Context, id string) (domain.EventsToStudent, error)
	DicrementRating(ctx context.Context, id string, rating int) error
	GetAllUserRatings(ctx context.Context) ([]domain.UserRating, error)
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

func (h *UserRouter) userEventsByID(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIdFromPath(w, r)
	events, err := h.service.UserEvents(r.Context(), id)
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

func (h *UserRouter) dicrementRating(w http.ResponseWriter, r *http.Request) {
	payload, err := auth.FromContext(r.Context())

	if err != nil {
		utils.HandleError(err, w)
		return
	}

	var rating schemas.RatingSchema
	err = utils.Decode(w, r, &rating)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = h.service.DicrementRating(r.Context(), payload.ID, rating.Rating)
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

func (h *UserRouter) getAllUserRatings(w http.ResponseWriter, r *http.Request) {
	ratings, err := h.service.GetAllUserRatings(r.Context())
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = utils.Encode(w, r, ratings)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (r *UserRouter) init() {
	key := utils.GetKey()
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("GET /my/events", r.userEvents)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("POST /my/rating", r.dicrementRating)
	r.mux.HandleFunc("GET /{id}/events", r.userEventsByID)
	r.mux.HandleFunc("GET /rating", r.getAllUserRatings)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("POST /my/events/{event_id}", r.addEventToUser)
}
