package events

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
	"github.com/KabanchikiDetected/hackaton/events/internal/server/schemas"
	"github.com/KabanchikiDetected/hackaton/events/internal/server/utils"
	auth "github.com/KabanchikiDetected/hackaton/events/pkg/users"
	"github.com/go-pkgz/routegroup"
)

type EventService interface {
	Event(ctx context.Context, id string) (domain.Event, error)
	Events(ctx context.Context, isFinished bool) ([]domain.Event, error)
	AddEvent(ctx context.Context, event domain.Event) (string, error)
	UpdateEvent(ctx context.Context, id string, event domain.Event) error
	DeleteEvent(ctx context.Context, id string) error
	InsertImage(ctx context.Context, id string, image string) error
	SearchByTitle(ctx context.Context, title string) ([]domain.Event, error)
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

	title := r.URL.Query().Get("title")
	if title != "" {
		events, err := h.service.SearchByTitle(r.Context(), title)
		if err != nil {
			utils.HandleError(err, w)
			return
		}
		err = utils.Encode(w, r, events)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

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
	var event schemas.EventSchema
	id := utils.GetIdFromPath(w, r)
	err := utils.Decode(w, r, &event)

	if err != nil {
		return
	}

	eventDomain, err := event.ToDomain()
	eventDomain.ID = id

	if err != nil {
		utils.SendErrorMessage(w, err.Error())
		return
	}

	err = h.service.UpdateEvent(r.Context(), id, eventDomain)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	err = utils.Encode(w, r, eventDomain)
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
	utils.SendResponceMessage(w, "Event deleted")
	w.WriteHeader(http.StatusOK)
}

func (h *EventRouter) insertImage(w http.ResponseWriter, r *http.Request) {
	id := utils.GetIdFromPath(w, r)

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
	err = h.service.InsertImage(r.Context(), id, filePath)
	if err != nil {
		utils.HandleError(err, w)
		return
	}
	utils.SendResponceMessage(w, "Image inserted")
	w.WriteHeader(http.StatusOK)
}

func (r *EventRouter) init() {
	key := utils.GetKey()
	r.mux.HandleFunc("GET /", r.events)
	r.mux.HandleFunc("GET /{id}", r.event)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("POST /", r.addEvent)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("PUT /{id}", r.updateEvent)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("DELETE /{id}", r.deleteEvent)
	r.mux.With(auth.MiddlwareJWT(key)).HandleFunc("POST /{id}/image", r.insertImage)
}
