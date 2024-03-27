package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		PrintNbr(-922337203685477580)
		z01.PrintRune('8')
		return
	} else if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	if n >= 10 {
		PrintNbr(n / 10)
		z01.PrintRune(rune(n%10 + '0'))
	} else if n <= 9 {
		z01.PrintRune(rune(n%10 + '0'))
	}
}
