package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (c *Circulo) getArea() float64 {
	return math.Pi * (c.Radio * c.Radio)
}

func (c *Circulo) getPerimetro() float64 {
	return 2 * math.Pi * c.Radio
}
