package main

import "fmt"

func main() {
	var a, b, res1, res2, res4 int
	var res3 float64
	fmt.Println("Введите два целых числа")
	fmt.Scan(&a, &b)
	var div0 bool = false
	if b != 0 {
		res3 = float64(a) / float64(b)
	} else {
		fmt.Println("Произошло деление на 0")
		div0 = true
	}

	res1 = a + b
	res2 = a - b

	res4 = a % b

	if div0 == true {
		fmt.Println(res1)
		fmt.Println(res2)

		fmt.Println(res4)

	} else {
		fmt.Println(res1)
		fmt.Println(res2)
		fmt.Printf("%.4f\n", res3)
		fmt.Println(res4)
	}

}
