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
	fmt.Println("main ",harsh.Name) // --> Harsh
	harsh.ChangeVal()
	fmt.Println("main ",harsh.Name) // --> veer
}
	
type User struct {
	Name string
	Email string
	Age int
	Status bool
}
		
func (u *User) ChangeVal() { // change value of status using Pointer
	u.Name = "veer"
	fmt.Println("ChangeVal :-",u.Name)
}