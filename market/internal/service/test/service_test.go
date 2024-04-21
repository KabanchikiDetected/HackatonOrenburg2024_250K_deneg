package service_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log/slog"
	"os"
	"testing"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/models"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/requests"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/service"
)

type Storage struct {
}

func (m *Storage) CreateProduct(ctx context.Context, product models.Product) error {
	return nil
}

func (m *Storage) Products(ctx context.Context) ([]models.Product, error) {
	return nil, nil
}

func (m *Storage) Product(ctx context.Context, id string) (models.Product, error) {
	return models.Product{}, nil
}

func TestService_CreateProductImageSaving(t *testing.T) {

	testCases := []struct {
		desc        string
		imageFormat string
		color       color.Color
	}{
		{
			desc: "png_image",
			color: color.RGBA{
				R: 255,
				G: 0,
				B: 0,
				A: 255,
			},
			imageFormat: "png",
		},
		{
			desc: "jpeg_image",
			color: color.RGBA{
				R: 0,
				G: 255,
				B: 0,
				A: 255,
			},
			imageFormat: "jpeg",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			storage := Storage{}
			log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
			service := service.New(log, &storage)
			err := service.CreateProduct(context.Background(), requests.CreateProduct{
				Image: base64EncodeImage(tC.imageFormat, generateImage(tC.color)),
			})
			if err != nil {
				dir, _ := os.Getwd()
				fmt.Println("Current dir: ", dir)
				t.Fatal(err)
			}
		})
	}
}

func generateImage(fillColor color.Color) image.Image {

	// colorR, colorG, colorB, _ :=
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, fillColor)
		}
	}
	return img
}

func base64EncodeImage(format string, img image.Image) string {

	buf := new(bytes.Buffer)
	switch format {
	case "png":
		err := png.Encode(buf, img)
		if err != nil {
			return ""
		}
	case "jpeg":
		err := jpeg.Encode(buf, img, nil)
		if err != nil {
			return ""
		}
	default:
		return ""
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
