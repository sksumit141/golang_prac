package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("NUmber is:", i)
	}
	counter := 0
	for {
		fmt.Println("Infinite loop")
		counter++
		if counter == 3 {
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("INdex of NUmbers is %d, and value is %d\n", index, value)
	}
}
