package domain

import "math"

type FractalImage struct {
	Data          []Pixel
	Width, Height int
}

func NewFractalImage(width, height int) *FractalImage {
	data := make([]Pixel, width*height)

	for i := range data {
		data[i] = Pixel{
			X: i % width,
			Y: i / width,
		}
	}

	return &FractalImage{data, width, height}
}

func (fi *FractalImage) Contains(x, y int) bool {
	return x > 0 && x < fi.Width && y > 0 && y < fi.Height
}

func (fi *FractalImage) Pixel(x, y int) *Pixel {
	if y*fi.Width+x >= len(fi.Data) || y*fi.Width+x < 0 {
		return &Pixel{}
	}

	return &fi.Data[y*fi.Width+x]
}

func (fi *FractalImage) GetAspectRatio() float64 {
	if fi.Height == 0 {
		return 0
	}

	return float64(fi.Width) / float64(fi.Height)
}

func (fi *FractalImage) GammaCorrection(gamma float64) {
	var maxNormal float64

	for row := 0; row < fi.Width; row++ {
		for col := 0; col < fi.Height; col++ {
			pixel := fi.Pixel(row, col)
			if pixel.HitCount != 0 {
				pixel.Normal = math.Log10(float64(pixel.HitCount))
				if pixel.Normal > maxNormal {
					maxNormal = pixel.Normal
				}
			}
		}
	}

	for row := 0; row < fi.Width; row++ {
		for col := 0; col < fi.Height; col++ {
			pixel := fi.Pixel(row, col)
			if pixel.HitCount != 0 {
				pixel.Normal /= maxNormal
				pixel.RGB.R = int(float64(pixel.RGB.R) * math.Pow(pixel.Normal, 1.0/gamma))
				pixel.RGB.G = int(float64(pixel.RGB.G) * math.Pow(pixel.Normal, 1.0/gamma))
				pixel.RGB.B = int(float64(pixel.RGB.B) * math.Pow(pixel.Normal, 1.0/gamma))
			}
		}
	}
}
