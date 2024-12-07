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

func (g *Generator) Generate(n, it int, gamma float64, gammaFlag, symmetryFlag bool) {
	var wg sync.WaitGroup

	wg.Add(len(g.NonlinearTransformation))

	for routine := 0; routine < len(g.NonlinearTransformation); routine++ {
		go g.Render(n, it, (g.NonlinearTransformation)[routine], symmetryFlag, &wg)
	}

	wg.Wait()

	if gammaFlag {
		g.FractalImage.GammaCorrection(gamma)
	}
}

func (g *Generator) Render(n, it int, nonlinearTransformations Transformation, symmetryFlag bool, wg *sync.WaitGroup) {
	defer wg.Done()

	XMIN, XMAX := g.FractalImage.GetAspectRatio()*-1, g.FractalImage.GetAspectRatio()
	YMIN, YMAX := -1.0, 1.0

	for num := 0; num < n; num++ {
		newX := XMIN + rand.Float64()*(XMAX-XMIN)
		newY := YMIN + rand.Float64()*(YMAX-YMIN)

		for step := -20; step < it; step++ {
			affineTransformation := (g.AffineTransformation)[rand.Intn(len(g.AffineTransformation))]

			newX, newY = affineTransformation.Apply(newX, newY)
			newX, newY = nonlinearTransformations.Apply(newX, newY)

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

					pixel.UpdateColorAndHit(affineTransformation.RGB)
				}
			}
		}
	}
}
