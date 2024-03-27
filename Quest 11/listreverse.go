package piscine

func ListReverse(l *List) {
	var a *NodeL = nil
	swp := l.Head
	var b *NodeL = nil

	for swp != nil {
		b = swp.Next
		swp.Next = a
		a = swp
		swp = b
	}
	l.Head = a
}
