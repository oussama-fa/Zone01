package picsine

func Atoi(s string) int {
	res := 0
	sig := 1
	if len(s) < 0 {
		return 0
	}
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sig = -1
		}
		s = s[1:]
	}
	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		}
		res = res*10 + int(i-48)
	}
	return res * sig
}
