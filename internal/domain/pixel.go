package domain

import "sync"

type Pixel struct {
	mt       sync.Mutex
	X, Y     int
	RGB      RGB
	HitCount uint64
}

// Hit increments the hit count of the pixel.
func (p *Pixel) UpdateColorAndHit(rgb RGB) {
	p.mt.Lock()
	defer p.mt.Unlock()

	if p.HitCount == 0 {
		p.RGB = rgb
	} else {
		p.RGB.R = (p.RGB.R + rgb.R) / 2
		p.RGB.G = (p.RGB.G + rgb.G) / 2
		p.RGB.B = (p.RGB.B + rgb.B) / 2
	}

	p.HitCount++
}
