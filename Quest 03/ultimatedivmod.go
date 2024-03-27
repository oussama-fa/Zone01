package piscine

func UltimateDivMod(a *int, b *int) {
	var c int
	c = *a
	*a = c / *b
	*b = c % *b
}
