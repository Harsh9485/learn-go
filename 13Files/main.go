package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to golang course talk abote files ")

	 writeFile("./newfile.txt", "this is a new file")
	readfile("./newfile.txt")
}

func writeFile(filename string, data string) *os.File {
	file, err := os.Create(filename)
	if err!= nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.WriteString(file, data)
	return file
}
func readfile(fileName string)  {
	databyte, err := ioutil.ReadFile(fileName)
	if err!= nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}