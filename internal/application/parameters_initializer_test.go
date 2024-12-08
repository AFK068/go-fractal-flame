package application_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	parametrsMock "github.com/es-debug/backend-academy-2024-go-template/internal/application/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInitializeParameters(t *testing.T) {
	mockInitializer := parametrsMock.NewParametersInitializer(t)

	iterations := 1000.0
	mockInitializer.On("InitializeIterations").Return(iterations, nil)

	gamma := 2.2
	gammaFlag := true
	mockInitializer.On("InitializeGamma").Return(gamma, gammaFlag, nil)

	symmetry := true
	mockInitializer.On("InitializeSymmetry").Return(symmetry, nil)

	concurrent := true
	mockInitializer.On("InitializeConcurrent").Return(concurrent, nil)

	parameters, err := application.InitializeParameters(mockInitializer)

	assert.NoError(t, err)
	assert.NotNil(t, parameters)

	assert.Equal(t, iterations, parameters.Iterations)
	assert.Equal(t, gamma, parameters.Gamma)
	assert.Equal(t, gammaFlag, parameters.GammaFlag)
	assert.Equal(t, symmetry, parameters.Symmetry)
}
