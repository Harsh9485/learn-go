package main

import (
	"fmt"
	"io"

	"net/http"
	"net/url"
	
)

func main() {
	URL := "http://localhost:3000" // --> any url is works
	fmt.Println("welcome to golang course talk abote send from data ")
	data := url.Values{}
	data.Add("name", "ahmed")
	data.Add("age", "25")
	data.Add("login", "true")
	res, err := http.PostForm(URL, data) // --> send get request to the url
	errChack(err)
	defer res.Body.Close()             // close the body when we are done with it
	contant, err := io.ReadAll(res.Body) // read the data from the body and convert []byte
	errChack(err)
	fmt.Println(string(contant)) // convert the string and print it

}

func errChack(err error) {
	if err != nil {
		fmt.Println(err)
	}
}