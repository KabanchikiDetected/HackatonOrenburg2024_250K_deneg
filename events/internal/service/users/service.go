package users

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
	"github.com/KabanchikiDetected/hackaton/events/internal/errors"
)

type Storage interface {
	AddEventToUser(ctx context.Context, id string, eventID string) error
	UserEvents(ctx context.Context, id string) (domain.EventsToStudent, error)
}

type Service struct {
	log     *slog.Logger
	storage Storage
}

func New(log *slog.Logger, storage Storage) *Service {
	log = log.With(slog.String("service", "users"))
	return &Service{
		log:     log,
		storage: storage,
	}
}

func (s *Service) AddEventToUser(ctx context.Context, id string, eventID string) error {
	const op = "service.AddEventToUser"

	log := s.log.With("operation", op)
	log.Info("Adding event to user")
	err := s.storage.AddEventToUser(ctx, id, eventID)
	if err != nil {
		s.log.Error("Error adding event to user", err)
		return fmt.Errorf("%w: %s", errors.NotFound, "error adding event to user")
	}
	log.Info("Added event to user")
	return nil
}

func (s *Service) UserEvents(ctx context.Context, id string) (domain.EventsToStudent, error) {
	const op = "service.UserEvents"

	log := s.log.With("operation", op)
	log.Info("Getting user events")
	eventsToStudent, err := s.storage.UserEvents(ctx, id)
	if err != nil {
		s.log.Error("Error getting user events", err)
		return domain.EventsToStudent{}, fmt.Errorf("%w: %s", errors.NotFound, "error getting user events")
	}
	log.Info("Got user events")
	return eventsToStudent, nil
}
