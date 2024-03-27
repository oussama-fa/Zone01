package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := '0'
	for a <= '7' {
		b := a + 1
		for b <= '8' {
			c := b + 1
			for c <= '9' {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a != '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				c++
			}
			b++
		}
		a++
	}
	z01.PrintRune('\n')
}
