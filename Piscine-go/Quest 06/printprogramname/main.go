package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	if len(args) >= 2 && args[0] == '.' && args[1] == '/' {
		args = args[2:]
	}
	for _, me := range args {
		z01.PrintRune(me)
	}
	z01.PrintRune('\n')
}
