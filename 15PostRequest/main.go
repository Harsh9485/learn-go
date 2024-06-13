package main

import (
	"fmt"
	"io"
	
	"net/http"
	"strings"
)

func main() {
	URL := "http://localhost:3000" // --> any url is works
	fmt.Println("welcome to golang course talk abote POST Request ")
	reqBody := strings.NewReader(`{
		"massge" : "hi there is something",
		"time" : 25,
		"login": true
	}`)
	res, err := http.Post(URL, "application/json", reqBody) // --> send get request to the url
	errChack(err)
	defer res.Body.Close()             // close the body when we are done with it
	contant, _ := io.ReadAll(res.Body) // read the data from the body and convert []byte
	//  string to struct

	fmt.Println(string(contant)) // convert the string and print it

}

func errChack(err error) {
	if err!= nil {
		fmt.Println(err)
	}
}