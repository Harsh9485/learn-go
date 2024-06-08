package main

import "fmt"

func main() {
	fmt.Println("welcome to golang course talk abote methods")
	harsh := User{
		Name: "Harsh",
		Email: "harsh@gmail",
		Age: 20,
		Status: true,
		
	}
	// fmt.Println(harsh.Status)
	harsh.SetStatus()
	// fmt.Println(harsh.Status)
}

type User struct {
	Name string
	Email string
	Age int
	Status bool
	
}

func (u User) SetStatus() {
	togal := &u.Status
	u.Status = !*togal
	fmt.Printf("New status is %v\n", u.Status)
}