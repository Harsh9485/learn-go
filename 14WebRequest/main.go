package main

import (
	"fmt"
	"io"
	

	"net/http"
)

func main() {
	URL := "https://github.com/" // --> any url is works
	fmt.Println("welcome to golang course talk abote Web Request ")
	res, err := http.Get(URL) // --> send get request to the url
	errChack(err)
	defer res.Body.Close() // close the body when we are done with it
	contant, _ := io.ReadAll(res.Body) // read the data from the body and convert []byte 
	fmt.Print(string(contant)) // convert the string and print it
	
}

func errChack(err error) {
	if err!= nil {
		fmt.Println(err)
	}
}