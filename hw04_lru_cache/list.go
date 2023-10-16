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

// NewListItem constructor for ListItem entity
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
	return nil
}

func (l *list) Back() *ListItem {
	return nil
}

func (l *list) PushFront(v interface{}) *ListItem {
	return nil
}

func (l *list) PushBack(v interface{}) *ListItem {
	return nil
}

func (l *list) Remove(i *ListItem) {
}

func (l *list) MoveToFront(i *ListItem) {
}
