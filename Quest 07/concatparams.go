package piscine

func ConcatParams(args []string) string {
	var rslt string
	for i := 0; i < len(args); i++ {
		rslt = rslt + (args[i])
		if i < len(args)-1 {
			rslt += "\n"
		}
	}
	return rslt
}
