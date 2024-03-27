package piscine

func DescendAppendRange(max, min int) []int {
	r := []int{}
	for i := max; i > min; i-- {
		r = append(r, i)
	}
	return r
}
