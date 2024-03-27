package piscine

func space(s string) bool {
	ary := []rune(s)
	for _, i := range ary {
		if i != ' ' {
			return false
		}
	}
	return true
}

func LoafOfBread(str string) string {
	if len(str) > 1 && AlphaCount(str) < 5 {
		if space(str) {
			return "\n"
		}
		return "Invalid Output\n"
	}
	if str == "" {
		return "\n"
	}
	ary := []rune(str)
	reslt := []string{}
	res := ""
	tmp := ""
	for _, i := range ary {
		if len(tmp) < 5 {
			if i != ' ' {
				tmp += string(i)
			}
		} else {
			res += tmp
			tmp = ""
		}
	}
	res += tmp
	s := len(res)
	for i := 0; i+5 <= s; i += 5 {
		reslt = append(reslt, res[i:i+5])
	}
	reslt = append(reslt, res[s/5*5:])
	result := Join(reslt, " ")
	ary = []rune(result)
	if ary[len(ary)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result + "\n"
}
