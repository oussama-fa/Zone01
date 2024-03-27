package main

import (
	"fmt"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}
func BasicJoin(elems []string) string {
	for _, v := range elems {
		print(v)
	}
	return ""
}
