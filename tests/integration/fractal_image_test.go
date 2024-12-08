package integration_test

import (
	"math"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestFractalImage_GammaCorrection(t *testing.T) {
	// Create a new image.
	width, height := 100, 100

	img := domain.NewFractalImage(width, height)
	tempImage := domain.NewFractalImage(width, height)

	// Fill the image with test data.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := img.Pixel(x, y)
			tempPixel := tempImage.Pixel(x, y)

			pixel.HitCount = (uint64(x) + uint64(y)) % 10 //nolint:gosec // unreal overflow
			tempPixel.HitCount = pixel.HitCount

			pixel.RGB = domain.RGB{
				R: x % 256,
				G: y % 256,
				B: (x + y) % 256,
			}

			tempPixel.RGB = pixel.RGB
		}
	}

	// Calculate the maximum normal.
	var maxNormal float64

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := tempImage.Pixel(x, y)
			if pixel.HitCount != 0 {
				pixel.Normal = math.Log10(float64(pixel.HitCount))
				if pixel.Normal > maxNormal {
					maxNormal = pixel.Normal
				}
			}
		}
	}

	// Apply gamma correction.
	gamma := 2.2
	img.GammaCorrection(gamma)

	// Check that the gamma correction was applied correctly.
	for row := 0; row < width; row++ {
		for col := 0; col < height; col++ {
			pixelOrigin := img.Pixel(row, col)
			pixelTemp := tempImage.Pixel(row, col)

			if pixelTemp.HitCount != 0 {
				pixelTemp.Normal /= maxNormal

				expectedR := int(float64(pixelTemp.RGB.R) * math.Pow(pixelTemp.Normal, 1.0/gamma))
				expectedG := int(float64(pixelTemp.RGB.G) * math.Pow(pixelTemp.Normal, 1.0/gamma))
				expectedB := int(float64(pixelTemp.RGB.B) * math.Pow(pixelTemp.Normal, 1.0/gamma))

				if (expectedR != pixelOrigin.RGB.R) || (expectedG != pixelOrigin.RGB.G) || (expectedB != pixelOrigin.RGB.B) {
					require.Fail(t, "Gamma correction failed", "Expected R: %d, G: %d, B: %d, but got R: %d, G: %d, B: %d at pixel (%d, %d)",
						expectedR, expectedG, expectedB, pixelOrigin.RGB.R, pixelOrigin.RGB.G, pixelOrigin.RGB.B, row, col)
				}
			}
		}
	}
}
