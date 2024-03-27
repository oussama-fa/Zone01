package main

import "fmt"

func AddFront(s string, slice []string) []string {
	if len(s) == 0 {
		return slice
	}
	return append([]string{s}, slice...)
}

func main() {
	fmt.Println(AddFront("Hello", []string{"world"}))
	fmt.Println(AddFront("Hello", []string{"world", "!"}))
	fmt.Println(AddFront("Hello", []string{}))
}
