package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test driven development")
	Test()
}
func Test() {
	o := calculate("+", 1, 1)
	if o != 2 {
		fmt.Printf("Test Failed ! expected %d output : %d \n", 2, o)
		return
	}
}
func calculate(op string, a, b int) int {
	return 0
}
