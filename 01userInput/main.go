package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	reder := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your username")
	input , _ := reder.ReadString('\n')
	fmt.Println("Enter your password: ", input)
	num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	fmt.Printf("type %T", num)
	if err != nil {
		fmt.Println(err)
	}
}