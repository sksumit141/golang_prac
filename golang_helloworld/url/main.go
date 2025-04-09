package main

import (
	"fmt"
	"net/url"
)

func main() {

	myUrl := ""

	parsedURL, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("Scheme", parsedURL.Scheme)
	// similar function is HOst Path RAwQuer

	//Modify url componenets

	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=sumit"

	newUrl := parsedURL.String()

	fmt.Println("new url", newUrl)
}
