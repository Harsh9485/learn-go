package main

import "fmt"

func main() {
	fmt.Println("welcome to golang course talk abote Defers")

	fmt.Println("connecting to database")
	fmt.Print("opration in database")
	defer fmt.Println("disconnecting from database")
}