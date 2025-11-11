package main

import "fmt"

func main() {
	var x int
	fmt.Println("Введите целое число")
	fmt.Scan(&x)
	x = x + 3
	fmt.Println("Промежуточное:", x)
	x += 2
	fmt.Println("Промежуточное:", x)
	x *= 4
	fmt.Println("Промежуточное:", x)
	x -= 10

	fmt.Println("Итоговое:", x)
}
