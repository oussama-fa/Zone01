package piscine

func Join(strs []string, sep string) string {
	str := ""
	i := 0
	for _, j := range strs {
		str += j
		i++
		if len(strs) > i {
			str += sep
		}
	}
	return str
}
