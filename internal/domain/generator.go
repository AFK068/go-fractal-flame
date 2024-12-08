package domain

import (
	"math/rand"
	"sync"
)

type Transformation interface {
	Apply(x, y float64) (float64, float64)
}

type Generator struct {
	FractalImage            *FractalImage
	AffineTransformation    []AffineTransformation
	NonlinearTransformation []Transformation
}

func NewGenerator(fractalImage FractalImage, affine []AffineTransformation, nonlinear []Transformation) *Generator {
	return &Generator{
		FractalImage:            &fractalImage,
		AffineTransformation:    affine,
		NonlinearTransformation: nonlinear,
	}
}

func (g *Generator) Generate(n, it int, gamma float64, gammaFlag, symmetryFlag, useConcurrency bool) {
	if useConcurrency {
		g.generateConcurrent(n, it, symmetryFlag)
	} else {
		g.generateSequential(n, it, symmetryFlag)
	}

	if gammaFlag {
		g.FractalImage.GammaCorrection(gamma)
	}
}

// Multi-threaded version.
func (g *Generator) generateConcurrent(n, it int, symmetryFlag bool) {
	var wg sync.WaitGroup

	wg.Add(len(g.NonlinearTransformation))

	for _, transformation := range g.NonlinearTransformation {
		go g.render(n, it, transformation, symmetryFlag, &wg)
	}

	wg.Wait()
}

// Single-threaded version (it is said in the technical specifications to implement).
func (g *Generator) generateSequential(n, it int, symmetryFlag bool) {
	for _, transformation := range g.NonlinearTransformation {
		g.render(n, it, transformation, symmetryFlag, nil)
	}
}

func (g *Generator) render(n, it int, nonlinearTransformations Transformation, symmetryFlag bool, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	XMIN, XMAX := g.FractalImage.GetAspectRatio()*-1, g.FractalImage.GetAspectRatio()
	YMIN, YMAX := -1.0, 1.0

	for num := 0; num < n; num++ {
		// Ctypto rand is work so slow, so we use math/rand.
		newX := XMIN + rand.Float64()*(XMAX-XMIN) //nolint:gosec // crypto/rand bad performance
		newY := YMIN + rand.Float64()*(YMAX-YMIN) //nolint:gosec // crypto/rand bad performance

		for step := -20; step < it; step++ {
			affineTransformation := (g.AffineTransformation)[rand.Intn(len(g.AffineTransformation))] //nolint:gosec // crypto/rand bad performance

			// Apply affine transformations.
			newX, newY = affineTransformation.Apply(newX, newY)

			// Apply nonlinear transformations.
			newX, newY = nonlinearTransformations.Apply(newX, newY)

			// Symmetry.
			tempX, tempY := newX, newY

			if symmetryFlag {
				switch {
				case step%4 == 0:
					tempX = -tempX
				case step%3 == 0:
					tempX, tempY = -tempX, -tempY
				case step%2 == 0:
					tempY = -tempY
				}
			}

			if step >= 0 && newX >= XMIN && newX <= XMAX && newY >= YMIN && newY <= YMAX {
				x1 := int((tempX - XMIN) / (XMAX - XMIN) * float64(g.FractalImage.Width))
				y1 := int((tempY - YMIN) / (YMAX - YMIN) * float64(g.FractalImage.Height))

				if g.FractalImage.Contains(x1, y1) {
					pixel := g.FractalImage.Pixel(x1, y1)

					// Update pixel color and hit count.
					pixel.UpdateColorAndHit(affineTransformation.RGB)
				}
			}
		}
	}
}
