package domain

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
	return &fi.Data[y*fi.Width+x]
}

func (fi *FractalImage) GetAspectRatio() float64 {
	return float64(fi.Width) / float64(fi.Height)
}
