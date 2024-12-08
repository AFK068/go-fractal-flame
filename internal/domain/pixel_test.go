package domain_test

import (
	"sync"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUpdateColorAndHit(t *testing.T) {
	pixel := &domain.Pixel{}

	colorValue := domain.RGB{R: 100, G: 100, B: 100}

	var wg sync.WaitGroup

	// Test that the pixel color is updated correctly.
	// That is count of goroutines is equal to HitCount.
	numGoroutines := 3000
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			pixel.UpdateColorAndHit(colorValue)
		}()
	}

	wg.Wait()

	assert.Equal(
		t, (numGoroutines), int(pixel.HitCount), //nolint:gosec // unreal overflow
		"HitCount should be equal to the number of goroutines",
	)
}
