package application_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	generatorMock "github.com/es-debug/backend-academy-2024-go-template/internal/application/mocks"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	transformations "github.com/es-debug/backend-academy-2024-go-template/pkg/nonlinear_transformations"
	"github.com/stretchr/testify/assert"
)

func TestInitializeGenerator(t *testing.T) {
	mockInitializer := generatorMock.NewGeneratorInitializer(t)

	image := domain.NewFractalImage(1000, 1000)

	nonlinearTransformations := []domain.Transformation{
		&transformations.SinusoidalTransformation{},
		&transformations.SphericalTransformation{},
	}

	affineTransformations := []domain.AffineTransformation{
		*domain.NewAffineTransformation(),
		*domain.NewAffineTransformation(),
	}

	mockInitializer.On("InitializeFractalImage").Return(image, nil)
	mockInitializer.On("InitializeNonlinearTransformations").Return(nonlinearTransformations, nil)
	mockInitializer.On("InitializeAffineTransformations").Return(affineTransformations, nil)

	generator, err := application.InitializeGenerator(mockInitializer)

	assert.NoError(t, err)
	assert.NotNil(t, generator)

	assert.Equal(t, image, generator.FractalImage)
	assert.Equal(t, affineTransformations, generator.AffineTransformation)
	assert.Equal(t, nonlinearTransformations, generator.NonlinearTransformation)
}
