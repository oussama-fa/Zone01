package piscine

func JumpOver(str string) string {
	r := ""
	if len(str) >= 3 {
		for i, j := range str {
			if i%3 == 2 {
				r += string(j)
			}
		}
	}
	return r + "\n"
}
