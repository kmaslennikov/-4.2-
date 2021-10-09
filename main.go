package main

import "fmt"

func main() {
	limit := 5

	var digit1 int
	var digit2 int
	var digit3 int
	fmt.Println("Три числа.")

	fmt.Println("Введите первое число:")
	fmt.Scan(&digit1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&digit2)
	fmt.Println("Введите третье число:")
	fmt.Scan(&digit3)

	if digit1 > limit {
		fmt.Println("Среди введённых чисел есть число больше", limit, ".")
	} else if digit2 > limit {
		fmt.Println("Среди введённых чисел есть число больше", limit, ".")
	} else if digit3 > limit {
		fmt.Println("Среди введённых чисел есть число больше", limit, ".")
	} else {
		fmt.Println("Среди введённых чисел нет числа больше", limit, ".")
	}
}
