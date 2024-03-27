package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	}
	n := 1
	for i := 0; i < power; i++ {
		n *= nb
	}
	return n
}
