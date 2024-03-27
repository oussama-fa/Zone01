package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	cnt := 0
	for l != nil {
		if cnt == pos {
			return l
		}
		cnt++
		l = l.Next
	}
	return nil
}
