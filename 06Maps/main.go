package main

import "fmt"

func main() {
	fmt.Println("welcome to golang course talk abote Maps")

	languages := make(map[string]string)
	languages["Go"] = "Golang"
	languages["c"] = "C"
	languages["py"] = "python"
	languages["cpp"] = "c++"
	languages["java"] = "java"
	languages["js"] = "javascript"

	for key, value := range languages {
		fmt.Printf("The file exetantion is %v than languages is %v\n", key, value)
	}
}