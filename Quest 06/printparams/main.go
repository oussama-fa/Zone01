package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	tmp := len(arg)
	for i := 1; i < tmp; i++ {
		for _, p := range arg[i] {
			z01.PrintRune(p)
		}
		z01.PrintRune('\n')
	}
}
