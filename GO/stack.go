package main

import (
	"fmt"
)

type StackNode struct {
	data string
	next *StackNode
}

type Stack struct {
	top  *StackNode
	size int
}

func NewStack() *Stack {
	return &Stack{
		top:  nil,
		size: 0,
	}
}

func (s *Stack) Push(value string) {
	newNode := &StackNode{
		data: value,
		next: s.top,
	}
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() {
	if s.top == nil {
		return
	}
	s.top = s.top.next
	s.size--
}

func (s *Stack) Top() (string, error) {
	if s.top == nil {
		return "", fmt.Errorf("stack is empty")
	}
	return s.top.data, nil
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Clear() {
	for !s.IsEmpty() {
		s.Pop()
	}
}

func (s *Stack) Print() {
	fmt.Print("Stack [ ")
	s.printRecursive(s.top)
	fmt.Println("]")
}

func (s *Stack) printRecursive(node *StackNode) {
	if node == nil {
		return
	}
	s.printRecursive(node.next)
	fmt.Printf("\"%s\" ", node.data)
}

// ToVector - возвращает элементы ОТ ВЕРШИНЫ К ОСНОВАНИЮ
func (s *Stack) ToVector() []string {
	result := make([]string, 0, s.size)
	current := s.top
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

// FromVector - загружает элементы В ОБРАТНОМ ПОРЯДКЕ
// Чтобы восстановить стек, нужно добавлять элементы с конца
func (s *Stack) FromVector(elements []string) {
	s.Clear()
	// Добавляем элементы с конца, чтобы первый элемент стал вершиной
	for i := len(elements) - 1; i >= 0; i-- {
		s.Push(elements[i])
	}
}

// ToVectorReverse - возвращает элементы ОТ ОСНОВАНИЯ К ВЕРШИНЕ
func (s *Stack) ToVectorReverse() []string {
	result := make([]string, 0, s.size)
	s.collectReverse(s.top, &result)
	return result
}

func (s *Stack) collectReverse(node *StackNode, result *[]string) {
	if node == nil {
		return
	}
	s.collectReverse(node.next, result)
	*result = append(*result, node.data)
}

// Equals сравнивает два стека
func (s *Stack) Equals(other *Stack) bool {
	if s.size != other.size {
		return false
	}

	// Сравниваем элементы, двигаясь от вершины
	current1 := s.top
	current2 := other.top

	for current1 != nil && current2 != nil {
		if current1.data != current2.data {
			return false
		}
		current1 = current1.next
		current2 = current2.next
	}

	return true
}

// Copy создает копию стека
func (s *Stack) Copy() *Stack {
	newStack := NewStack()
	elements := s.ToVectorReverse() // берем от основания к вершине
	for i := 0; i < len(elements); i++ {
		newStack.Push(elements[i])
	}
	return newStack
}