package main

import "github.com/01-edu/z01"

// Level 01

func main() {
	for z := 'z'; z >= 'a'; z-- {
		if z%2 != 0 {
			tmp := z - 32
			z01.PrintRune(tmp)
		} else {
			z01.PrintRune(z)
		}
	}
	z01.PrintRune('\n')
}
