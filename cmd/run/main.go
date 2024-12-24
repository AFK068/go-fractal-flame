package main

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/imageutils"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/logger"
)

const (
	saveImageFilename = "fractal.png"
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

	rendererConfig := application.RendererConfig{}

	renderer, err := application.InitializeRenderer(&rendererConfig)
	if err != nil {
		slog.Error("failed to initialize renderer", slog.String("error", err.Error()))

		return
	}

	parametersConfig := application.ParametersConfig{}

	parametrs, err := application.InitializeParameters(&parametersConfig)
	if err != nil {
		slog.Error("failed to initialize parameters", slog.String("error", err.Error()))

		return
	}

	renderer.Generate(
		int(parametrs.Iterations),
		parametrs.Gamma,
		parametrs.GammaFlag,
		parametrs.Symmetry,
		parametrs.Concurrent,
	)

	err = imageutils.SaveImage(renderer.FractalImage, saveImageFilename)
	if err != nil {
		slog.Error("failed to save image", slog.String("error", err.Error()))

		return
	}

	slog.Info("image saved", slog.String("filename", saveImageFilename))
}
