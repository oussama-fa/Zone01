# count-character
Instructions

write a function that takes a string and a character as arguments and returns the number of times the character appears in the string.

    if the character is not in the string return 0
    if the string is empty return 0

Expected Function

func CountChar(str string, c rune) int {
    ...
}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}

And its output :

$ go run .
3
0
1
1

