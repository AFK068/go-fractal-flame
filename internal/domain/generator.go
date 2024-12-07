package domain

type Transformation interface {
	Apply(x, y float64) (float64, float64)
}

type Generator struct {
	FractalImage            *FractalImage
	AffineTransformation    []AffineTransformation
	NonlinearTransformation []Transformation
}
