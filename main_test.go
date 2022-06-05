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
