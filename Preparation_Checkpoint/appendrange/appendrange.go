package main

import (
	"fmt"
)

func AppendRange(min, max int) []int {
	tmp := []int{}
	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		tmp = append(tmp, i)
	}
	return tmp
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
