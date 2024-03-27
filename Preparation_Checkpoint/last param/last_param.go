package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Level 01

func main() {
	A := os.Args
	if len(A) > 1 {
		B := A[len(A)-1]
		for _, C := range B {
			z01.PrintRune(C)
		}
	}
	z01.PrintRune('\n')
}
