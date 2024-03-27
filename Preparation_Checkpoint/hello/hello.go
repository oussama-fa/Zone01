package main

import "github.com/01-edu/z01"

func Str(s string) {
	for a := 0; a < len(s); a++ {
		z01.PrintRune(rune(s[a]))
	}
}

func main() {
	Str("hello world!\n")
}
