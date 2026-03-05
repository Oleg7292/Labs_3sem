package main

import "fmt"

type SLLNode struct {
	data string
	next *SLLNode
}

type SinglyLinkedList struct {
	head *SLLNode
	tail *SLLNode
	size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (s *SinglyLinkedList) PushBack(value string) {
	newNode := &SLLNode{data: value}
	if s.head == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		s.tail.next = newNode
		s.tail = newNode
	}
	s.size++
}

func (s *SinglyLinkedList) PushFront(value string) {
	newNode := &SLLNode{data: value}
	if s.head == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
	s.size++
}

func (s *SinglyLinkedList) PopFront() {
	if s.head == nil {
		return
	}
	s.head = s.head.next
	s.size--
	if s.head == nil {
		s.tail = nil
	}
}

func (s *SinglyLinkedList) Insert(index int, value string) {
	if index > s.size {
		return
	}
	if index == 0 {
		s.PushFront(value)
		return
	}
	if index == s.size {
		s.PushBack(value)
		return
	}
	
	current := s.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	
	newNode := &SLLNode{data: value}
	newNode.next = current.next
	current.next = newNode
	s.size++
}

func (s *SinglyLinkedList) Remove(index int) {
	if index >= s.size {
		return
	}
	if index == 0 {
		s.PopFront()
		return
	}
	
	current := s.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	
	current.next = current.next.next
	if current.next == nil {
		s.tail = current
	}
	s.size--
}

func (s *SinglyLinkedList) Get(index int) (string, error) {
	if index >= s.size {
		return "", fmt.Errorf("index out of range")
	}
	
	current := s.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data, nil
}

func (s *SinglyLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *SinglyLinkedList) Size() int {
	return s.size
}

func (s *SinglyLinkedList) Clear() {
	for !s.IsEmpty() {
		s.PopFront()
	}
}

func (s *SinglyLinkedList) ToVector() []string {
	result := make([]string, 0, s.size)
	current := s.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func (s *SinglyLinkedList) FromVector(elements []string) {
	s.Clear()
	for _, element := range elements {
		s.PushBack(element)
	}
}
