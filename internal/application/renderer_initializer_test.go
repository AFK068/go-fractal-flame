package application_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	rendererMock "github.com/es-debug/backend-academy-2024-go-template/internal/application/mocks"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	transformations "github.com/es-debug/backend-academy-2024-go-template/pkg/nonlinear_transformations"
)

func TestInitializeRenderer(t *testing.T) {
	mockInitializer := rendererMock.NewRendererInitializer(t)

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

	renderer, err := application.InitializeRenderer(mockInitializer)

	assert.NoError(t, err)
	assert.NotNil(t, renderer)

	assert.Equal(t, image, renderer.FractalImage)
	assert.Equal(t, affineTransformations, renderer.AffineTransformation)
	assert.Equal(t, nonlinearTransformations, renderer.NonlinearTransformation)
}
