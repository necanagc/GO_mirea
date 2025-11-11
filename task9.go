package main

import "fmt"

func sum(a, b int) int {
	return a + b
}
func main() {
	var a, b int
	fmt.Println("Введите два целых числа")
	fmt.Scan(&a, &b)
	fmt.Println(sum(a, b))
}
