package application

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

type rendererInitializer interface {
	InitializeFractalImage() (*domain.FractalImage, error)
	InitializeNonlinearTransformations() (nonlinearTransformations []domain.Transformation, err error)
	InitializeAffineTransformations() (affineTransformations []domain.AffineTransformation, err error)
}

func InitializeRenderer(initializer rendererInitializer) (*domain.Renderer, error) {
	fractalImage, err := initializer.InitializeFractalImage()
	if err != nil {
		slog.Error("failed to initialize fractal image", slog.String("error", err.Error()))

		return nil, fmt.Errorf("initializing resolution: %w", err)
	}

	nonlinearTransformations, err := initializer.InitializeNonlinearTransformations()
	if err != nil {
		slog.Error("failed to initialize nonlinear transformations", slog.String("error", err.Error()))

		return nil, fmt.Errorf("initializing nonlinear transformations: %w", err)
	}

	affineTransformations, err := initializer.InitializeAffineTransformations()
	if err != nil {
		slog.Error("failed to initialize affine transformations", slog.String("error", err.Error()))

		return nil, fmt.Errorf("initializing affine transformations: %w", err)
	}

	slog.Info("renderer initialized successfully")

	return domain.NewRenderer(*fractalImage, affineTransformations, nonlinearTransformations), nil
}
