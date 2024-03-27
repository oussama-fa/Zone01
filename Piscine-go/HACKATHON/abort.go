package piscine

func sort(tab []int) {
	for i := 0; i < len(tab); i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] > tab[j] {
				swp := tab[i]
				tab[i] = tab[j]
				tab[j] = swp
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	sort(tab)
	return tab[2]
}
