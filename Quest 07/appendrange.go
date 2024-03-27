package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		var t []int = nil
		return t
	} else {
		t := []int{}

		for i := min; i < max; i++ {
			t = append(t, i)
		}
		return t
	}
}
