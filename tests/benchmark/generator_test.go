package benchmark_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	transformations "github.com/es-debug/backend-academy-2024-go-template/pkg/nonlinear_transformations"
)

func BenchmarkRenderer_Sequential(b *testing.B) {
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

	renderer := domain.NewRenderer(*img, affineTransformations, nonlinearTransformations)

	for i := 0; i < b.N; i++ {
		renderer.Generate(1000000, 2.2, false, false, false)
	}
}

func BenchmarkRenderer_Concurrent(b *testing.B) {
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

	renderer := domain.NewRenderer(*img, affineTransformations, nonlinearTransformations)

	for i := 0; i < b.N; i++ {
		renderer.Generate(1000000, 2.2, false, false, true)
	}
}

func BenchmarkRendererWithGammaCorrection_Sequential(b *testing.B) {
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

	renderer := domain.NewRenderer(*img, affineTransformations, nonlinearTransformations)

	for i := 0; i < b.N; i++ {
		renderer.Generate(1000000, 2.2, true, false, false)
	}
}

func BenchmarkRendererWithGammaCorrection_Concurrent(b *testing.B) {
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

	renderer := domain.NewRenderer(*img, affineTransformations, nonlinearTransformations)

	for i := 0; i < b.N; i++ {
		renderer.Generate(1000000, 2.2, true, false, true)
	}
}
