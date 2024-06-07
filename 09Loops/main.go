package main

import "fmt"

func main() {
	fmt.Println("welcome to golang course talk abote loops")
	// weekest := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// for i:=0; i < len(weekest); i++ {
	// 	fmt.Println(weekest[i])
	// }
	// for indxe, day := range weekest {
	// 	fmt.Printf("%v %v \n", indxe, day)
	// }

	for i:=1; i < 10; i++ {
		if i == 2 {
			goto hi
		}
	}
	hi:
		fmt.Println("jumping to golang course")

}