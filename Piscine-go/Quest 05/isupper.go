package piscine

func IsUpper(s string) bool {
	A := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			A += 1
		}
	}
	if A == len(s) {
		return true
	} else {
		return false
	}
}
