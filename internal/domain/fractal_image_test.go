package domain_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestFractalImage_Contains_success(t *testing.T) {
	width, height := 100, 100
	img := domain.NewFractalImage(width, height)

	assert.True(t, img.Contains(50, 50))
	assert.True(t, img.Contains(1, 1))
	assert.True(t, img.Contains(width-1, height-1))
	assert.True(t, img.Contains(1, height-1))
	assert.True(t, img.Contains(width-1, 1))
}

func TestFractalImage_Contains_failure(t *testing.T) {
	width, height := 100, 100
	img := domain.NewFractalImage(width, height)

	assert.False(t, img.Contains(-10, -10))
	assert.False(t, img.Contains(width+1, height+1))
	assert.False(t, img.Contains(width, height))
	assert.False(t, img.Contains(0, height))
	assert.False(t, img.Contains(width, 0))
}

func TestFractalImage_Pixel(t *testing.T) {
	t.Parallel()

	width, height := 100, 100
	img := domain.NewFractalImage(width, height)

	tests := []struct {
		x, y     int
		expected domain.Pixel
	}{
		{
			x:        50,
			y:        50,
			expected: domain.Pixel{X: 50, Y: 50},
		},
		{
			x:        0,
			y:        0,
			expected: domain.Pixel{X: 0, Y: 0},
		},
		{
			x:        99,
			y:        99,
			expected: domain.Pixel{X: 99, Y: 99},
		},
		{
			x:        -1,
			y:        -1,
			expected: domain.Pixel{X: 0, Y: 0},
		},
	}

	for i := range tests {
		tt := &tests[i]

		t.Run("", func(t *testing.T) {
			pixel := img.Pixel(tt.x, tt.y)
			assert.Equal(t, tt.expected.X, pixel.X)
			assert.Equal(t, tt.expected.Y, pixel.Y)
		})
	}
}

func TestFractalImage_GetAspectRatio(t *testing.T) {
	tests := []struct {
		width, height int
		expected      float64
	}{
		{width: 100, height: 50, expected: 2.0},
		{width: 50, height: 100, expected: 0.5},
		{width: 200, height: 100, expected: 2.0},
		{width: 100, height: 200, expected: 0.5},
		{width: 100, height: 0, expected: 0.0},
		{width: 0, height: 200, expected: 0.0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			img := domain.NewFractalImage(tt.width, tt.height)
			assert.Equal(t, tt.expected, img.GetAspectRatio())
		})
	}
}
