package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test driven development")
	Test()
	Test2()
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
	o = calculate("+", 5, -1)
	if o != 4 {
		fmt.Printf("Test Failed ! expected %d output : %d \n", 4, o)
		return
	}
	o = calculate("-", 5, 1)
	if o != 4 {
		fmt.Printf("Test Failed ! expected %d output : %d \n", 4, o)
		return
	}
	o = calculate("-", 5, 5)
	if o != 0 {
		fmt.Printf("Test Failed ! expected %d output : %d \n", 0, o)
		return
	}
	o = calculate("*", 5, 5)
	if o != 25 {
		fmt.Printf("Test Failed ! expected %d output : %d \n", 25, o)
		return
	}
	o = calculate("*", 5, -5)
	if o != -25 {
		fmt.Printf("Test Failed ! expected %d output : %d \n", -25, o)
		return
	}
	o = calculate("/", 5, 5)
	if o != 1 {
		fmt.Printf("Test Failed ! expected %d output : %d \n", 1, o)
		return
	}
	fmt.Println("Test success!")
}
func Test2() {
	if !testCalculate("test1", "+", 1, 1, 2) {
		return
	}
	if !testCalculate("test2", "+", 5, 1, 6) {
		return
	}
	if !testCalculate("test3", "+", 5, -1, 4) {
		return
	}
	if !testCalculate("test4", "-", 5, 1, 4) {
		return
	}
	if !testCalculate("test5", "-", 5, 5, 0) {
		return
	}
	if !testCalculate("test6", "*", 5, 5, 25) {
		return
	}
	if !testCalculate("test7", "*", 5, -5, -25) {
		return
	}
	if !testCalculate("test8", "/", 5, 5, 1) {
		return
	}
}
func testCalculate(testcase, op string, a, b, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed! expected : %d output: %d \n", testcase, expected, o)
		return false
	}
	return true
}
func calculate(op string, a, b int) int {
	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "*" {
		return a * b
	} else if op == "/" {
		return a / b
	}
	return 0
}
func Calculate(op string, a, b int) int {
	res := 0
	if op == "+" {
		res = a + b
	} else if op == "-" {
		res = a - b
	} else if op == "*" {
		res = a * b
	} else if op == "/" {
		res = a / b
	}
	return res
}
