package domain

import (
	randgenerator "github.com/es-debug/backend-academy-2024-go-template/pkg/rand_generator"
)

type AffineTransformation struct {
	A, B, C, D, E, F float64 //  Coefficients of the affine transformation matrix.
	RGB              RGB
}

func NewAffineTransformationMatrix() *AffineTransformation {
	var a, b, c, d, e, f float64

	for {
		a = randgenerator.Float64()*2 - 1
		b = randgenerator.Float64()*2 - 1
		d = randgenerator.Float64()*2 - 1
		e = randgenerator.Float64()*2 - 1
		c = randgenerator.Float64()*2 - 1
		f = randgenerator.Float64()*2 - 1

		// Check that the transformation is compressive.
		if a*a+d*d < 1 && b*b+e*e < 1 && a*a+b*b+d*d+e*e < 1+(a*e-b*d)*(a*e-b*d) {
			break
		}
	}

	return &AffineTransformation{
		A: a, B: b, C: c, D: d, E: e, F: f,
		RGB: RGB{
			randgenerator.Uint8(255),
			randgenerator.Uint8(255),
			randgenerator.Uint8(255),
		},
	}
}

func (t *AffineTransformation) Apply(x, y float64) (newX, newY float64) {
	newX = t.A*x + t.B*y + t.C
	newY = t.D*x + t.E*y + t.F

	return newX, newY
}
