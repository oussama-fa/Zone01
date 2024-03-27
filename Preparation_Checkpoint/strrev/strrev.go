package main

import (
	"fmt"
)

// Level 01

func StrRev(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
