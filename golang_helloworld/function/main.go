package main

import "fmt"

func simpleFunction() {
	fmt.Println("Hello")
}
func add(a, b int) int {
	return a + b
}
func multi(a, b int) int {
	return a * b
}
func main() {
	simpleFunction()

	ans := add(2, 4)
	fmt.Println("add", ans)

	multiply := multi(2, 2)

	fmt.Println("result", multiply)
}
