package application

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/userinteraction"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/menu"
)

const (
	// Width of the fractal image.
	minWidthImageSize = 100
	maxWidthImageSize = 9000

	// Height of the fractal image.
	minHeightImageSize = 100
	maxHeightImageSize = 9000

	// Number of affine transformations.
	minAffineTransformations = 1
	maxAffineTransformations = 10
)

type RendererConfig struct{}

func (df *RendererConfig) InitializeFractalImage() (*domain.FractalImage, error) {
	width, err := userinteraction.GetValue(
		fmt.Sprintf("Enter width (min: %d, max: %d): ", minWidthImageSize, maxWidthImageSize),
		minWidthImageSize, maxWidthImageSize,
	)

	if err != nil {
		slog.Error("failed to get width", slog.String("error", err.Error()))

		return nil, fmt.Errorf("getting width: %w", err)
	}

	height, err := userinteraction.GetValue(
		fmt.Sprintf("Enter height (min: %d, max: %d): ", minHeightImageSize, maxHeightImageSize),
		minHeightImageSize, maxHeightImageSize,
	)

	if err != nil {
		slog.Error("failed to get height", slog.String("error", err.Error()))

		return nil, fmt.Errorf("getting height: %w", err)
	}

	slog.Info("fractal image initialized", slog.Int("width", int(width)), slog.Int("height", int(height)))

	return domain.NewFractalImage(int(width), int(height)), nil
}

func (df *RendererConfig) InitializeAffineTransformations() (affineTransformations []domain.AffineTransformation, err error) {
	count, err := userinteraction.GetValue(
		fmt.Sprintf("Enter the number of affine transformations (min: %d, max: %d): ", minAffineTransformations, maxAffineTransformations),
		minAffineTransformations, maxAffineTransformations,
	)

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

func (df *RendererConfig) InitializeNonlinearTransformations() (nonlinearTransformations []domain.Transformation, err error) {
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
