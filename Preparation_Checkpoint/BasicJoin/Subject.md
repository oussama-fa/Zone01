# basicjoin
Instructions

Write a function that returns a concatenated string from the ‘strings’ passed as arguments.
Expected function

func BasicJoin(elems []string) string {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.BasicJoin(elems))
}

And its output :

$ go run .
Hello! How are you?
$

