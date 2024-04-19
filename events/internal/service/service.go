package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
)

var (
	InternalServerError = errors.New("internal server error")
	BadRequest          = errors.New("bad request")
	NotFound            = errors.New("not found")
)

type Storage interface {
	Event(ctx context.Context, id string) (domain.Event, error)
	Events(ctx context.Context, isFinished bool) ([]domain.Event, error)
	AddEvent(ctx context.Context, event domain.Event) (string, error)
	UpdateEvent(ctx context.Context, id string, event domain.Event) error
	DeleteEvent(ctx context.Context, id string) error
}

type Service struct {
	log     *slog.Logger
	storage Storage
}

func New(log *slog.Logger, storage Storage) *Service {
	log = log.With(slog.String("service", "events"))
	return &Service{
		log:     log,
		storage: storage,
	}
}

func (s *Service) Event(ctx context.Context, id string) (domain.Event, error) {
	const op = "service.Event"

	log := s.log.With("operation", op)
	log.Info("Geting event")

	event, err := s.storage.Event(ctx, id)
	if err != nil {
		s.log.Error("Error getting event", err)
		return domain.Event{}, fmt.Errorf("%w: %s", NotFound, "error getting event")
	}
	log.Info("Got event")
	return event, nil
}

func (s *Service) Events(ctx context.Context, isFinished bool) ([]domain.Event, error) {
	const op = "service.Events"
	log := s.log.With("operation", op)
	log.Info("Geting events")
	events, err := s.storage.Events(ctx, isFinished)
	if err != nil {
		s.log.Error("Error getting events", err)
		return nil, fmt.Errorf("%w: %s", InternalServerError, "error getting events")
	}
	log.Info("Got events")
	return events, nil
}

func (s *Service) AddEvent(ctx context.Context, event domain.Event) (string, error) {
	const op = "service.AddEvent"

	log := s.log.With("operation", op)
	log.Info("Adding event")
	id, err := s.storage.AddEvent(ctx, event)
	if err != nil {
		s.log.Error("Error adding event", err)
		return "", fmt.Errorf("%w: %s", InternalServerError, "error adding event")
	}
	log.Info("Added event")
	return id, nil
}

func (s *Service) UpdateEvent(ctx context.Context, id string, event domain.Event) error {
	const op = "service.UpdateEvent"

	log := s.log.With("operation", op)
	log.Info("Updating event")
	err := s.storage.UpdateEvent(ctx, id, event)
	if err != nil {
		s.log.Error("Error updating event", err)
		return fmt.Errorf("%w: %s", NotFound, "error updating event")
	}
	log.Info("Updated event")
	return nil
}

func (s *Service) DeleteEvent(ctx context.Context, id string) error {
	const op = "service.DeleteEvent"

	log := s.log.With("operation", op)
	log.Info("Deleting event")
	err := s.storage.DeleteEvent(ctx, id)
	if err != nil {
		s.log.Error("Error deleting event", err)
		return fmt.Errorf("%w: %s", NotFound, "error deleting event")
	}
	log.Info("Deleted event")
	return nil
}
