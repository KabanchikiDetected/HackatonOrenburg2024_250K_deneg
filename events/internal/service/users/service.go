package users

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
	customErrors "github.com/KabanchikiDetected/hackaton/events/internal/errors"
)

type Storage interface {
	AddEventToUser(ctx context.Context, id string, eventID string) error
	UserEvents(ctx context.Context, id string) (domain.EventsToStudent, error)
	DicrementRating(ctx context.Context, id string, rating int) error
	GetAllUserRatings(ctx context.Context) ([]domain.UserRating, error)
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
		if errors.Is(err, customErrors.NotFound) {
			return fmt.Errorf("%w: %s", customErrors.NotFound, "error adding event to user")
		}
		if errors.Is(err, customErrors.BadRequest) {
			return fmt.Errorf("%w: %s", customErrors.BadRequest, "error adding event to user")
		}
		return fmt.Errorf("%w: %s", customErrors.InternalServerError, "error adding event to user")
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
		return domain.EventsToStudent{}, fmt.Errorf("%w: %s", customErrors.NotFound, "error getting user events")
	}
	log.Info("Got user events")
	return eventsToStudent, nil
}

func (s *Service) DicrementRating(ctx context.Context, id string, rating int) error {
	const op = "service.DicrementRating"

	log := s.log.With("operation", op)
	log.Info("Dicrementing rating")
	err := s.storage.DicrementRating(ctx, id, rating)
	if err != nil {
		if errors.Is(err, customErrors.NotFound) {
			return fmt.Errorf("%w: %s", customErrors.NotFound, "error dicrementing rating")
		} else if errors.Is(err, customErrors.BadRequest) {
			return fmt.Errorf("%w: %s", customErrors.BadRequest, "error dicrementing rating")
		} else {
			return fmt.Errorf("%w: %s", customErrors.InternalServerError, "error dicrementing rating")
		}
	}
	log.Info("Dicrementing rating")
	return nil
}

func (s *Service) GetAllUserRatings(ctx context.Context) ([]domain.UserRating, error) {
	const op = "service.GetAllUserRatings"

	log := s.log.With("operation", op)
	log.Info("Getting all user ratings")
	userRatings, err := s.storage.GetAllUserRatings(ctx)
	if err != nil {
		s.log.Error("Error getting all user ratings", err)
		return nil, fmt.Errorf("%w: %s", customErrors.NotFound, "error getting all user ratings")
	}
	log.Info("Got all user ratings")
	return userRatings, nil
}
