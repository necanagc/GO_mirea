package main

import "fmt"

func main() {
	var a string
	fmt.Println("Введите своё имя")
	fmt.Scan(&a)
	fmt.Printf("Hello, %s", a)

}
