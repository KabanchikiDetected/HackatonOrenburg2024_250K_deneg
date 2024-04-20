package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/config"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/domain/models"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/domain/requests"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
	ErrNotFound       = errors.New("not found")
)

type Storage interface {
	SaveUser(ctx context.Context, user models.User) error
	User(ctx context.Context, id string) (models.User, error)
	UserByEmail(ctx context.Context, email string) (models.User, error)
	MakeDeputy(ctx context.Context, id string) error
}

type Service struct {
	log     *slog.Logger
	storage Storage
}

func New(log *slog.Logger, storage Storage) *Service {
	return &Service{
		log:     log,
		storage: storage,
	}
}

func (s *Service) Register(ctx context.Context, user requests.Register) error {
	const op = "service.Register"
	log := s.log.With("operation", op)

	log.Info("register user")
	log.Debug("user", slog.Any("user", user))
	if user.Password != user.RepeatPassword {
		log.Debug("passwords do not match", "password", user.Password, "repeat_password", user.RepeatPassword)
		return fmt.Errorf("%w: %s", ErrBadRequest, "passwords do not match")
	}

	log.Debug("create user")

	// Hash password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		log.Debug("failed to hash password", "err", err.Error())
		return fmt.Errorf("%w: %s", ErrInternalServer, "failed to hash password")
	}
	log.Debug("saving user")
	err = s.storage.SaveUser(ctx, models.User{
		Role:     models.BaseUser,
		Email:    user.Email,
		Password: hashedPassword,
	})
	if err != nil {
		log.Debug("failed to save user", "err", err.Error())
		return fmt.Errorf("%w: %s", ErrInternalServer, "failed to save user")
	}
	log.Debug("user saved")
	return nil
}

func (s *Service) Login(ctx context.Context, user requests.Login) (string, error) {
	const op = "service.Login"
	log := s.log.With("operation", op)

	log.Info("login user")

	log.Debug("get user")
	userDB, err := s.storage.UserByEmail(ctx, user.Email)
	if err != nil {
		log.Debug(err.Error())
		return "", fmt.Errorf("%w: %s", ErrNotFound, "user not found")
	}

	log.Debug("check password")
	if !checkPasswordHash(user.Password, userDB.Password) {
		log.Debug("invalid password")
		return "", fmt.Errorf("%w: %s", ErrBadRequest, "invalid password")
	}
	token := createToken(userDB)
	return token, nil
}

func (s *Service) MakeDeputy(ctx context.Context, id string) (newToken string, err error) {
	const op = "service.MakeDeputy"
	log := s.log.With("operation", op)

	log.Info("make user as deputy")

	err = s.storage.MakeDeputy(ctx, id)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrNotFound, "failed to make user as deputy")
	}

	user, err := s.storage.User(ctx, id)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrNotFound, "failed to get user")
	}
	newToken = createToken(user)
	return newToken, nil
}

func createToken(user models.User) string {
	// Calculate expiration time
	exp := time.Now().Add(config.Config().JWT.TokenTTl).Unix()

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  exp,
	})

	// Sign and get the complete encoded token as string.
	signedToken, err := token.SignedString(config.Config().JWT.PrivateKey)
	if err != nil {
		return ""
	}

	// return signed token
	return signedToken

}

func CreateToken() {
	token := createToken(models.User{
		ID:   primitive.NewObjectID(),
		Role: models.BaseUser,
	})
	fmt.Println(token)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
