package main

import (
	"fmt"

	"net/http"
)

func main() {
	URL := "https://google.com"
	fmt.Println("welcome to golang course talk abote Web Request ")
	res, err := http.Get(URL)
	errChack(err)
	defer res.Body.Close() // close the body when we are done with it
	
	
	
}

func errChack(err error) {
	if err!= nil {
		fmt.Println(err)
	}
}