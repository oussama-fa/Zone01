package piscine

func Map(f func(int) bool, arr []int) []bool {
	tmp := make([]bool, len(arr))

	for i, j := range arr {
		tmp[i] = f(j)
	}
	return tmp
}
