package piscine

func splt(str string) []string {
	m := []string{}
	a := ""
	for _, i := range str {
		if i != ' ' {
			a = a + string(i)
		} else {
			m = append(m, a)
			a = ""
		}
	}

	if a != " " {
		m = append(m, a)
	}
	return m
}

func ShoppingSummaryCounter(str string) map[string]int {
	tmp := splt(str)
	cnt := make(map[string]int)
	for _, i := range tmp {
		cnt[i]++
	}
	return cnt
}
