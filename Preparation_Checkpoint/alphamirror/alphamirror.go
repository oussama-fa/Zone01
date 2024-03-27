package main

import (
	"os"

	"github.com/01-edu/z01"
)

func miror(c rune) {
	if c >= 'a' && c <= 'z' {
		c = 'z' - (c - 'a')
	} else if c >= 'A' && c <= 'Z' {
		c = 'Z' - (c - 'A')
	}
	z01.PrintRune(c)
}

func main() {
	if len(os.Args) >= 2 {
		for _, c := range os.Args[1] {
			miror(c)
		}
	}
	z01.PrintRune('\n')
}
