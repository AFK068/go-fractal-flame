package infrastructure_test

import (
	"image"
	"image/color"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
)

func TestSaveImage(t *testing.T) {
	width, height := 100, 100
	img := domain.NewFractalImage(width, height)

	// Fill the image with test data.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := img.Pixel(x, y)
			pixel.RGB = domain.RGB{
				R: x % 256,
				G: y % 256,
				B: (x + y) % 256,
			}
		}
	}

	// Save the image.
	filename := "test_fractal.png"
	err := infrastructure.SaveImage(img, filename)
	assert.NoError(t, err, "SaveImage should not return an error")

	// Check that the file was created.
	_, err = os.Stat(filename)
	assert.NoError(t, err, "File should be created")

	// Open the file to check its contents.
	file, err := os.Open(filename)
	assert.NoError(t, err, "Failed to open file")
	defer file.Close()

	// Decode the image.
	decodedImg, _, err := image.Decode(file)
	assert.NoError(t, err, "Failed to decode image")

	// Check the contents of the image.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			expectedColor := color.RGBA{
				R: uint8(x % 256),       //nolint:gosec // unreal overflow, positive mod 256
				G: uint8(y % 256),       //nolint:gosec // unreal overflow, positive mod 256
				B: uint8((x + y) % 256), //nolint:gosec // unreal overflow, positive mod 256
				A: 255,
			}

			assert.Equal(t, expectedColor, decodedImg.At(x, y), "Pixel color mismatch at (%d, %d)", x, y)
		}
	}

	// Remove the test file.
	err = os.Remove(filename)
	assert.NoError(t, err, "Failed to remove test file")
}
