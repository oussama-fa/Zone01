package main

import "github.com/01-edu/z01"

func main() {
	for z := 'z'; z >= 'a'; z-- {
		z01.PrintRune(z)
	}
	z01.PrintRune('\n')
}
