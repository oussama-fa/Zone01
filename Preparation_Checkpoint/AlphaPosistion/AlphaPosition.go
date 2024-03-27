package main

import "fmt"

func AlphaPosition(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	} else if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 1)
	}
	return -1
}

func main() {
	fmt.Println(AlphaPosition('a'))
	fmt.Println(AlphaPosition('z'))
	fmt.Println(AlphaPosition('B'))
	fmt.Println(AlphaPosition('Z'))
	fmt.Println(AlphaPosition('0'))
	fmt.Println(AlphaPosition(' '))
}
