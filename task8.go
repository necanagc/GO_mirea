package main

import "fmt"

func main() {

	var n int
	var res int
	res = 0
	fmt.Println("Введите число N >= 0")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		res += i
	}

	fmt.Println(res)

}
