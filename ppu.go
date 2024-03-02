package main

import (
	"fmt"
)

func main() {
	var length, width int
	fmt.Print("Введите длину прямоугольника: ")
	fmt.Scanln(&length)
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scanln(&width)
	perimeter := 2 * (length + width)
	fmt.Printf("Периметр прямоугольника равен %d", perimeter)
	fmt.Scan(&length)
	fmt.Scan(&width)
}
