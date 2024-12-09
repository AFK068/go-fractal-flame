package transformations_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	transformations "github.com/es-debug/backend-academy-2024-go-template/pkg/nonlinear_transformations"
)

func TestSinusoidalTransformation(t *testing.T) {
	transformation := transformations.SinusoidalTransformation{}
	newX, newY := transformation.Apply(math.Pi/4, math.Pi/2)

	assert.InEpsilon(t, math.Sqrt(2)/2, newX, 1e-9)
	assert.InEpsilon(t, 1.0, newY, 1e-9)
}

func TestSphericalTransformation(t *testing.T) {
	transformation := transformations.SphericalTransformation{}
	newX, newY := transformation.Apply(1.0, 2.0)

	assert.InEpsilon(t, 1.0/5.0, newX, 1e-9)
	assert.InEpsilon(t, 2.0/5.0, newY, 1e-9)
}

func TestSwirlTransformation(t *testing.T) {
	transformation := transformations.SwirlTransformation{}
	newX, newY := transformation.Apply(1.0, 2.0)

	assert.InEpsilon(t, math.Atan(2.0/1.0)/math.Pi, newX, 1e-9)
	assert.InEpsilon(t, math.Sqrt(1.0*1.0+2.0*2.0)-1, newY, 1e-9)
}

func TestHeartTransformation(t *testing.T) {
	transformation := transformations.HeartTransformation{}
	newX, newY := transformation.Apply(1.0, 2.0)

	expectedX := math.Sqrt(1.0*1.0+2.0*2.0) * math.Sin(math.Sqrt(1.0*1.0+2.0*2.0)*math.Atan(2.0/1.0))
	expectedY := -math.Sqrt(1.0*1.0+2.0*2.0) * math.Cos(math.Sqrt(1.0*1.0+2.0*2.0)*math.Atan(2.0/1.0))

	assert.InEpsilon(t, expectedX, newX, 1e-9)
	assert.InEpsilon(t, expectedY, newY, 1e-9)
}

func TestDiscTransformation(t *testing.T) {
	transformation := transformations.DiscTransformation{}
	newX, newY := transformation.Apply(1.0, 2.0)

	expectedX := 1 / math.Pi * math.Atan(2.0/1.0) * math.Sin(math.Pi*math.Sqrt(1.0*1.0+2.0*2.0))
	expectedY := 1 / math.Pi * math.Atan(2.0/1.0) * math.Cos(math.Pi*math.Sqrt(1.0*1.0+2.0*2.0))

	assert.InEpsilon(t, expectedX, newX, 1e-9)
	assert.InEpsilon(t, expectedY, newY, 1e-9)
}

func TestExTransformation(t *testing.T) {
	transformation := transformations.ExTransformation{}
	newX, newY := transformation.Apply(1.0, 2.0)

	r := math.Sqrt(1.0*1.0 + 2.0*2.0)
	p0 := math.Sin(math.Atan(1.0/2.0) + r)
	p1 := math.Cos(math.Atan(1.0/2.0) - r)

	expectedX := r * (p0*p0*p0 + p1*p1*p1)
	expectedY := r * (p0*p0*p0 - p1*p1*p1)

	assert.InEpsilon(t, expectedX, newX, 1e-9)
	assert.InEpsilon(t, expectedY, newY, 1e-9)
}

func TestExponentialTransformation(t *testing.T) {
	transformation := transformations.ExponentialTransformation{}
	newX, newY := transformation.Apply(1.0, 2.0)

	expectedX := math.Exp(1.0-1) * math.Cos(math.Pi*2.0)
	expectedY := math.Exp(1.0-1) * math.Sin(math.Pi*2.0)

	assert.InEpsilon(t, expectedX, newX, 1e-9)
	assert.InEpsilon(t, expectedY, newY, 1e-9)
}

func TestFisheyeTransformation(t *testing.T) {
	transformation := transformations.FisheyeTransformation{}
	newX, newY := transformation.Apply(1.0, 2.0)

	coef := 2 / (1 + math.Sqrt(1.0*1.0+2.0*2.0))
	expectedX := 1.0 * coef
	expectedY := 2.0 * coef

	assert.InEpsilon(t, expectedX, newX, 1e-9)
	assert.InEpsilon(t, expectedY, newY, 1e-9)
}

func TestCosineTransformation(t *testing.T) {
	transformation := transformations.CosineTransformation{}
	newX, newY := transformation.Apply(1.0, 2.0)

	expectedX := math.Cos(math.Pi*1.0) * math.Cosh(2.0)
	expectedY := -1 * math.Sin(math.Pi*1.0) * math.Sinh(2.0)

	assert.InEpsilon(t, expectedX, newX, 1e-9)
	assert.InEpsilon(t, expectedY, newY, 1e-9)
}

func TestTangentTransformation(t *testing.T) {
	transformation := transformations.TangentTransformation{}
	newX, newY := transformation.Apply(1.0, 2.0)

	expectedX := math.Sin(1.0) / math.Cos(2.0)
	expectedY := math.Tan(2.0)

	assert.InEpsilon(t, expectedX, newX, 1e-9)
	assert.InEpsilon(t, expectedY, newY, 1e-9)
}

func TestNoiseTransformation(t *testing.T) {
	transformation := transformations.NoiseTransformation{}
	newX, newY := transformation.Apply(1.0, 1.0)

	assert.True(t, newX >= -1.0+1e-9 && newX <= 1+1e-9)
	assert.True(t, newY >= -1.0-1e-9 && newY <= 1+1e-9)
}
