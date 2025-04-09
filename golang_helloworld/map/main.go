package main

import "fmt"

func main() {
	studentgrade := make(map[string]int)

	studentgrade["Sumit"] = 89
	studentgrade["bob"] = 43
	studentgrade["sam"] = 46
	studentgrade["ban"] = 96

	fmt.Println("Matks of sumit,", studentgrade["Sumit"])

	//checking if a key exists

	Grades, Exists := studentgrade["Sumit"]
	fmt.Printf("grades and boolean %d %t", Grades, Exists)

	for index, value := range studentgrade {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

}
