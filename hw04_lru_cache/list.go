package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newHeadNode := &ListItem{Value: v}
	if l.head == nil {
		l.tail = newHeadNode
	} else {
		newHeadNode.Next = l.head
		l.head.Prev = newHeadNode
	}
	l.head = newHeadNode
	l.len++

	return newHeadNode
}

func (l *list) PushBack(v interface{}) *ListItem {
	newTailNode := &ListItem{Value: v}
	if l.head == nil {
		l.head = newTailNode
	} else {
		newTailNode.Prev = l.tail
		l.tail.Next = newTailNode
	}
	l.tail = newTailNode
	l.len++

	return newTailNode
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.head = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		l.tail = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.Len() > 1 {
		l.Remove(i)
		l.PushFront(i.Value)
	}
}

func NewList() List {
	return new(list)
}
