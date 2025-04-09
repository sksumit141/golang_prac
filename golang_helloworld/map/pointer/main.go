package main

import "fmt"

func modifyvalue(num *int) {
	*num = *num * 2
}

func main() {
	value := 10

	modifyvalue(&value)
	fmt.Println("Value is ", value)
}
