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

// NewListItem constructor for ListItem entity.
func NewListItem(value interface{}, prev *ListItem, next *ListItem) ListItem {
	return ListItem{
		Value: value,
		Prev:  prev,
		Next:  next,
	}
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func NewList() List {
	return &list{
		len:  0,
		head: nil,
		tail: nil,
	}
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
	l.len++
	newItem := NewListItem(v, nil, l.head)

	if l.head == nil {
		l.tail = &newItem
	} else {
		l.head.Prev = &newItem
	}

	l.head = &newItem
	return &newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.len++
	newItem := NewListItem(v, l.tail, nil)

	if l.tail == nil {
		l.head = &newItem
	} else {
		l.tail.Next = &newItem
	}

	l.tail = &newItem
	return &newItem
}

func (l *list) Remove(i *ListItem) {
}

func (l *list) MoveToFront(i *ListItem) {
}
