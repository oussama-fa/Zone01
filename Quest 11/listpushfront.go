package piscine

func ListPushFront(l *List, data interface{}) {
	tmp := &NodeL{Data: data, Next: l.Head}
	if l.Head == nil {
		l.Head = tmp
		l.Tail = tmp
	} else {
		l.Head = tmp
	}
}
