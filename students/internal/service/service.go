package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/KabanchikiDetected/hackaton/students/internal/domain"
	"github.com/KabanchikiDetected/hackaton/students/internal/errors"
)

type Storage interface {
	User(ctx context.Context, id string) (domain.User, error)
	Users(ctx context.Context) ([]domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
	UpdateUser(ctx context.Context, id string, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
	InsertImage(ctx context.Context, id string, image string) error
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

func (s *Service) User(ctx context.Context, id string) (domain.User, error) {
	const op = "service.User"

	log := s.log.With("operation", op)
	log.Info("Getting User")

	user, err := s.storage.User(ctx, id)
	if err != nil {
		s.log.Error("Error getting user", err)
		return domain.User{}, fmt.Errorf("%w: %s", errors.NotFound, "error getting user")
	}
	log.Info("Got User")
	return user, nil
}

func (s *Service) Users(ctx context.Context) ([]domain.User, error) {
	const op = "service.Users"

	log := s.log.With("operation", op)
	log.Info("Getting Users")

	users, err := s.storage.Users(ctx)
	if err != nil {
		s.log.Error("Error getting users", err)
		return nil, fmt.Errorf("%w: %s", errors.NotFound, "error getting users")
	}
	log.Info("Got Users")
	return users, nil
}

func (s *Service) CreateUser(ctx context.Context, user domain.User) error {
	const op = "service.CreateUser"

	log := s.log.With("operation", op)
	log.Info("Creating User")

	err := s.storage.CreateUser(ctx, user)
	if err != nil {
		s.log.Error("Error creating user", err)
		return fmt.Errorf("%w: %s", errors.NotFound, "error creating user")
	}
	log.Info("Created User")
	return nil
}

func (s *Service) UpdateUser(ctx context.Context, id string, user domain.User) error {
	const op = "service.UpdateUser"

	log := s.log.With("operation", op)
	log.Info("Updating User")

	if err := s.storage.UpdateUser(ctx, id, user); err != nil {
		s.log.Error("Error updating user", err)
		return fmt.Errorf("%w: %s", errors.NotFound, "error updating user")
	}
	log.Info("Updated User")
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	const op = "service.DeleteUser"

	log := s.log.With("operation", op)
	log.Info("Deleting User")

	if err := s.storage.DeleteUser(ctx, id); err != nil {
		s.log.Error("Error deleting user", err)
		return fmt.Errorf("%w: %s", errors.NotFound, "error deleting user")
	}
	log.Info("Deleted User")
	return nil
}

func (s *Service) InsertImage(ctx context.Context, id string, image string) error {
	const op = "service.InsertImage"

	log := s.log.With("operation", op)
	log.Info("Inserting image")

	if err := s.storage.InsertImage(ctx, id, image); err != nil {
		s.log.Error("Error inserting image", err)
		return fmt.Errorf("%w: %s", errors.NotFound, "error inserting image")
	}
	log.Info("Inserted image")
	return nil
}
