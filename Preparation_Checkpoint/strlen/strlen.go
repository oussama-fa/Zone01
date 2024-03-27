package main

import "fmt"

// Level 01

func StrLen(s string) int {
	return len([]rune(s))
}

func main() {
	fmt.Println(StrLen("Hello World!"))
}
