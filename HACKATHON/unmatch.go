package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if i != j && a[i] == a[j] {
				a[i] = -1
				a[j] = -1
			}
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] != -1 {
			return a[i]
		}
	}
	return -1
}
