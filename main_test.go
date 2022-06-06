package main

import (
	"testing"
)

func TestTest(t *testing.T) {
	Test()
}
func TestTest2(t *testing.T) {
	Test()
}
func TestTest3(t *testing.T) {
	Test()
}

/*
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test()
	}
}
func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test2()
	}
}
func BenchmarkTest3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test3()
	}
}
*/
func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculate("*", 5, 5)
	}
}
func BenchmarkCalculate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate("*", 5, 5)
	}
}

func BenchmarkCalculate3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate2("*", 5, 5)
	}
}
