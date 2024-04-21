package service

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/models"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/requests"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/responses"
	eventsclient "github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/lib/events_client"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
	ErrNotFound       = errors.New("not found")
)

type Storage interface {
	CreateProduct(ctx context.Context, product models.Product) error
	Product(ctx context.Context, id string) (models.Product, error)
	Products(ctx context.Context) ([]models.Product, error)
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

func (s *Service) CreateProduct(ctx context.Context, product requests.CreateProduct) error {
	const op = "service.CreateProduct"
	log := s.log.With("op", op)
	log.Info("create product")

	log.Debug("save image")
	imageName, err := saveImage(product.Image)
	if err != nil {
		log.Debug("error save image", "err", err)
		return fmt.Errorf("%s: %w", op, ErrInternalServer)
	}

	log.Debug("save product")
	err = s.storage.CreateProduct(ctx, models.Product{
		ID:    primitive.NewObjectID(),
		Name:  product.Name,
		Image: imageName,
		Price: product.Price,
	})
	if err != nil {
		log.Debug("error save product", "err", err)
		return fmt.Errorf("service.CreateProduct: %w", ErrInternalServer)
	}
	return nil
}

func (s *Service) Product(ctx context.Context, id string) (responses.Product, error) {
	const op = "service.Product"
	log := s.log.With("op", op)
	log.Info("get product")

	product, err := s.storage.Product(ctx, id)
	if err != nil {
		log.Debug("error get product", "err", err)
		return responses.Product{}, fmt.Errorf("%s: %w", op, ErrNotFound)
	}
	return responses.Product{
		ID:    product.ID.Hex(),
		Name:  product.Name,
		Image: product.Image,
		Price: product.Price,
	}, nil
}

func (s *Service) Products(ctx context.Context) ([]responses.Product, error) {
	const op = "service.Products"
	log := s.log.With("op", op)
	log.Info("get products")

	products, err := s.storage.Products(ctx)
	if err != nil {
		log.Debug("error get products", "err", err)
		return nil, fmt.Errorf("%s: %w", op, ErrInternalServer)
	}
	respProducts := make([]responses.Product, 0, len(products))
	for _, product := range products {
		respProducts = append(respProducts, responses.Product{
			ID:    product.ID.Hex(),
			Name:  product.Name,
			Image: product.Image,
			Price: product.Price,
		})
	}
	return respProducts, nil
}

func (s *Service) Buy(ctx context.Context, userId, productId string) error {
	const op = "service.Buy"
	log := s.log.With("op", op)
	log.Info("buy product")

	product, err := s.storage.Product(ctx, productId)
	if err != nil {
		log.Debug("error get product", "err", err)
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}
	rating, err := eventsclient.GetUserRating(ctx, userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rating < product.Price {
		return fmt.Errorf("%s: %w", op, ErrBadRequest)
	}

	err = eventsclient.DecreaseUserRating(ctx, userId, product.Price)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func saveImage(imageBase64 string) (string, error) {
	decodedBytes, err := bytesFromBase64(imageBase64)
	if err != nil {
		return "", err
	}
	format := getFormat(decodedBytes[:4])
	name := generateImageName(format)

	// save image
	f, err := os.Create("./media/" + name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.Write(decodedBytes)
	if err != nil {
		return "", err
	}

	return name, nil
}

func bytesFromBase64(imageBase64 string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return nil, err
	}
	return decodedBytes, nil
}

func generateImageName(format string) string {
	prefix := time.Now().Unix()
	return fmt.Sprintf("%d_%d.%s", prefix, rand.Int(), strings.ToLower(format))
}

func getFormat(data []byte) string {
	if data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 {
		return "png"
	}
	if data[0] == 0xFF && data[1] == 0xD8 {
		return "jpg"
	}
	if data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x38 {
		return "gif"
	}
	if data[0] == 0x42 && data[1] == 0x4D {
		return "bmp"
	}
	return ""
}
