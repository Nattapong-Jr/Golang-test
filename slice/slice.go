package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java","Python"}
	fmt.Println(courseName)

	courseName = append(courseName, "JavaScript", "C", "c++","HTML")
	fmt.Println(courseName)

	courseWeb := courseName[4:6]
	fmt.Println(courseWeb)

	courseWeb = courseName[:3]
	fmt.Println(courseWeb)
	
}