package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) calculate_rectangle_area() float64 {
	return r.width * r.height
}

func main() {
	var rectangle *Rectangle = &Rectangle{
		width: 10.25,
		height: 3.24,
	}

	rectangle_area := rectangle.calculate_rectangle_area()

	fmt.Println("Area", rectangle_area)

}
