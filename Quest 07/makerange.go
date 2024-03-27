package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	res := make([]int, max-min)

	for i := 0; i < max-min; i++ {
		res[i] = i + min
	}

	return res
}
