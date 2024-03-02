package main

import (
	"fmt"
	"time"
)

func main() {
	printLog("Программа запущена")

	var age int

	fmt.Print("Введите свой возраст: ")
	fmt.Scanln(&age)

	yearsToJubilee := (10 - (age % 10)) % 10

	printLog(fmt.Sprintf("До вашего следующего юбилея осталось: %d лет", yearsToJubilee))

	printLog("Хэллоу")
	printLog("Программа завершена")
	fmt.Scan(&age)
}

func printLog(message string) {
	currentTime := time.Now()
	fmt.Printf("%02d:%02d %02d/%02d/%d %s\n", currentTime.Hour(), currentTime.Minute(), currentTime.Month(), currentTime.Day(), currentTime.Year(), message)

}
