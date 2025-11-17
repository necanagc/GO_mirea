package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {

	return float64(c)*9/5 + 32
}
func main() {
	var a float64
	fmt.Println("Введите число: ")
	fmt.Scanln(&a)

	b := Celsius(a).ToFahrenheit()

	fmt.Printf("%.2f", b)

}
