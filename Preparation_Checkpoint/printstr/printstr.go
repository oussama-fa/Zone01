package main

import "github.com/01-edu/z01"

// Level 01

func PrintStr(s string) {
	for a := 0; a < len(s); a++ {
		z01.PrintRune(rune(s[a]))
	}

}

func main() {
	PrintStr("Hello World!")
}
