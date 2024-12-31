package main

import "fmt"

type Stack struct {
	head    *Node
	size    int
	maxSize int
}

func newStack(mSize int) *Stack {
	return &Stack{
		head:    nil,
		size:    0,
		maxSize: mSize,
	}
}

func (s *Stack) Peek() string {
	return s.head.GetValue()
}

func (s *Stack) IsEmpty() bool {
	return s.size < 1
}

func (s *Stack) IsFull() bool {
	return s.size < s.maxSize
}

func (s *Stack) Push(value string) {
	node_to_add := NewNode(value, nil)
	if !s.IsFull() {
		fmt.Println("Sorry! the stack is full.")
		return
	}

	if s.IsEmpty() {
		s.head = node_to_add
	} else {
		node_to_add.SetLinkedNode(s.head)
		s.head = node_to_add
	}
	s.size += 1

}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("Sorry the stack is empty.")
	} else {
		if s.size == 1 {
			s.size -= 1
			s.head = nil
		} else {
			s.head = s.head.GetLinkedNode()
			s.size -= 1
		}
	}
}

func (s *Stack) PrintStack() {
	if s.IsEmpty() {
		fmt.Println("Sorry! the stack is empty.")
		return
	} else {
		current := s.head
		fmt.Println("The current stack is:")
		for current != nil {
			fmt.Println(current.GetValue())
			fmt.Println("_________")
			current = current.GetLinkedNode()
		}
	}
}

func main() {
	s1 := newStack(3)
	s1.Push("v3")
	s1.Push("v2")
	s1.Push("v1")

	s1.PrintStack()

	s1.Pop()
	s1.Pop()
	s1.Pop()

	s1.PrintStack()
}
