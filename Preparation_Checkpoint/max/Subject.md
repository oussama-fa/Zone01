# max

Instructions

Write a function Max that will return the maximum value in a slice of integers. If the slice is empty it will return 0.
Expected function

func Max(a []int) int {

}

Usage

Here is a possible program to test your function :

package main

import (
"fmt"
"piscine"
)

func main() {
a := []int{23, 123, 1, 11, 55, 93}
max := piscine.Max(a)
fmt.Println(max)
}

And its output :

$ go run .
123
$
