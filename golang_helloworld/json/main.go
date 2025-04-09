package main

import "fmt"

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("there is an error", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.SatusOK {
		fmt.Println("error in grtting response", res.Status)
		return
	}
	
}
