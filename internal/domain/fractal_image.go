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
