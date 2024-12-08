package main

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/logger"
)

func main() {
	customLogger, err := logger.InitLogger()
	if err != nil {
		fmt.Println("failed to initialize logger:", err)
		return
	}

	slog.SetDefault(customLogger.Logger)

	defer func() {
		if err := logger.CloseLogger(customLogger); err != nil {
			fmt.Println("failed to close logger:", err)
		}
	}()

	generatorConfig := application.GeneratorConfig{}

	generator, err := application.InitializeGenerator(&generatorConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	parametersConfig := application.ParametersConfig{}

	parametrs, err := application.InitializeParameters(&parametersConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	generator.Generate(
		int(parametrs.Iterations),
		100, // Count of it.
		parametrs.Gamma,
		parametrs.GammaFlag,
		parametrs.Symmetry,
		parametrs.Concurrent,
	)

	err = infrastructure.SaveImage(generator.FractalImage, "fractal.png")
	if err != nil {
		fmt.Println(err)
		return
	}
}
