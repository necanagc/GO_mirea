package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите целое число")
	fmt.Scan(&n)
	if n%2 == 0 {
		if n == 0 {
			fmt.Println("zero")
		}
		if n > 0 {
			fmt.Println("positive even")
		}
		if n < 0 {
			fmt.Println("negative even")
		}
	} else {
		if n == 0 {
			fmt.Println("zero")
		}
		if n > 0 {
			fmt.Println("positive odd")
		}
		if n < 0 {
			fmt.Println("negative odd")
		}
	}

}
