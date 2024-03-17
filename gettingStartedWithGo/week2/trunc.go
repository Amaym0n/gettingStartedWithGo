package main

import "fmt"

func trunc(floatNum float64) int {
	return int(floatNum)
}

func main() {
	var inputFloat float64
	fmt.Scanln(&inputFloat)
	fmt.Println(trunc(inputFloat))
}
