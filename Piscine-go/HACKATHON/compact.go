package piscine

func Compact(ptr *[]string) int {
	Slc := []string{}
	for _, i := range *ptr {
		if i != "" {
			Slc = append(Slc, i)
		}
	}
	*ptr = Slc
	return len(Slc)
}
