package main

import (
	"fmt"
	"os"
)

// Level 01

func main() {
	if len(os.Args) == 3 {
		if Wdmatch(os.Args[1], os.Args[2]) {
			fmt.Println(os.Args[1])
		}
	}
}

func Wdmatch(s1 string, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	var rest string
	cnt := 0
	for i := 0; i < len(r1); i++ {
		for j := cnt; j < len(r2); j++ {
			if r1[i] == r2[j] {
				rest += string(r1[i])
				j = len(r2) - 1
			}
			cnt++
		}
	}
	return s1 == rest
}
