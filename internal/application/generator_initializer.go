package application

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

type GeneratorInitializer interface {
	InitializeFractalImage() (*domain.FractalImage, error)
	InitializeNonlinearTransformations() (nonlinearTransformations []domain.Transformation, err error)
	InitializeAffineTransformations() (affineTransformations []domain.AffineTransformation, err error)
}

func InitializeGenerator(initializer GeneratorInitializer) (*domain.Generator, error) {
	fractalImage, err := initializer.InitializeFractalImage()
	if err != nil {
		return nil, fmt.Errorf("initializing resolution: %w", err)
	}

	nonlinearTransformations, err := initializer.InitializeNonlinearTransformations()
	if err != nil {
		return nil, fmt.Errorf("initializing nonlinear transformations: %w", err)
	}

	affineTransformations, err := initializer.InitializeAffineTransformations()
	if err != nil {
		return nil, fmt.Errorf("initializing affine transformations: %w", err)
	}

	return domain.NewGenerator(*fractalImage, affineTransformations, nonlinearTransformations), nil
}
