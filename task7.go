package main

import "fmt"

func main() {
	var score int
	fmt.Println("Введите целое число")
	fmt.Scan(&score)
	if score > 100 || score < 0 {
		score = score / 10
	}

	switch score {
	case 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100:
		fmt.Println("A")
	case 80, 81, 82, 83, 84, 85, 86, 87, 88, 89:
		fmt.Println("B")
	case 70, 71, 72, 73, 74, 75, 76, 77, 78, 79:
		fmt.Println("C")
	case 60, 61, 62, 63, 64, 65, 66, 67, 68, 69:
		fmt.Println("D")

	default:
		fmt.Println("F")
	}

}
