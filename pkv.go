package main

import "fmt"

func main() {
	var a int
	fmt.Print("Введите длину квадрата : ")
	fmt.Scanln(&a)
	perimeter := 4 * a
	fmt.Printf("Периметр квадрата равен %d", perimeter)
	fmt.Scan(&a)
}
