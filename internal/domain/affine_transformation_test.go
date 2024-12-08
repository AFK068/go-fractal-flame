package domain_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewAffineTransformation(t *testing.T) {
	transformation := domain.NewAffineTransformation()

	// Check if the transformation is compressive.
	assert.Less(t, transformation.A*transformation.A+transformation.D*transformation.D, 1.0, "A*A + D*D should be less than 1")
	assert.Less(t, transformation.B*transformation.B+transformation.E*transformation.E, 1.0, "B*B + E*E should be less than 1")
	assert.Less(
		t, transformation.A*transformation.A+
			transformation.B*transformation.B+
			transformation.D*transformation.D+
			transformation.E*transformation.E,
		1.0+(transformation.A*transformation.E-transformation.B*transformation.D)*
			(transformation.A*transformation.E-transformation.B*transformation.D),
		"A*A + B*B + D*D + E*E should be less than 1 + (A*E - B*D)*(A*E - B*D)",
	)

	// Check that RGB values are within acceptable limits.
	assert.GreaterOrEqual(t, transformation.RGB.R, 0, "R component should be greater than or equal to 0")
	assert.Less(t, transformation.RGB.R, 256, "R component should be less than 256")

	assert.GreaterOrEqual(t, transformation.RGB.G, 0, "G component should be greater than or equal to 0")
	assert.Less(t, transformation.RGB.G, 256, "G component should be less than 256")

	assert.GreaterOrEqual(t, transformation.RGB.B, 0, "B component should be greater than or equal to 0")
	assert.Less(t, transformation.RGB.B, 256, "B component should be less than 256")
}

func TestAffineTransformationApply(t *testing.T) {
	transformations := domain.NewAffineTransformation()

	// Set the transformation matrix.
	transformations.A = 0.3
	transformations.B = 0.4
	transformations.D = 0.2
	transformations.E = 0.5

	transformations.C = 0.1
	transformations.F = 0.15

	// Check that the transformation is applied correctly.
	x, y := 1.0, 2.0
	newX, newY := transformations.Apply(x, y)

	assert.Equal(t, 0.3*x+0.4*y+0.1, newX, "New X should be equal to 0.3*1.0 + 0.4*2.0 + 0.1")
	assert.Equal(t, 0.2*x+0.5*y+0.15, newY, "New Y should be equal to 0.2*1.0 + 0.5*2.0 + 0.15")
}
