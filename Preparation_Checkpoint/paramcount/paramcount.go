package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	A := os.Args
	fmt.Print(len(A))

	z01.PrintRune('\n')
}
