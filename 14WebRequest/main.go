package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	URL := "https://jsonplaceholder.typicode.com/todos/1" // --> any url is works
	fmt.Println("welcome to golang course talk abote Web Request ")
	res, err := http.Get(URL) // --> send get request to the url
	errChack(err)
	defer res.Body.Close() // close the body when we are done with it
	contant, _ := ioutil.ReadAll(res.Body) // read the data from the body and convert []byte 
	var data interface{}
	json.Unmarshal(contant, &data) // convert
	fmt.Print(data) // convert the string and print it
	
}

type data struct {}

func errChack(err error) {
	if err!= nil {
		fmt.Println(err)
	}
}