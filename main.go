package main

import (
	// "digitRegexp"
	"fmt"
)

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
// func Filter(s []int, fn func(int) bool) []int {
// 	var p []int // == nil
// 	for _, v := range s {
// 		if fn(v) {
// 			p = append(p, v)
// 		}
// 	}
// 	return p
// }

// func CopyDigits(filename string) []byte {
// 	b, _ := ioutil.ReadFile(filename)
// 	b = digitRegexp.Find(b)

// 	c := append([]byte{}, b...)

// 	return c
// }

func main() {
	// 1 задача
	week := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	s := week[2:]
	fmt.Println(s)
	res1, res2 := len(week), cap(week)

	res3, res4 := len(s), cap(s)

	fmt.Println("Длина и ёмкость массива:", res1, res2)
	fmt.Println("Длина и ёмкость среза:", res3, res4)

	// 2 задача

	fmt.Println(week)
	fmt.Println(week[1])
	fmt.Println(week[:1], week[2:])

	// 3 задача

	week1 := week[:]

	for i := 0; i < 7; i++ {
		week1 = append(week1, "Пятница")
	}
	fmt.Println(week1)

	// 4 задача

}
