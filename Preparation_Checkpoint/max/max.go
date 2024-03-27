package main

import (
	"fmt"
	"piscine-go"
)

// Level 01

func Max(a []int) int {
	res := 0
	for _, nb := range a {
		if nb > res {
			res = nb
		}
	}
	return res

}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := piscine.Max(a)
	fmt.Println(max)
}
