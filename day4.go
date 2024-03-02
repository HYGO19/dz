package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Println("Программа запущена")
	defer log.Println("Программа завершена")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(`Здравствуйте! Пожалуйста, выберите день обучения:
- 1 День. Знакомство
- 2 День. Первая программа
...`)

	day, _ := reader.ReadString('\n')
	day = strings.TrimSpace(day)
	//day = day[:len(day)-1] // Удаление символа новой строки

	dayInt, err := strconv.Atoi(day)
	if err != nil || dayInt < 1 || dayInt > 2 { // Предполагаем, что у нас есть информация только за 2 дня
		fmt.Println("Извините, информация за этот день ещё не доступна.")
		return
	}

	switch dayInt {
	case 1:
		fmt.Println("На первый день домашних заданий нет.")
	case 2:
		fmt.Println(`Выберите задание:
- 1. Периметр квадрата
- 2. Бонусное задание: Таблица умножения
- 3. Бонусное задание: Числа Фибоначчи`)

		task, _ := reader.ReadString('\n')
		task = strings.TrimSpace(task)
		taskInt, err := strconv.Atoi(task)

		if err != nil || taskInt < 1 || taskInt > 3 {
			fmt.Println("Такого задания не существует.")
			return
		}

		switch taskInt {
		case 1:
			fmt.Println("Задание на периметр квадрата.")
			calculateSquarePerimeter()
		case 2:
			printMultiplicationTable()
		case 3:
			fmt.Println("Введите число для расчета числа Фибоначчи:")
			nStr, _ := reader.ReadString('\n')
			nStr = strings.TrimSpace(nStr) // Удаление символа новой строки
			n, _ := strconv.Atoi(nStr)
			fmt.Printf("Число Фибоначчи для %d равно %d\n", n, fibonacci(n))
		}
	}
}

func calculateSquarePerimeter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите длину стороны квадрата:")

	sideStr, _ := reader.ReadString('\n')
	sideStr = strings.TrimSpace(sideStr)
	side, err := strconv.Atoi(sideStr)

	if err != nil {
		fmt.Println("Ошибка: введено не число.")
		return
	}

	perimeter := 4 * side
	fmt.Printf("Периметр квадрата со стороной %d равен %d.\n", side, perimeter)
}

func printMultiplicationTable() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
