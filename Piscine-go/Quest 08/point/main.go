package main

import "github.com/01-edu/z01"

var str = "x = 42, y = 21"

func main() {
	for _, a := range str { // PrintSTR
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}
