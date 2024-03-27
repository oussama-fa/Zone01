package main

import (
	"fmt"
)

func AlphaCount(s string) int {
	cnt := 0
	for _, a := range s {
		if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') {
			cnt++
		}
	}
	return cnt
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
