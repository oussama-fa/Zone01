package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb < 0 || nb >= 21 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb > 0 {
		for i := 1; i <= nb; i++ {
			n *= i
		}
	}
	return n
}
