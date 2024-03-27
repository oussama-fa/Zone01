package piscine

func ToLowers(s string) string {
	r := ""
	for _, a := range s {
		if a >= 'A' && a <= 'Z' {
			r += string(a + 32)
		} else {
			r += string(a)
		}
	}
	return r
}

func Capitalize(s string) string {
	s = ToLowers(s)
	res := []rune(s)

	for i := 0; i < len(res); i++ {
		if i == 0 && (res[i] >= 'a' && res[i] <= 'z') {
			res[i] = rune(res[i] - 32)
		}
		if (res[i] >= 32 && res[i] <= 47) || (res[i] >= 58 && res[i] <= 64) || (res[i] >= 91 && res[i] <= 96) || (res[i] >= 123 && res[i] <= 126) {
			if i+1 < len(res) && res[i+1] >= 'a' && res[i+1] <= 'z' {
				res[i+1] = rune(res[i+1] - 32)
			}
		}
	}
	return string(res)
}
