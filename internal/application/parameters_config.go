package application

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/userinteraction"
)

const (
	// Number of iterations.
	minCountOfIterations = 100000
	maxCountOfIterations = 3000000

	// Gamma value.
	minGamma = 0.1
	maxGamma = 10.0
)

type ParametersConfig struct{}

func (df *ParametersConfig) InitializeIterations() (iterations float64, err error) {
	iterations, err = userinteraction.GetValue(
		fmt.Sprintf("Enter iterations (min: %d, max: %d): ", minCountOfIterations, maxCountOfIterations),
		minCountOfIterations, maxCountOfIterations,
	)

	if err != nil {
		slog.Error("failed to get iterations", slog.String("error", err.Error()))

		return 0, fmt.Errorf("getting iterations: %w", err)
	}

	slog.Info("iterations initialized", slog.Int("iterations", int(iterations)))

	return iterations, nil
}

func (df *ParametersConfig) InitializeGamma() (gamma float64, flag bool, err error) {
	gammaFlag, err := userinteraction.GetTrueOrFalse("Do you want to set gamma?")
	if err != nil {
		slog.Error("failed to get gamma flag", slog.String("error", err.Error()))

		return 0, false, fmt.Errorf("getting gamma flag: %w", err)
	}

	if !gammaFlag {
		slog.Info("gamma flag initialized", slog.Bool("gamma flag", gammaFlag))

		return 0, false, nil
	}

	gamma, err = userinteraction.GetValue(
		fmt.Sprintf("Enter gamma (min: %.1f, max: %.1f, recommended: 2.2): ", minGamma, maxGamma),
		minGamma, maxGamma,
	)

	if err != nil {
		slog.Error("failed to get gamma", slog.String("error", err.Error()))

		return 0, false, fmt.Errorf("getting gamma: %w", err)
	}

	slog.Info("gamma initialized", slog.Float64("gamma", gamma))

	return gamma, true, nil
}

func (df *ParametersConfig) InitializeSymmetry() (bool, error) {
	symmetry, err := userinteraction.GetTrueOrFalse("Do you want to set symmetry?")
	if err != nil {
		slog.Error("failed to get symmetry", slog.String("error", err.Error()))

		return false, fmt.Errorf("getting symmetry: %w", err)
	}

	slog.Info("symmetry initialized", slog.Bool("symmetry", symmetry))

	return symmetry, nil
}

func (df *ParametersConfig) InitializeConcurrent() (bool, error) {
	concurrent, err := userinteraction.GetTrueOrFalse("Do you want to use concurrent mode?")
	if err != nil {
		slog.Error("failed to get concurrent", slog.String("error", err.Error()))

		return false, fmt.Errorf("getting concurrent: %w", err)
	}

	slog.Info("concurrent initialized", slog.Bool("concurrent", concurrent))

	return concurrent, nil
}
