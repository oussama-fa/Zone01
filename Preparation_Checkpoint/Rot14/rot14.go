package main

func Rot14(s string) string {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] = ((((str[i] + 14) - 'a') % 26) + 'a')
		} else if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = ((((str[i] + 14) - 'A') % 26) + 'A')
		}
	}
	return string(str)
}
