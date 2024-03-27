package piscine

func TrimAtoi(s string) int {
	var r []rune
	l := 0

	for i, let := range s {
		if let >= '0' && let <= '9' {
			r = append(r, let)
			l = i + 1
		} else if let == '-' && l == 0 {
			r = append(r, '-')
		} else {
			continue
		}
	}
	return atoi(r)
}

func atoi(r []rune) int {
	a := 0
	res := 0
	for i := range r {
		i++
		a++
	}
	if a == 0 {
		return 0
	}

	x := 1
	for k := 0; k < a-1; k++ {
		if r[k] == '+' || r[k] == '-' {
			continue
		}
		x *= 10
	}

	for i := range r {
		if (r[0] == '-' || r[0] == '+') && i == 0 {
			continue
		}
		if r[i] < '0' || r[i] > '9' {
			return 0
		}
		n := 0
		for j := '0'; j < r[i]; j++ {
			n++
		}
		res += n * x
		x /= 10
	}
	if r[0] == '-' {
		res *= -1
	}
	return res
}
