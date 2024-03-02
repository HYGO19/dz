package main

import (
	"fmt"
)

func main() {
	var width, height int

	fmt.Println("Введите ширину прямоугольника:")
	fmt.Scanln(&width)

	fmt.Println("Введите высоту прямоугольника:")
	fmt.Scanln(&height)

	area := width * height
	fmt.Println("Площадь прямоугольника равна:", area)
	fmt.Scan(&width)
	fmt.Scan(&height)
}
