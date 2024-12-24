package domain

import (
	"math/rand"
	"runtime"
	"sync"
)

const (
	countOfSteps = 100
)

type Transformation interface {
	Apply(x, y float64) (float64, float64)
}

type Renderer struct {
	FractalImage            *FractalImage
	AffineTransformation    []AffineTransformation
	NonlinearTransformation []Transformation
}

func NewRenderer(fractalImage FractalImage, affine []AffineTransformation, nonlinear []Transformation) *Renderer {
	return &Renderer{
		FractalImage:            &fractalImage,
		AffineTransformation:    affine,
		NonlinearTransformation: nonlinear,
	}
}

func (g *Renderer) Generate(n int, gamma float64, gammaFlag, symmetryFlag, useConcurrency bool) {
	numGoroutines := 1

	if useConcurrency {
		numGoroutines = runtime.NumCPU() * 2
	}

	sema := make(chan struct{}, numGoroutines)

	var wg sync.WaitGroup

	for _, transformation := range g.NonlinearTransformation {
		wg.Add(1)

		go func(trans Transformation) {
			defer wg.Done()
			sema <- struct{}{}

			g.render(n, trans, symmetryFlag)
			<-sema
		}(transformation)
	}

	wg.Wait()

	if gammaFlag {
		g.FractalImage.GammaCorrection(gamma)
	}
}

func (g *Renderer) render(n int, nonlinearTransformations Transformation, symmetryFlag bool) {
	XMIN, XMAX := g.FractalImage.GetAspectRatio()*-1, g.FractalImage.GetAspectRatio()
	YMIN, YMAX := -1.0, 1.0

	for num := 0; num < n; num++ {
		// Ctypto rand is work so slow, so we use math/rand.
		newX := XMIN + rand.Float64()*(XMAX-XMIN)
		newY := YMIN + rand.Float64()*(YMAX-YMIN)

		for step := -20; step < countOfSteps; step++ {
			affineTransformation := (g.AffineTransformation)[rand.Intn(len(g.AffineTransformation))]

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
