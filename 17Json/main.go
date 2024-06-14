package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("welcome to golang course talk abote Json data ")
	lenCode := []Course{
		{
			Name: "Golang",
			Price: 100,
			Duration: 2,
			Rating: 5.0,
			Pasword: "cdad41d5sa",
			Flag: []string{"wev-dev", "Backend"},
		},
		{
			Name: "Node js",
			Price: 200,
			Duration: 1,
			Rating: 4.5,
			Pasword: "dfdf4df1d",
			Flag: nil,
		},
	}
	jsonFormat, err := json.MarshalIndent(lenCode, "", "\t")
	errChack(err)
	fmt.Printf("%s\n",jsonFormat)

}

func errChack(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Course struct {
	Name string `json:"coursName"`
	Price int 
	Duration int 
	Rating float32
	Pasword string `json:"-"`
	Flag []string  `json:"tage,omitempty"`
}