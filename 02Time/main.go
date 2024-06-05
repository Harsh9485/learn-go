package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to golang course toke abote Time package")
	prsantTIme := time.Now()
	fmt.Println(prsantTIme.Format("2006/01/02"))
	fmt.Printf("type %T", prsantTIme)
	time.Sleep(time.Second * 10)
}