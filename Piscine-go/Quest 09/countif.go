package piscine

func CountIf(f func(string) bool, tab []string) int {
	cnt := 0
	for _, i := range tab {
		if f(i) {
			cnt++
		}
	}
	return cnt
}
