package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb >= 21 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb > 0 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}
