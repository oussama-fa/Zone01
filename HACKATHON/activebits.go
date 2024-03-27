package piscine

func ActiveBits(n int) int {
	cnt := 0
	for n != 0 {
		if n&1 == 1 {
			cnt++
		}
		n >>= 1
	}
	return cnt
}
