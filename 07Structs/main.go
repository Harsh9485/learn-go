package main

import "fmt"

func main() {
	fmt.Println("welcome to golang course talk abote Structs")
	// struct is like a class in java and ander langauges
	// no inharitance in go and no super or parent
	harsh := User{
		Name: "Harsh",
		Email: "harsh@gmail",
		Age: 20,
		Status: true,
		
	}
	fmt.Println(harsh.Age)
}

type User struct {
	Name string
	Email string
	Age int
	Status bool
	
}