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
	o = calculate("+", 5, 1)
	if o != 6 {
		fmt.Printf("Test Failed ! expected %d output : %d \n", 6, o)
		return
	}

	fmt.Println("Test success!")
}
func calculate(op string, a, b int) int {
	return 2
}
