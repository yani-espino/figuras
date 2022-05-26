package figuras

import "fmt"

type FiguraI interface {
	getArea() float64
	getPerimetro() float64
}

func CalculaArea(figura FiguraI) {
	fmt.Println(figura.getArea())
}

func CalculaPerimetro(figura FiguraI) {
	fmt.Println(figura.getPerimetro())
}
