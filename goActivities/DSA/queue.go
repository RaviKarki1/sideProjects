package main

type Queue struct {
	head    *Node
	tail    *Node
	size    int
	maxSize int
}

func NewQueue(maxSize int) *Queue {
	return &Queue{
		head:    nil,
		tail:    nil,
		size:    0,
		maxSize: maxSize,
	}
}

func (q *Queue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *Queue) hasSpace() bool {
	if q.size < q.maxSize {
		return true
	}
	return false
}

func (q *Queue) Peek() string {
	return q.head.GetValue()
}

func (q *Queue) enQueue() {
	if q.hasSpace() {

	}
}
