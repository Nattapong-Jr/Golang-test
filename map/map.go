package main

import "fmt"

var product = make(map[string]float64)
func main() {
	fmt.Println("product = ", product)

	// add
	product["Macbook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 20000
	fmt.Println("product = ", product)

	//delet
	delete(product, "iPad")
	fmt.Println("product = ", product)

	//update
	product["iPhone"] = 25000
	fmt.Println("product = ", product)

	value := product["iPhone"]
	fmt.Println("value = ", value)

	courseName := map[string]string{"101":"Java","102":"Python"}
	fmt.Println("courseName = ", courseName)
}
