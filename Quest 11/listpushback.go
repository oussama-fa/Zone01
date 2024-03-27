package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	tmp := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = tmp
	} else {
		l.Tail.Next = tmp
	}
	l.Tail = tmp
}
