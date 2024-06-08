package main

import (
	"fmt"
	
	"strconv"
)

func main() {
	fmt.Println("welcome to golang course talk abote functions")

	printName("Golang")

	sum := addTwoNum(10, 20)
	fmt.Printf("sum of 10 and 20 is %T \n", sum)
	sum2 := proAdder(10, 20, 30, 40, 50)
	fmt.Println(sum2)
}
func printName(name string) {
	fmt.Println(name)
}
func addTwoNum(val int, val2 int) string{
	sum := val + val2
	return strconv.Itoa(sum)
}

func proAdder(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}