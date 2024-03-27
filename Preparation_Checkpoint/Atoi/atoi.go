package main

import "fmt"

func Atoi(s string) int {
	res := 0
	sig := 1
	if len(s) < 0 {
		return 0
	}
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sig = -1
		}
		s = s[1:]
	}
	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		}
		res = res*10 + int(i-48)
	}
	return res * sig
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
