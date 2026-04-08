package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

// Calcula a area do retangulo usando a referencia de um retangulo
func (rectangle *Rectangle) calculate_rectangle_area() float64 {
	return rectangle.width * rectangle.height
}

// Escala a altura e a largura do retângulo original, modificando via ponteiro
func (rectangle *Rectangle) scale(factor float64) {
	rectangle.width *= factor
	rectangle.height *= factor
}

func main() {
	// Declara rectangle como um ponteiro para Rectangle, ou seja, armazena o endereço de memória do struct, não uma cópia dele.
	var rectangle *Rectangle = &Rectangle{
		width:  10.25,
		height: 3.24,
	}

	fmt.Println("Original Width", rectangle.width)
	fmt.Println("Original Height", rectangle.height)

	rectangle_area := rectangle.calculate_rectangle_area()

	fmt.Println("Area", rectangle_area)

	// Dobra as dimensões do retângulo original (factor = 2)
	rectangle.scale(2)

	fmt.Println(rectangle.width)  // 20.50
	fmt.Println(rectangle.height) // 6.48

	fmt.Println("Nova area: ", rectangle.calculate_rectangle_area())
}
