package main

import "fmt"

func main() {
	const PI float64 = 3.14159
	var r float64
	fmt.Println("Введите r")
	fmt.Scanln(&r)
	var area float64 = PI * r * r

	fmt.Printf("%.3f", area)
}
