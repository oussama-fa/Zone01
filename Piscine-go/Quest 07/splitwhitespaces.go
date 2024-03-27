package piscine

func help(s rune) bool {
	return s == ' ' || s == '\t' || s == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var res []string
	var tmp string

	for _, i := range s {
		if help(i) == true {
			if tmp != "" {
				res = append(res, tmp)
				tmp = ""
			}
		} else {
			tmp += string(i)
		}
	}
	if tmp != "" {
		res = append(res, tmp)
		tmp = ""
	}
	return res
}
