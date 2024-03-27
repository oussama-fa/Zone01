package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ag := os.Args
	l := len(ag)
	for i := 1; i < l-1; i++ {
		for j := 1; j < l-1; j++ {
			if ag[j] >= ag[j+1] {
				ag[j], ag[j+1] = ag[j+1], ag[j]
			}
		}
	}
	for i := 1; i < l; i++ {
		for _, p := range ag[i] {
			z01.PrintRune(p)
		}
		z01.PrintRune('\n')
	}
}
