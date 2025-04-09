package main

import "strconv"

func main() {

	//var num int = 20

	//var data float64 = float64(num)

	num := 123

	str := strconv.Itoa(num)

	number_string := "2242"

	number_int, _ := strconv.Atoi(number_string)
}

//defer keyword is used to excute code at last of main function
