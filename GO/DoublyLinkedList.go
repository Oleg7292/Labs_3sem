package main

import "fmt"

type DLLNode struct {
	data string
	prev *DLLNode
	next *DLLNode
}

type DoublyLinkedList struct {
	head *DLLNode
	tail *DLLNode
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (d *DoublyLinkedList) PushBack(value string) {
	newNode := &DLLNode{data: value}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	}
	d.size++
}

func (d *DoublyLinkedList) PushFront(value string) {
	newNode := &DLLNode{data: value}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.size++
}

func (d *DoublyLinkedList) PopFront() {
	if d.head == nil {
		return
	}
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}
	d.size--
}

func (d *DoublyLinkedList) PopBack() {
	if d.tail == nil {
		return
	}
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.size--
}

func (d *DoublyLinkedList) Insert(index int, value string) {
	if index > d.size {
		return
	}
	if index == 0 {
		d.PushFront(value)
		return
	}
	if index == d.size {
		d.PushBack(value)
		return
	}
	
	current := d.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	
	newNode := &DLLNode{data: value}
	newNode.prev = current.prev
	newNode.next = current
	current.prev.next = newNode
	current.prev = newNode
	d.size++
}

func (d *DoublyLinkedList) Remove(index int) {
	if index >= d.size {
		return
	}
	if index == 0 {
		d.PopFront()
		return
	}
	if index == d.size-1 {
		d.PopBack()
		return
	}
	
	current := d.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	
	current.prev.next = current.next
	current.next.prev = current.prev
	d.size--
}

func (d *DoublyLinkedList) Get(index int) (string, error) {
	if index >= d.size {
		return "", fmt.Errorf("index out of range")
	}
	
	current := d.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data, nil
}

func (d *DoublyLinkedList) GetFront() (string, error) {
	if d.head == nil {
		return "", fmt.Errorf("list is empty")
	}
	return d.head.data, nil
}

func (d *DoublyLinkedList) GetBack() (string, error) {
	if d.tail == nil {
		return "", fmt.Errorf("list is empty")
	}
	return d.tail.data, nil
}

func (d *DoublyLinkedList) IsEmpty() bool {
	return d.size == 0
}

func (d *DoublyLinkedList) Size() int {
	return d.size
}

func (d *DoublyLinkedList) Clear() {
	for !d.IsEmpty() {
		d.PopFront()
	}
}

func (d *DoublyLinkedList) ToVector() []string {
	result := make([]string, 0, d.size)
	current := d.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func (d *DoublyLinkedList) FromVector(elements []string) {
	d.Clear()
	for _, element := range elements {
		d.PushBack(element)
	}
}
