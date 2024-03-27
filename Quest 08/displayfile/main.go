package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fnlength := len(os.Args[1:])
	if fnlength == 0 {
		fmt.Println("File name missing")
	} else if fnlength == 1 {
		ref, _ := ioutil.ReadFile("quest8.txt")
		fmt.Print(string(ref))
	} else if fnlength > 1 {
		fmt.Println("Too many arguments")
	}
}
