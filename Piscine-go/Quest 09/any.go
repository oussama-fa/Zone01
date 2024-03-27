package piscine

func Any(f func(string) bool, a []string) bool {
	for _, j := range a {
		if f(j) {
			return true
		}
	}
	return false
}
