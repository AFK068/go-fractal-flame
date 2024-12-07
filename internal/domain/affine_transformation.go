package domain

type AffineTransformation struct {
	A, B, C, D, E, F float64 //  Coefficients of the affine transformation matrix.
	RGB              RGB
}
