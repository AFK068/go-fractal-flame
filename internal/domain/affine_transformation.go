package domain

import "math/rand"

type AffineTransformation struct {
	A, B, C, D, E, F float64 //  Coefficients of the affine transformation matrix.
	RGB              RGB
}

func NewAffineTransformation() *AffineTransformation {
	var a, b, c, d, e, f float64

	for {
		a = rand.Float64()*2 - 1 //nolint:gosec // crypto/rand bad performance
		b = rand.Float64()*2 - 1 //nolint:gosec // crypto/rand bad performance
		d = rand.Float64()*2 - 1 //nolint:gosec // crypto/rand bad performance
		e = rand.Float64()*2 - 1 //nolint:gosec // crypto/rand bad performance
		c = rand.Float64()*2 - 1 //nolint:gosec // crypto/rand bad performance
		f = rand.Float64()*2 - 1 //nolint:gosec // crypto/rand bad performance

		// Check that the transformation is compressive.
		if a*a+d*d < 1 && b*b+e*e < 1 && a*a+b*b+d*d+e*e < 1+(a*e-b*d)*(a*e-b*d) {
			break
		}
	}

	return &AffineTransformation{
		A: a, B: b, C: c, D: d, E: e, F: f,
		RGB: RGB{
			R: rand.Intn(256), //nolint:gosec // crypto/rand bad performance
			G: rand.Intn(256), //nolint:gosec // crypto/rand bad performance
			B: rand.Intn(256), //nolint:gosec // crypto/rand bad performance
		},
	}
}

func (t *AffineTransformation) Apply(x, y float64) (newX, newY float64) {
	newX = t.A*x + t.B*y + t.C
	newY = t.D*x + t.E*y + t.F

	return newX, newY
}
