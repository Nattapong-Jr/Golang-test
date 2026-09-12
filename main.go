package main

import (
	"fmt"
)

var number,number2 int = 1000, 2000
var msg string = "Hello"

func main() {

	numberfloat := 34.4
	fmt.Println("Hello, Son!")
	fmt.Println("The number is:", number)
	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(number + number2)
	fmt.Println(float64(number + number2) + numberfloat)
	fmt.Println(msg + "World")
	fmt.Println("My money = ",number2)
}