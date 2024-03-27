package piscine

func IsLower(s string) bool {
	a := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			a += 1
		}
	}
	if a == len(s) {
		return true
	} else {
		return false
	}
}
