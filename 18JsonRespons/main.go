package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	URL := "http://localhost:3000" // --> any url is works
	fmt.Println("welcome to golang course talk abote send from data ")
	res, _ := http.Get(URL)
	defer res.Body.Close()
	constant, _ := io.ReadAll(res.Body)
	var decodeJson map[string]interface{}
	json.Unmarshal(constant, &decodeJson)
	fmt.Println(decodeJson)
	// for k, v := range decodeJson {
	// 	fmt.Printf("%v: %v type: %t\n", k,v,v)
	// }
}

func errChack(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

