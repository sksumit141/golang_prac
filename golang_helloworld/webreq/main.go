package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")
	//just add link
	res, err := http.Get("")

	if err != nil {
		fmt.Println("Error in receiving data", err)
		return
	}

	defer res.Body.Close()
	//fmt.Println("content is", res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("the errror is", err)
		return
	}
	fmt.Println("response", string(data))
}
