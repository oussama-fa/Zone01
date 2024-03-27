package piscine

func StringToIntSlice(str string) []int {
	var r []int
	for _, i := range str {
		r = append(r, int(i))
	}
	return r
}
