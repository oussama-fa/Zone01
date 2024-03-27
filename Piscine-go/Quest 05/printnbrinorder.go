package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var arr []int
		a := 0

		b := 0
		var tmp int
		for n != 0 {
			a = n % 10
			n /= 10
			arr = append(arr, a)
		}

		for i := range arr {
			b = i + 1
		}
		for i := 0; i < b-1; i++ {
			for j := 0; j < b-i-1; j++ {
				if arr[j] > arr[j+1] {
					tmp = arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = tmp
				}
			}
		}
		for i := 0; i < b; i++ {
			z01.PrintRune(rune(arr[i] + 48))
		}
	}
}
