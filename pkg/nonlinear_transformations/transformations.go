package transformations

import (
	"math"

	randgenerator "github.com/es-debug/backend-academy-2024-go-template/pkg/rand_generator"
)

type SinusoidalTransformation struct{}

func (t *SinusoidalTransformation) Apply(x, y float64) (newX, newY float64) {
	newX = math.Sin(x)
	newY = math.Sin(y)

	return
}

type SphericalTransformation struct{}

func (t *SphericalTransformation) Apply(x, y float64) (newX, newY float64) {
	newX = x / (x*x + y*y)
	newY = y / (x*x + y*y)

	return
}

type SwirlTransformation struct{}

func (t *SwirlTransformation) Apply(x, y float64) (newX, newY float64) {
	newX = math.Atan(y/x) / math.Pi
	newY = math.Sqrt(x*x+y*y) - 1

	return
}

type HeartTransformation struct{}

func (t *HeartTransformation) Apply(x, y float64) (newX, newY float64) {
	newX = math.Sqrt(x*x+y*y) * math.Sin(math.Sqrt(x*x+y*y)*math.Atan(y/x))
	newY = -math.Sqrt(x*x+y*y) * math.Cos(math.Sqrt(x*x+y*y)*math.Atan(y/x))

	return
}

type DiscTransformation struct{}

func (t *DiscTransformation) Apply(x, y float64) (newX, newY float64) {
	newX = 1 / math.Pi * math.Atan(y/x) * math.Sin(math.Pi*math.Sqrt(x*x+y*y))
	newY = 1 / math.Pi * math.Atan(y/x) * math.Cos(math.Pi*math.Sqrt(x*x+y*y))

	return
}

type ExTransformation struct{}

func (t *ExTransformation) Apply(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	p0 := math.Sin(math.Atan(x/y) + r)
	p1 := math.Cos(math.Atan(x/y) - r)

	newX = r * (p0*p0*p0 + p1*p1*p1)
	newY = r * (p0*p0*p0 - p1*p1*p1)

	return
}

type FisheyeTransformation struct{}

func (t *FisheyeTransformation) Apply(x, y float64) (newX, newY float64) {
	coef := 2 / (1 + math.Sqrt(x*x+y*y))

	newX = x * coef
	newY = y * coef

	return
}

type CosineTransformation struct{}

func (t *CosineTransformation) Apply(x, y float64) (newX, newY float64) {
	newX = math.Cos(math.Pi*x) * math.Cosh(y)
	newY = -1 * math.Sin(math.Pi*x) * math.Sinh(y)

	return
}

type TangentTransformation struct{}

func (t *TangentTransformation) Apply(x, y float64) (newX, newY float64) {
	newX = math.Sin(x) / math.Cos(y)
	newY = math.Tan(y)

	return
}

type NoiseTransformation struct{}

func (t *NoiseTransformation) Apply(x, y float64) (newX, newY float64) {
	Psi1 := randgenerator.Float64()
	Psi2 := randgenerator.Float64()

	newX = Psi1 * x * math.Cos(2*math.Pi*Psi2)
	newY = Psi1 * y * math.Sin(2*math.Pi*Psi2)

	return
}
