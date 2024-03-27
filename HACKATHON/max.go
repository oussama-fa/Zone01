package piscine

func Max(a []int) int {
	max := 0
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
