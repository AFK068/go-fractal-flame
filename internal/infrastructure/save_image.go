package infrastructure

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func SaveImage(img *domain.FractalImage, filename string) error {
	outImg := image.NewRGBA(image.Rect(0, 0, img.Width, img.Height))

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			pixel := img.Pixel(x, y)

			col := color.RGBA{
				R: uint8(pixel.RGB.R),
				G: uint8(pixel.RGB.G),
				B: uint8(pixel.RGB.B),
				A: uint8(255),
			}

			outImg.Set(x, y, col)
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	return png.Encode(file, outImg)
}
