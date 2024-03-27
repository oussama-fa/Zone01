package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for a := 'a'; a <= 'z'; a++ {
		if a%2 == 0 {
			b := a - 32
			z01.PrintRune(b)
		} else {
			z01.PrintRune(a)
		}
	}
	z01.PrintRune('\n')
}
