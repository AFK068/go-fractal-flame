package benchmark_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	transformations "github.com/es-debug/backend-academy-2024-go-template/pkg/nonlinear_transformations"
)

func BenchmarkGenerate_Sequential(b *testing.B) {
	width, height := 2000, 2000
	img := domain.NewFractalImage(width, height)

	affineTransformations := []domain.AffineTransformation{
		*domain.NewAffineTransformation(),
		*domain.NewAffineTransformation(),
		*domain.NewAffineTransformation(),
	}

	nonlinearTransformations := []domain.Transformation{
		&transformations.SinusoidalTransformation{},
		&transformations.SphericalTransformation{},
		&transformations.SwirlTransformation{},
		&transformations.HeartTransformation{},
		&transformations.DiscTransformation{},
		&transformations.ExTransformation{},
		&transformations.FisheyeTransformation{},
		&transformations.CosineTransformation{},
		&transformations.TangentTransformation{},
		&transformations.NoiseTransformation{},
		&transformations.ExponentialTransformation{},
	}

	generator := domain.NewGenerator(*img, affineTransformations, nonlinearTransformations)

	for i := 0; i < b.N; i++ {
		generator.Generate(1000000, 100, 2.2, false, false, false)
	}
}

func BenchmarkGenerate_Concurrent(b *testing.B) {
	width, height := 2000, 2000
	img := domain.NewFractalImage(width, height)

	affineTransformations := []domain.AffineTransformation{
		*domain.NewAffineTransformation(),
		*domain.NewAffineTransformation(),
		*domain.NewAffineTransformation(),
	}

	nonlinearTransformations := []domain.Transformation{
		&transformations.SinusoidalTransformation{},
		&transformations.SphericalTransformation{},
		&transformations.SwirlTransformation{},
		&transformations.HeartTransformation{},
		&transformations.DiscTransformation{},
		&transformations.ExTransformation{},
		&transformations.FisheyeTransformation{},
		&transformations.CosineTransformation{},
		&transformations.TangentTransformation{},
		&transformations.NoiseTransformation{},
		&transformations.ExponentialTransformation{},
	}

	generator := domain.NewGenerator(*img, affineTransformations, nonlinearTransformations)

	for i := 0; i < b.N; i++ {
		generator.Generate(1000000, 100, 2.2, false, false, true)
	}
}

func BenchmarkGenerateWithGammaCorrection_Sequential(b *testing.B) {
	width, height := 2000, 2000
	img := domain.NewFractalImage(width, height)

	affineTransformations := []domain.AffineTransformation{
		*domain.NewAffineTransformation(),
		*domain.NewAffineTransformation(),
		*domain.NewAffineTransformation(),
	}

	nonlinearTransformations := []domain.Transformation{
		&transformations.SinusoidalTransformation{},
		&transformations.SphericalTransformation{},
		&transformations.SwirlTransformation{},
		&transformations.HeartTransformation{},
		&transformations.DiscTransformation{},
		&transformations.ExTransformation{},
		&transformations.FisheyeTransformation{},
		&transformations.CosineTransformation{},
		&transformations.TangentTransformation{},
		&transformations.NoiseTransformation{},
		&transformations.ExponentialTransformation{},
	}

	generator := domain.NewGenerator(*img, affineTransformations, nonlinearTransformations)

	for i := 0; i < b.N; i++ {
		generator.Generate(1000000, 100, 2.2, true, false, false)
	}
}

func BenchmarkGenerateWithGammaCorrection_Concurrent(b *testing.B) {
	width, height := 2000, 2000
	img := domain.NewFractalImage(width, height)

	affineTransformations := []domain.AffineTransformation{
		*domain.NewAffineTransformation(),
		*domain.NewAffineTransformation(),
		*domain.NewAffineTransformation(),
	}

	nonlinearTransformations := []domain.Transformation{
		&transformations.SinusoidalTransformation{},
		&transformations.SphericalTransformation{},
		&transformations.SwirlTransformation{},
		&transformations.HeartTransformation{},
		&transformations.DiscTransformation{},
		&transformations.ExTransformation{},
		&transformations.FisheyeTransformation{},
		&transformations.CosineTransformation{},
		&transformations.TangentTransformation{},
		&transformations.NoiseTransformation{},
		&transformations.ExponentialTransformation{},
	}

	generator := domain.NewGenerator(*img, affineTransformations, nonlinearTransformations)

	for i := 0; i < b.N; i++ {
		generator.Generate(1000000, 100, 2.2, true, false, true)
	}
}
