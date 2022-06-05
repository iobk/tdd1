package main

import (
	"fmt"
)

var opMap map[string]func(int, int) int

func initOpMap() {
	opMap = make(map[string]func(int, int) int)
	opMap["+"] = add
	opMap["-"] = sub
	opMap["*"] = mul
	opMap["/"] = div

}
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}

func main() {
	fmt.Println("Test driven development")
	initOpMap()
	Test()
	Test2()
	Test3()
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
func Test3() {
	if !testCalculate2("test1", "+", 1, 1, 2) {
		return
	}
	if !testCalculate2("test2", "+", 5, 1, 6) {
		return
	}
	if !testCalculate2("test3", "+", 5, -1, 4) {
		return
	}
	if !testCalculate2("test4", "-", 5, 1, 4) {
		return
	}
	if !testCalculate2("test5", "-", 5, 5, 0) {
		return
	}
	if !testCalculate2("test6", "*", 5, 5, 25) {
		return
	}
	if !testCalculate2("test7", "*", 5, -5, -25) {
		return
	}
	if !testCalculate2("test8", "/", 5, 5, 1) {
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
func testCalculate2(testcase, op string, a, b, expected int) bool {
	o := Calculate2(op, a, b)
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
func Calculate2(op string, a, b int) int {
	if v, ok := opMap[op]; ok {
		return v(a, b)
	}
	return 0

}
