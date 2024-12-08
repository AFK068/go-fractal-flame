package application

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/menu"
)

type GeneratorConfig struct{}

func (df *GeneratorConfig) InitializeFractalImage() (*domain.FractalImage, error) {
	width, err := infrastructure.GetValue("Enter width (min: 100, max: 9000): ", 100, 9000)
	if err != nil {
		slog.Error("failed to get width", slog.String("error", err.Error()))

		return nil, fmt.Errorf("getting width: %w", err)
	}

	height, err := infrastructure.GetValue("Enter height (min: 100, max: 9000): ", 100, 9000)
	if err != nil {
		slog.Error("failed to get height", slog.String("error", err.Error()))

		return nil, fmt.Errorf("getting height: %w", err)
	}

	slog.Info("fractal image initialized", slog.Int("width", int(width)), slog.Int("height", int(height)))

	return domain.NewFractalImage(int(width), int(height)), nil
}

func (df *GeneratorConfig) InitializeAffineTransformations() (affineTransformations []domain.AffineTransformation, err error) {
	count, err := infrastructure.GetValue("Enter the number of affine transformations (min: 1, max: 10, recommended: 3): ", 1, 10)
	if err != nil {
		slog.Error("failed to get the number of affine transformations", slog.String("error", err.Error()))

		return nil, fmt.Errorf("getting the number of affine transformations: %w", err)
	}

	for i := 0; i < int(count); i++ {
		affineTransformations = append(affineTransformations, *domain.NewAffineTransformation())
	}

	slog.Info("affine transformations initialized", slog.Any("affine transformations", affineTransformations))

	return affineTransformations, nil
}

func (df *GeneratorConfig) InitializeNonlinearTransformations() (nonlinearTransformations []domain.Transformation, err error) {
	availableTransformations := GetTransformations()

	// Create a new menu.
	mainMenu := menu.NewMenu("Select transformations (use enter to select/unselect, esc to continue):")
	for _, t := range availableTransformations {
		mainMenu.AddItem(t.Name)
	}

	// Display the menu.
	selectedItems, err := mainMenu.Display()
	if err != nil {
		slog.Error("failed to display menu", slog.String("error", err.Error()))

		return nil, fmt.Errorf("displaying menu: %w", err)
	}

	// Get the selected transformations.
	var nonlinearTransformation []domain.Transformation

	selectedNames := make(map[string]bool)

	for _, item := range selectedItems {
		selectedNames[item.Text] = true
	}

	for _, t := range availableTransformations {
		if selectedNames[t.Name] {
			slog.Info("user selected transformation", slog.String("name", t.Name))

			nonlinearTransformation = append(nonlinearTransformation, t.Transformation)
		}
	}

	return nonlinearTransformation, nil
}
