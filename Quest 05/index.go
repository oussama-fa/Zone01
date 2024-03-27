package piscine

func Index(s string, toFind string) int {
	index := 0
	s1 := []rune(s)
	s2 := []rune(toFind)
	if len(s) == 0 || len(toFind) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s1[i] == s2[0] {
			index = i
			break
		} else {
			index = -1
		}
	}
	return index
}
