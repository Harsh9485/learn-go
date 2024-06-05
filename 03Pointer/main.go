package main

import "fmt"

func main() {
	fmt.Println("welcome to golang course talk abote Pointers ")

	// var a *int
	// fmt.Println("value for a", a) --> nil
	// fmt.Printf("type %T\n", a) --> *int
	
	// myNumber := 10
	// ref := &myNumber
	// fmt.Println("value for ref", ref) --> memory address of myNumber EX- 0xc0000100a0
	// fmt.Printf("type %T", ref) --> *int this is a pointer to reference to myNumber
	// fmt.Println("value for *ref", *ref) --> 10  references to myNumber value 10

	// *ref = *ref + 10
	// fmt.Println("New value of myNumber", myNumber) --> 20

	// dubRef := &ref
	// fmt.Println("value for ref", *ref) --> 10
	// fmt.Println("value for dubRef", *dubRef) --> memory address of muNumber 
	// fmt.Println("value for dubRef", **dubRef) --> 10 references to myNumber value 10 because this pointer references to ref this is also a pointer to reference to myNumber
}