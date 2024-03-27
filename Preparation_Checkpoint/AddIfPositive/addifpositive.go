package main

import "fmt"

func AddIfPositive(a int, b int) int {
	if a >= 0 && b >= 0 {
		return a + b
	}
	return 0
}

func main() {
	fmt.Println(AddIfPositive(1, 2))
	fmt.Println(AddIfPositive(1, -2))
	fmt.Println(AddIfPositive(-1, 2))
	fmt.Println(AddIfPositive(-1, -2))
	fmt.Println(AddIfPositive(10, 20))
	fmt.Println(AddIfPositive(0, 20))
}
