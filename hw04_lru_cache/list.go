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
	FirstItem *ListItem
	LastItem  *ListItem
	size      int
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.FirstItem
}

func (l *list) Back() *ListItem {
	return l.LastItem
}

func (l *list) PushFront(v interface{}) *ListItem {

	li := ListItem{Value: v}

	if l.FirstItem != nil {
		li.Next = l.FirstItem
		l.FirstItem.Prev = &li
	}

	if l.FirstItem == nil && l.LastItem == nil {
		l.LastItem = &li
	}

	l.FirstItem = &li
	l.size++
	return &li
}

func (l *list) PushBack(v interface{}) *ListItem {

	li := ListItem{Value: v}

	if l.LastItem != nil {
		li.Prev = l.LastItem
		l.LastItem.Next = &li
	}

	if l.FirstItem == nil && l.LastItem == nil {
		l.FirstItem = &li
	}

	l.LastItem = &li
	l.size++
	return &li
}

func (l *list) Remove(i *ListItem) {
	l.size--

	if l.FirstItem == i && l.LastItem == i {
		l.FirstItem = nil
		l.LastItem = nil
		return
	}

	if l.FirstItem == i {
		l.FirstItem = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if l.LastItem == i {
		l.LastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	i.Prev = nil
	i.Next = nil
}

func (l *list) MoveToFront(i *ListItem) {
	if l.FirstItem == i {
		return
	}
	if l.LastItem == i {
		i.Prev.Next = nil
		l.LastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	i.Prev.Next = i.Next
	i.Prev = nil
	i.Next = l.FirstItem
	l.FirstItem.Prev = i
	l.FirstItem = i
}

func NewList() List {
	return new(list)
}
