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

func (g *Generator) Generate(n, it int) {
	var wg sync.WaitGroup

	wg.Add(len(g.NonlinearTransformation))

	for routine := 0; routine < len(g.NonlinearTransformation); routine++ {
		go g.Render(n, it, (g.NonlinearTransformation)[routine], &wg)
	}

	wg.Wait()
}

func (g *Generator) Render(n, it int, nonlinearTransformations Transformation, wg *sync.WaitGroup) {
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

			if step >= 0 && newX >= XMIN && newX <= XMAX && newY >= YMIN && newY <= YMAX {
				x1 := int((newX - XMIN) / (XMAX - XMIN) * float64(g.FractalImage.Width))
				y1 := int((newY - YMIN) / (YMAX - YMIN) * float64(g.FractalImage.Height))

				if g.FractalImage.Contains(x1, y1) {
					pixel := g.FractalImage.Pixel(x1, y1)

					pixel.UpdateColorAndHit(affineTransformation.RGB)
				}
			}
		}
	}
}
