package figuras

type Cuadrado struct {
	Ancho  float64
	Altura float64
}

func (c *Cuadrado) getArea() float64 {
	return c.Ancho * c.Altura
}

func (c *Cuadrado) getPerimetro() float64 {
	return 2*c.Ancho + 2*c.Altura
}
