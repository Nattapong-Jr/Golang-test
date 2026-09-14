package main

import "fmt"

//no return no parameter
func hello() {
	fmt.Println("Hello Mate")
}
// add parameter
func plus(value1 int, value2 int){
	result1 := value1 + value2
	fmt.Println("result = ", result1)
}
// add parameter and return
func plus2(value1 int, value2 int) int {
	return value1 + value2
}

func plus4value(value1, value2, value3, value4 int) int{
	return value1 + value2 + value3 + value4
}

func main() {
	hello()
	plus(1,2)
	result := plus2(2,2)
	fmt.Println("result = ", result)

	result1 := plus4value(20,20,20,7)
	fmt.Println("result = ", result1)
}
