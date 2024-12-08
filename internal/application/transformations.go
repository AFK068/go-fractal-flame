package application

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	transformations "github.com/es-debug/backend-academy-2024-go-template/pkg/nonlinear_transformations"
)

type TransformationInfo struct {
	Transformation domain.Transformation
	Name           string
}

func GetTransformations() []TransformationInfo {
	nonLinearTransformations := []TransformationInfo{
		{&transformations.SinusoidalTransformation{}, "Sinusoidal"},
		{&transformations.SphericalTransformation{}, "Spherical"},
		{&transformations.SwirlTransformation{}, "Swirl"},
		{&transformations.HeartTransformation{}, "Heart"},
		{&transformations.DiscTransformation{}, "Disc"},
		{&transformations.ExTransformation{}, "Ex"},
		{&transformations.FisheyeTransformation{}, "Fisheye"},
		{&transformations.CosineTransformation{}, "Cosine"},
		{&transformations.TangentTransformation{}, "Tangent"},
		{&transformations.NoiseTransformation{}, "Noise"},
		{&transformations.ExponentialTransformation{}, "Exponential"},
	}

	return nonLinearTransformations
}
