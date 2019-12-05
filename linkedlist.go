package practice

// Can be used as a stupid array,
// a queue, or a stack.

type DoublyLinkedListNode struct {
	value int
	prev  *DoublyLinkedListNode
	next  *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	Size  int
	first *DoublyLinkedListNode
	last  *DoublyLinkedListNode
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return l.Size == 0
}

func (l *DoublyLinkedList) Add(value int) {
	l.AddLast(value)
}

// Push
func (l *DoublyLinkedList) AddLast(value int) {
	defer func() { l.Size++ }()
	n := DoublyLinkedListNode{value, nil, nil}
	if l.IsEmpty() {
		l.first = &n
		l.last = &n
		return
	}

	l.last.next = &n
	n.prev = l.last
	l.last = &n
}

// Shift
func (l *DoublyLinkedList) AddFirst(value int) {
	defer func() { l.Size++ }()
	n := DoublyLinkedListNode{value, nil, nil}
	if l.IsEmpty() {
		l.first = &n
		l.last = &n
		return
	}

	l.first.prev = &n
	n.next = l.first
	l.first = &n
}

func (l *DoublyLinkedList) Remove() (int, bool) {
	v, ok := l.RemoveFirst()
	return v, ok
}

// Pop
func (l *DoublyLinkedList) RemoveLast() (int, bool) {
	if l.IsEmpty() {
		return 0, false
	}
	defer func() { l.Size-- }()

	n := l.last
	newLast := l.last.prev

	n.prev = nil
	if newLast != nil {
		newLast.next = nil
	}

	l.last = newLast

	return n.value, true
}

// Unshift
func (l *DoublyLinkedList) RemoveFirst() (int, bool) {
	if l.IsEmpty() {
		return 0, false
	}
	defer func() { l.Size-- }()

	n := l.first
	newFirst := l.first.next

	n.next = nil
	if newFirst != nil {
		newFirst.prev = nil
	}

	l.first = newFirst

	return n.value, true
}
