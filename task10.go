package main

import "fmt"

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	var a, b int
	fmt.Println("Введите два целых числа")
	fmt.Scan(&a, &b)

	if b == 0 {
		fmt.Println("division by zero")
	} else {
		var res1, res2 int = divmod(a, b)
		fmt.Println(res1, res2)
	}
}
