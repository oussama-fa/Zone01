package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	l := 0
	for _, i := range s {
		l = l*10 + int(i-48)
	}
	return l
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
	fmt.Println(BasicAtoi("Hello World!"))
}
