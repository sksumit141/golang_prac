package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//file, err := os.Create("example.txt")
	//if err != nil {
	//fmt.Println("Error while creating file", err)
	//		return
	//	}
	//	defer file.Close()

	//	content := "hello to my golang file"

	//	byte, errors := io.WriteString(file, content)
	//	fmt.Println("total bytes,", byte)
	//	if errors != nil {
	//		fmt.Println("Error while creating file", errors)
	//		return
	//	}
	//
	//	fmt.Println("file creates sucessfully")
	/*
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error in opening file", err)
			return
		}
		defer file.Close()

		buffer := make([]byte, 1024)

		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error in printing", err)
				return
			}

			fmt.Println(string(buffer[:n]))
		}
	*/

	content, err := ioutil.ReadFile("example.txt")

	if err != nil {
		fmt.Println("Error while reading", err)
		return
	}
	fmt.Println(string(content))
}
