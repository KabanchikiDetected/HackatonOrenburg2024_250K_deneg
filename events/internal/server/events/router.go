package events

import (
	"context"
	"net/http"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
	"github.com/KabanchikiDetected/hackaton/events/internal/server/schemas"
	"github.com/KabanchikiDetected/hackaton/events/internal/server/utils"
	"github.com/go-pkgz/routegroup"
)

type EventService interface {
	Event(ctx context.Context, id string) (domain.Event, error)
	Events(ctx context.Context, isFinished bool) ([]domain.Event, error)
	AddEvent(ctx context.Context, event domain.Event) (string, error)
	UpdateEvent(ctx context.Context, event domain.Event) error
	DeleteEvent(ctx context.Context, id string) error
}

type EventRouter struct {
	mux     *routegroup.Bundle
	service EventService
}

func Register(service EventService, mux *routegroup.Bundle) *EventRouter {
	router := &EventRouter{
		mux:     mux,
		service: service,
	}
	router.init()
	return router
}

func (h *EventRouter) event(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIdFromPath(w, r)

	event, err := h.service.Event(r.Context(), id)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = utils.Encode(w, r, event)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *EventRouter) events(w http.ResponseWriter, r *http.Request) {
	isFinished := utils.GetBoolQuery(r.URL.Query().Get("is_finished"))

	events, err := h.service.Events(r.Context(), isFinished)
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

func (h *EventRouter) addEvent(w http.ResponseWriter, r *http.Request) {
	var event schemas.EventSchema
	err := utils.Decode(w, r, &event)
	if err != nil {
		return
	}

	domainEvent, err := event.ToDomain()
	if err != nil {
		utils.SendErrorMessage(w, err.Error())
		return
	}

	id, err := h.service.AddEvent(r.Context(), domainEvent)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	domainEvent.ID = id
	err = utils.Encode(w, r, domainEvent)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *EventRouter) updateEvent(w http.ResponseWriter, r *http.Request) {
	var event domain.Event
	err := utils.Decode(w, r, &event)
	if err != nil {
		return
	}
	err = h.service.UpdateEvent(r.Context(), event)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = utils.Encode(w, r, event)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *EventRouter) deleteEvent(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIdFromPath(w, r)
	err := h.service.DeleteEvent(r.Context(), id)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (r *EventRouter) init() {
	r.mux.HandleFunc("GET /", r.events)
	r.mux.HandleFunc("GET /{id}", r.event)
	r.mux.HandleFunc("POST /", r.addEvent)
	r.mux.HandleFunc("PUT /{id}", r.updateEvent)
	r.mux.HandleFunc("DELETE /{id}", r.deleteEvent)
}
