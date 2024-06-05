package main

import "fmt"

func main() {

	fmt.Println("welcome to golang course talk abote Slices")

	// var fruitList = []string{"apple", "banana", "orange", "mango", "watermelon"}
	// var fruitList2 = [5]string{}
	// fmt.Printf("type %T \n", fruitList) --> []string this is a slice becose it is dinamice array and not fixed array
	// fmt.Printf("type %T", fruitList2) --> [5]string this is a array becose it is size of array not insart fruitList2[5] = "orange" this will give error becose array is fixed is no incresing size

	// var a = append(fruitList, "pineapple") || append is a function which is used to add element to slice and return new slice
	// fmt.Println(a) --> [apple banana orange mango watermelon pineapple]
	// fmt.Println(fruitList) --> [apple banana orange mango watermelon]

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList) --> [banana orange mango watermelon]

	// removed value
	// index := 2
	// fruitList = append(fruitList[:index], fruitList[index+1:]...) --> fruitlist[:index] = ["apple", "banana"] + [ "mango", "watermalon"] || "orange" is sliced this slice for this logic
	// fmt.Println(fruitList) 

}