package application

import (
	"fmt"
	"log/slog"
)

type ParametersInitializer interface {
	InitializeIterations() (iterations float64, err error)
	InitializeGamma() (gamma float64, flag bool, err error)
	InitializeSymmetry() (bool, error)
	InitializeConcurrent() (bool, error)
}

type Parameters struct {
	Iterations float64
	Gamma      float64
	GammaFlag  bool
	Symmetry   bool
	Concurrent bool
}

func InitializeParameters(initializer ParametersInitializer) (*Parameters, error) {
	iterations, err := initializer.InitializeIterations()
	if err != nil {
		slog.Error("failed to initialize iterations", slog.String("error", err.Error()))

		return nil, fmt.Errorf("initializing iterations: %w", err)
	}

	gamma, gammaFlag, err := initializer.InitializeGamma()
	if err != nil {
		slog.Error("failed to initialize gamma", slog.String("error", err.Error()))

		return nil, fmt.Errorf("initializing gamma: %w", err)
	}

	symmetry, err := initializer.InitializeSymmetry()
	if err != nil {
		slog.Error("failed to initialize symmetry", slog.String("error", err.Error()))

		return nil, fmt.Errorf("initializing symmetry: %w", err)
	}

	concurrent, err := initializer.InitializeConcurrent()
	if err != nil {
		slog.Error("failed to initialize concurrent", slog.String("error", err.Error()))

		return nil, fmt.Errorf("initializing concurrent: %w", err)
	}

	slog.Info(
		"parameters initialized successfully",
		slog.Int("iterations", int(iterations)),
		slog.Float64("gamma", gamma),
		slog.Bool("gamma flag", gammaFlag),
		slog.Bool("symmetry", symmetry),
		slog.Bool("concurrent", concurrent),
	)

	return &Parameters{
		Iterations: iterations,
		Gamma:      gamma,
		GammaFlag:  gammaFlag,
		Symmetry:   symmetry,
		Concurrent: concurrent,
	}, nil
}
