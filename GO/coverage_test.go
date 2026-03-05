package main

import (
	"testing"
)

// ==================== ТЕСТЫ ДЛЯ ARRAY ====================

func TestArrayBasic(t *testing.T) {
	arr := NewArray("test", 2)
	defer arr.Free()

	// Проверка начального состояния
	if arr.Size() != 0 {
		t.Errorf("Expected size 0, got %d", arr.Size())
	}
	if arr.capacity != 2 {
		t.Errorf("Expected capacity 2, got %d", arr.capacity)
	}
	if arr.name != "test" {
		t.Errorf("Expected name 'test', got '%s'", arr.name)
	}
}

func TestArrayPushBack(t *testing.T) {
	arr := NewArray("test", 2)
	defer arr.Free()

	// Добавление элементов
	arr.PushBack("first")
	if arr.Size() != 1 {
		t.Errorf("Expected size 1, got %d", arr.Size())
	}
	if arr.Get(0) != "first" {
		t.Errorf("Expected 'first', got '%s'", arr.Get(0))
	}

	arr.PushBack("second")
	if arr.Size() != 2 {
		t.Errorf("Expected size 2, got %d", arr.Size())
	}
	if arr.Get(1) != "second" {
		t.Errorf("Expected 'second', got '%s'", arr.Get(1))
	}

	// Должен вызвать resize
	arr.PushBack("third")
	if arr.Size() != 3 {
		t.Errorf("Expected size 3, got %d", arr.Size())
	}
	if arr.capacity != 4 {
		t.Errorf("Expected capacity 4 after resize, got %d", arr.capacity)
	}
	if arr.Get(2) != "third" {
		t.Errorf("Expected 'third', got '%s'", arr.Get(2))
	}
}

func TestArrayInsert(t *testing.T) {
	arr := NewArray("test", 3)
	defer arr.Free()

	arr.PushBack("A")
	arr.PushBack("C")
	arr.PushBack("D")

	// Вставка в середину
	ok := arr.Insert(1, "B")
	if !ok {
		t.Error("Insert returned false")
	}
	if arr.Size() != 4 {
		t.Errorf("Expected size 4, got %d", arr.Size())
	}
	if arr.Get(1) != "B" {
		t.Errorf("Expected 'B' at index 1, got '%s'", arr.Get(1))
	}
	if arr.Get(2) != "C" {
		t.Errorf("Expected 'C' at index 2, got '%s'", arr.Get(2))
	}

	// Вставка в начало
	ok = arr.Insert(0, "Z")
	if !ok {
		t.Error("Insert returned false")
	}
	if arr.Get(0) != "Z" {
		t.Errorf("Expected 'Z' at index 0, got '%s'", arr.Get(0))
	}

	// Вставка в конец
	ok = arr.Insert(arr.Size(), "E")
	if !ok {
		t.Error("Insert returned false")
	}
	if arr.Get(arr.Size()-1) != "E" {
		t.Errorf("Expected 'E' at end, got '%s'", arr.Get(arr.Size()-1))
	}

	// Неверный индекс
	ok = arr.Insert(-1, "X")
	if ok {
		t.Error("Insert with negative index should return false")
	}
	ok = arr.Insert(100, "X")
	if ok {
		t.Error("Insert with too large index should return false")
	}
}

func TestArrayRemove(t *testing.T) {
	arr := NewArray("test", 3)
	defer arr.Free()

	arr.PushBack("A")
	arr.PushBack("B")
	arr.PushBack("C")

	// Удаление из середины
	ok := arr.Remove(1)
	if !ok {
		t.Error("Remove returned false")
	}
	if arr.Size() != 2 {
		t.Errorf("Expected size 2, got %d", arr.Size())
	}
	if arr.Get(0) != "A" || arr.Get(1) != "C" {
		t.Errorf("After remove: expected [A C], got [%s %s]", arr.Get(0), arr.Get(1))
	}

	// Удаление из начала
	arr.Insert(0, "Z")
	ok = arr.Remove(0)
	if !ok {
		t.Error("Remove returned false")
	}
	if arr.Get(0) != "A" {
		t.Errorf("After remove from start, expected 'A' at 0, got '%s'", arr.Get(0))
	}

	// Удаление из конца
	ok = arr.Remove(arr.Size() - 1)
	if !ok {
		t.Error("Remove returned false")
	}
	if arr.Size() != 1 {
		t.Errorf("Expected size 1, got %d", arr.Size())
	}

	// Неверный индекс
	ok = arr.Remove(-1)
	if ok {
		t.Error("Remove with negative index should return false")
	}
	ok = arr.Remove(100)
	if ok {
		t.Error("Remove with too large index should return false")
	}
}

func TestArrayReplace(t *testing.T) {
	arr := NewArray("test", 2)
	defer arr.Free()

	arr.PushBack("A")
	arr.PushBack("B")

	// Замена
	ok := arr.Replace(0, "X")
	if !ok {
		t.Error("Replace returned false")
	}
	if arr.Get(0) != "X" {
		t.Errorf("Expected 'X', got '%s'", arr.Get(0))
	}

	// Неверный индекс
	ok = arr.Replace(-1, "Y")
	if ok {
		t.Error("Replace with negative index should return false")
	}
	ok = arr.Replace(100, "Y")
	if ok {
		t.Error("Replace with too large index should return false")
	}
}

func TestArrayGet(t *testing.T) {
	arr := NewArray("test", 2)
	defer arr.Free()

	arr.PushBack("A")
	arr.PushBack("B")

	// Получение
	if arr.Get(0) != "A" {
		t.Errorf("Expected 'A', got '%s'", arr.Get(0))
	}
	if arr.Get(1) != "B" {
		t.Errorf("Expected 'B', got '%s'", arr.Get(1))
	}

	// Неверный индекс (должен вернуть пустую строку и напечатать "incorrect")
	val := arr.Get(-1)
	if val != "" {
		t.Errorf("Expected empty string for invalid index, got '%s'", val)
	}
	val = arr.Get(100)
	if val != "" {
		t.Errorf("Expected empty string for invalid index, got '%s'", val)
	}
}

func TestArrayEquals(t *testing.T) {
	arr1 := NewArray("test", 3)
	arr2 := NewArray("test", 3)
	arr3 := NewArray("different", 3)
	defer arr1.Free()
	defer arr2.Free()
	defer arr3.Free()

	arr1.PushBack("A")
	arr1.PushBack("B")
	arr2.PushBack("A")
	arr2.PushBack("B")
	arr3.PushBack("A")
	arr3.PushBack("C")

	if !arr1.Equals(arr2) {
		t.Error("arr1 should equal arr2")
	}
	if arr1.Equals(arr3) {
		t.Error("arr1 should not equal arr3")
	}
}

func TestArrayCopy(t *testing.T) {
	original := NewArray("test", 3)
	defer original.Free()

	original.PushBack("A")
	original.PushBack("B")
	original.PushBack("C")

	copy := original.Copy()
	defer copy.Free()

	if !original.Equals(copy) {
		t.Error("Copy should equal original")
	}

	// Изменение копии не должно влиять на оригинал
	copy.Replace(0, "X")
	if original.Get(0) == "X" {
		t.Error("Changing copy should not affect original")
	}
}

func TestArrayToFromSlice(t *testing.T) {
	arr := NewArray("test", 3)
	defer arr.Free()

	arr.PushBack("A")
	arr.PushBack("B")
	arr.PushBack("C")

	slice := arr.toSlice()
	if len(slice) != 3 || slice[0] != "A" || slice[1] != "B" || slice[2] != "C" {
		t.Errorf("toSlice returned %v, expected [A B C]", slice)
	}

	newArr := NewArray("new", 1)
	defer newArr.Free()
	newArr.fromSlice(slice)

	if !arr.Equals(newArr) {
		t.Error("fromSlice should recreate the array")
	}
}

// ==================== ТЕСТЫ ДЛЯ STACK ====================

func TestStackBasic(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
}

func TestStackPush(t *testing.T) {
	s := NewStack()

	s.Push("A")
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
	top, _ := s.Top()
	if top != "A" {
		t.Errorf("Expected top 'A', got '%s'", top)
	}

	s.Push("B")
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}
	top, _ = s.Top()
	if top != "B" {
		t.Errorf("Expected top 'B', got '%s'", top)
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack()

	s.Push("A")
	s.Push("B")
	s.Push("C")

	s.Pop()
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}
	top, _ := s.Top()
	if top != "B" {
		t.Errorf("Expected top 'B', got '%s'", top)
	}

	s.Pop()
	s.Pop()
	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all")
	}

	// Pop on empty stack (should not panic)
	s.Pop()
}

func TestStackTop(t *testing.T) {
	s := NewStack()

	// Top on empty stack
	_, err := s.Top()
	if err == nil {
		t.Error("Top on empty stack should return error")
	}

	s.Push("A")
	val, err := s.Top()
	if err != nil || val != "A" {
		t.Errorf("Expected 'A', got '%s', error: %v", val, err)
	}
}

func TestStackPrintSimple(t *testing.T) {
    s := NewStack()
    
    // Просто вызываем Print - проверяем что не паникует
    s.Print() // Пустой стек
    
    s.Push("A")
    s.Push("B")
    s.Push("C")
    s.Print() // Стек с элементами
    
    // Если код дошел сюда - тест пройден
    // (мы не проверяем вывод, только что функция не падает)
}

func TestStackClear(t *testing.T) {
	s := NewStack()
	s.Push("A")
	s.Push("B")
	s.Push("C")

	s.Clear()
	if !s.IsEmpty() {
		t.Error("Stack should be empty after Clear")
	}
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
}

func TestStackToFromVector(t *testing.T) {
	s := NewStack()
	s.Push("bottom")
	s.Push("middle")
	s.Push("top")

	// ToVector должен вернуть от вершины к основанию
	vec := s.ToVector()
	expected := []string{"top", "middle", "bottom"}
	for i, v := range vec {
		if v != expected[i] {
			t.Errorf("ToVector[%d] = %s, expected %s", i, v, expected[i])
		}
	}

	// FromVector должен восстановить стек
	s2 := NewStack()
	s2.FromVector(vec)

	if !s.Equals(s2) {
		t.Error("Stacks not equal after FromVector")
		top1, _ := s.Top()
		top2, _ := s2.Top()
		t.Logf("Original top: %s, Loaded top: %s", top1, top2)
	}
}

func TestStackToVectorReverse(t *testing.T) {
	s := NewStack()
	s.Push("bottom")
	s.Push("middle")
	s.Push("top")

	vec := s.ToVectorReverse()
	expected := []string{"bottom", "middle", "top"}
	for i, v := range vec {
		if v != expected[i] {
			t.Errorf("ToVectorReverse[%d] = %s, expected %s", i, v, expected[i])
		}
	}
}

func TestStackEquals(t *testing.T) {
	s1 := NewStack()
	s2 := NewStack()
	s3 := NewStack()

	s1.Push("A")
	s1.Push("B")
	s2.Push("A")
	s2.Push("B")
	s3.Push("A")
	s3.Push("C")

	if !s1.Equals(s2) {
		t.Error("s1 should equal s2")
	}
	if s1.Equals(s3) {
		t.Error("s1 should not equal s3")
	}
}

func TestStackCopy(t *testing.T) {
	original := NewStack()
	original.Push("A")
	original.Push("B")
	original.Push("C")

	copy := original.Copy()

	if !original.Equals(copy) {
		t.Error("Copy should equal original")
	}

	copy.Pop()
	if original.Size() == copy.Size() {
		t.Error("Changing copy should not affect original")
	}
}

// ==================== ТЕСТЫ ДЛЯ SINGLY LINKED LIST ====================

func TestSinglyLinkedListBasic(t *testing.T) {
	list := NewSinglyLinkedList()

	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}
}

func TestSinglyLinkedListPushBack(t *testing.T) {
	list := NewSinglyLinkedList()

	list.PushBack("A")
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
	val, _ := list.Get(0)
	if val != "A" {
		t.Errorf("Expected 'A' at index 0, got '%s'", val)
	}

	list.PushBack("B")
	list.PushBack("C")
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	val, _ = list.Get(2)
	if val != "C" {
		t.Errorf("Expected 'C' at index 2, got '%s'", val)
	}
}

func TestSinglyLinkedListPushFront(t *testing.T) {
	list := NewSinglyLinkedList()

	list.PushFront("A")
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
	val, _ := list.Get(0)
	if val != "A" {
		t.Errorf("Expected 'A' at index 0, got '%s'", val)
	}

	list.PushFront("B")
	val, _ = list.Get(0)
	if val != "B" {
		t.Errorf("Expected 'B' at front, got '%s'", val)
	}
	val, _ = list.Get(1)
	if val != "A" {
		t.Errorf("Expected 'A' at index 1, got '%s'", val)
	}
}

func TestSinglyLinkedListPopFront(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")

	list.PopFront()
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
	val, _ := list.Get(0)
	if val != "B" {
		t.Errorf("Expected 'B' at front, got '%s'", val)
	}

	list.PopFront()
	list.PopFront()
	if !list.IsEmpty() {
		t.Error("List should be empty")
	}

	// PopFront on empty list (should not panic)
	list.PopFront()
}

func TestSinglyLinkedListInsert(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack("A")
	list.PushBack("C")

	// Insert in middle
	list.Insert(1, "B")
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	val, _ := list.Get(1)
	if val != "B" {
		t.Errorf("Expected 'B' at index 1, got '%s'", val)
	}

	// Insert at beginning
	list.Insert(0, "Z")
	val, _ = list.Get(0)
	if val != "Z" {
		t.Errorf("Expected 'Z' at front, got '%s'", val)
	}

	// Insert at end
	list.Insert(list.Size(), "D")
	val, _ = list.Get(list.Size() - 1)
	if val != "D" {
		t.Errorf("Expected 'D' at end, got '%s'", val)
	}

	// Insert with invalid index (should do nothing)
	list.Insert(100, "X")
	if list.Size() != 5 {
		t.Errorf("Size should remain 5, got %d", list.Size())
	}
}

func TestSinglyLinkedListRemove(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")
	list.PushBack("D")

	// Remove from middle
	list.Remove(2)
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	val, _ := list.Get(2)
	if val != "D" {
		t.Errorf("Expected 'D' at index 2, got '%s'", val)
	}

	// Remove from beginning
	list.Remove(0)
	val, _ = list.Get(0)
	if val != "B" {
		t.Errorf("Expected 'B' at front, got '%s'", val)
	}

	// Remove from end
	list.Remove(list.Size() - 1)
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}

	// Remove with invalid index (should do nothing)
	list.Remove(100)
	if list.Size() != 1 {
		t.Errorf("Size should remain 1, got %d", list.Size())
	}
}

func TestSinglyLinkedListGet(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack("A")
	list.PushBack("B")

	val, err := list.Get(0)
	if err != nil || val != "A" {
		t.Errorf("Expected 'A', got '%s', error: %v", val, err)
	}

	val, err = list.Get(1)
	if err != nil || val != "B" {
		t.Errorf("Expected 'B', got '%s', error: %v", val, err)
	}

	// Invalid index
	_, err = list.Get(100)
	if err == nil {
		t.Error("Get with invalid index should return error")
	}
}

func TestSinglyLinkedListClear(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")

	list.Clear()
	if !list.IsEmpty() {
		t.Error("List should be empty after Clear")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}
}

func TestSinglyLinkedListToFromVector(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")

	vec := list.ToVector()
	expected := []string{"A", "B", "C"}
	for i, v := range vec {
		if v != expected[i] {
			t.Errorf("ToVector[%d] = %s, expected %s", i, v, expected[i])
		}
	}

	list2 := NewSinglyLinkedList()
	list2.FromVector(vec)

	if list.Size() != list2.Size() {
		t.Error("Lists not equal after FromVector")
	}
	for i := 0; i < list.Size(); i++ {
		v1, _ := list.Get(i)
		v2, _ := list2.Get(i)
		if v1 != v2 {
			t.Errorf("At index %d: %s vs %s", i, v1, v2)
		}
	}
}

// ==================== ТЕСТЫ ДЛЯ DOUBLY LINKED LIST ====================

func TestDoublyLinkedListBasic(t *testing.T) {
	list := NewDoublyLinkedList()

	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}
}

func TestDoublyLinkedListPushBack(t *testing.T) {
	list := NewDoublyLinkedList()

	list.PushBack("A")
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
	front, _ := list.GetFront()
	if front != "A" {
		t.Errorf("Expected front 'A', got '%s'", front)
	}
	back, _ := list.GetBack()
	if back != "A" {
		t.Errorf("Expected back 'A', got '%s'", back)
	}

	list.PushBack("B")
	list.PushBack("C")
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	back, _ = list.GetBack()
	if back != "C" {
		t.Errorf("Expected back 'C', got '%s'", back)
	}
}

func TestDoublyLinkedListPushFront(t *testing.T) {
	list := NewDoublyLinkedList()

	list.PushFront("A")
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
	front, _ := list.GetFront()
	if front != "A" {
		t.Errorf("Expected front 'A', got '%s'", front)
	}

	list.PushFront("B")
	front, _ = list.GetFront()
	if front != "B" {
		t.Errorf("Expected front 'B', got '%s'", front)
	}
	val, _ := list.Get(1)
	if val != "A" {
		t.Errorf("Expected 'A' at index 1, got '%s'", val)
	}
}

func TestDoublyLinkedListPopFront(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")

	list.PopFront()
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
	front, _ := list.GetFront()
	if front != "B" {
		t.Errorf("Expected front 'B', got '%s'", front)
	}

	list.PopFront()
	list.PopFront()
	if !list.IsEmpty() {
		t.Error("List should be empty")
	}

	// PopFront on empty list (should not panic)
	list.PopFront()
}

func TestDoublyLinkedListPopBack(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")

	list.PopBack()
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
	back, _ := list.GetBack()
	if back != "B" {
		t.Errorf("Expected back 'B', got '%s'", back)
	}

	list.PopBack()
	list.PopBack()
	if !list.IsEmpty() {
		t.Error("List should be empty")
	}

	// PopBack on empty list (should not panic)
	list.PopBack()
}

func TestDoublyLinkedListInsert(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack("A")
	list.PushBack("C")
	list.PushBack("D")

	// Insert in middle
	list.Insert(1, "B")
	if list.Size() != 4 {
		t.Errorf("Expected size 4, got %d", list.Size())
	}
	val, _ := list.Get(1)
	if val != "B" {
		t.Errorf("Expected 'B' at index 1, got '%s'", val)
	}

	// Insert at beginning
	list.Insert(0, "Z")
	val, _ = list.Get(0)
	if val != "Z" {
		t.Errorf("Expected 'Z' at front, got '%s'", val)
	}

	// Insert at end
	list.Insert(list.Size(), "E")
	val, _ = list.Get(list.Size() - 1)
	if val != "E" {
		t.Errorf("Expected 'E' at end, got '%s'", val)
	}

	// Insert with invalid index (should do nothing)
	list.Insert(100, "X")
	if list.Size() != 6 {
		t.Errorf("Size should remain 6, got %d", list.Size())
	}
}

func TestDoublyLinkedListRemove(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")
	list.PushBack("D")

	// Remove from middle
	list.Remove(2)
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	val, _ := list.Get(2)
	if val != "D" {
		t.Errorf("Expected 'D' at index 2, got '%s'", val)
	}

	// Remove from beginning
	list.Remove(0)
	val, _ = list.Get(0)
	if val != "B" {
		t.Errorf("Expected 'B' at front, got '%s'", val)
	}

	// Remove from end
	list.Remove(list.Size() - 1)
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}

	// Remove with invalid index (should do nothing)
	list.Remove(100)
	if list.Size() != 1 {
		t.Errorf("Size should remain 1, got %d", list.Size())
	}
}

func TestDoublyLinkedListGet(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack("A")
	list.PushBack("B")

	val, err := list.Get(0)
	if err != nil || val != "A" {
		t.Errorf("Expected 'A', got '%s', error: %v", val, err)
	}

	val, err = list.Get(1)
	if err != nil || val != "B" {
		t.Errorf("Expected 'B', got '%s', error: %v", val, err)
	}

	// Invalid index
	_, err = list.Get(100)
	if err == nil {
		t.Error("Get with invalid index should return error")
	}
}

func TestDoublyLinkedListFrontBack(t *testing.T) {
	list := NewDoublyLinkedList()

	// Empty list
	_, err := list.GetFront()
	if err == nil {
		t.Error("GetFront on empty list should return error")
	}
	_, err = list.GetBack()
	if err == nil {
		t.Error("GetBack on empty list should return error")
	}

	list.PushBack("A")
	list.PushBack("B")

	front, _ := list.GetFront()
	if front != "A" {
		t.Errorf("Expected front 'A', got '%s'", front)
	}
	back, _ := list.GetBack()
	if back != "B" {
		t.Errorf("Expected back 'B', got '%s'", back)
	}
}

func TestDoublyLinkedListClear(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")

	list.Clear()
	if !list.IsEmpty() {
		t.Error("List should be empty after Clear")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}
}

func TestDoublyLinkedListToFromVector(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")

	vec := list.ToVector()
	expected := []string{"A", "B", "C"}
	for i, v := range vec {
		if v != expected[i] {
			t.Errorf("ToVector[%d] = %s, expected %s", i, v, expected[i])
		}
	}

	list2 := NewDoublyLinkedList()
	list2.FromVector(vec)

	if list.Size() != list2.Size() {
		t.Error("Lists not equal after FromVector")
	}
	for i := 0; i < list.Size(); i++ {
		v1, _ := list.Get(i)
		v2, _ := list2.Get(i)
		if v1 != v2 {
			t.Errorf("At index %d: %s vs %s", i, v1, v2)
		}
	}
}

// ==================== ТЕСТЫ ДЛЯ HASH TABLE ====================

func TestHashTableBasic(t *testing.T) {
	ht := NewOpenAddressingHashTable(10)

	if !ht.IsEmpty() {
		t.Error("New hash table should be empty")
	}
	if ht.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", ht.GetSize())
	}
	if ht.capacity != 10 {
		t.Errorf("Expected capacity 10, got %d", ht.capacity)
	}
}

func TestHashTableInsert(t *testing.T) {
	ht := NewOpenAddressingHashTable(5)

	// Вставка новых элементов
	ht.Insert('a', 1)
	if ht.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", ht.GetSize())
	}
	val, ok := ht.Search('a')
	if !ok || val != 1 {
		t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
	}

	ht.Insert('b', 2)
	ht.Insert('c', 3)
	if ht.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", ht.GetSize())
	}

	// Перезапись существующего ключа
	ht.Insert('a', 100)
	if ht.GetSize() != 3 {
		t.Errorf("Size should remain 3 after overwrite, got %d", ht.GetSize())
	}
	val, _ = ht.Search('a')
	if val != 100 {
		t.Errorf("After overwrite, expected 100, got %d", val)
	}
}

func TestHashTableSearch(t *testing.T) {
	ht := NewOpenAddressingHashTable(5)
	ht.Insert('a', 1)
	ht.Insert('b', 2)

	// Поиск существующих
	val, ok := ht.Search('a')
	if !ok || val != 1 {
		t.Errorf("Search 'a': expected (1, true), got (%d, %v)", val, ok)
	}

	val, ok = ht.Search('b')
	if !ok || val != 2 {
		t.Errorf("Search 'b': expected (2, true), got (%d, %v)", val, ok)
	}

	// Поиск несуществующего
	val, ok = ht.Search('x')
	if ok || val != 0 {
		t.Errorf("Search 'x': expected (0, false), got (%d, %v)", val, ok)
	}
}

func TestHashTableRemove(t *testing.T) {
	ht := NewOpenAddressingHashTable(5)
	ht.Insert('a', 1)
	ht.Insert('b', 2)
	ht.Insert('c', 3)

	// Удаление существующего
	ok := ht.Remove('b')
	if !ok {
		t.Error("Remove should return true for existing key")
	}
	if ht.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", ht.GetSize())
	}
	_, ok = ht.Search('b')
	if ok {
		t.Error("Key 'b' should not exist after removal")
	}

	// Удаление несуществующего
	ok = ht.Remove('x')
	if ok {
		t.Error("Remove should return false for non-existing key")
	}
}

func TestHashTableClear(t *testing.T) {
	ht := NewOpenAddressingHashTable(5)
	ht.Insert('a', 1)
	ht.Insert('b', 2)
	ht.Insert('c', 3)

	ht.Clear()
	if !ht.IsEmpty() {
		t.Error("Hash table should be empty after Clear")
	}
	if ht.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", ht.GetSize())
	}

	// Проверка, что можно добавлять после очистки
	ht.Insert('a', 1)
	if ht.GetSize() != 1 {
		t.Errorf("Expected size 1 after insert, got %d", ht.GetSize())
	}
}

func TestHashTableCollision(t *testing.T) {
	ht := NewOpenAddressingHashTable(3)

	// Заполняем таблицу
	ht.Insert('a', 1) // хеш 97 % 3 = 1
	ht.Insert('b', 2) // хеш 98 % 3 = 2
	ht.Insert('c', 3) // хеш 99 % 3 = 0

	// Должна возникнуть коллизия
	ht.Insert('d', 4) // хеш 100 % 3 = 1 (коллизия с 'a')

	if ht.GetSize() != 4 {
		t.Errorf("Expected size 4, got %d", ht.GetSize())
	}

	// Проверка, что все элементы доступны
	val, ok := ht.Search('a')
	if !ok || val != 1 {
		t.Error("Key 'a' should exist")
	}
	val, ok = ht.Search('d')
	if !ok || val != 4 {
		t.Error("Key 'd' should exist despite collision")
	}
}

func TestHashTableFull(t *testing.T) {
	ht := NewOpenAddressingHashTable(2)

	ht.Insert('a', 1)
	ht.Insert('b', 2)
	ht.Insert('c', 3) // Таблица полна, вставка должна игнорироваться

	if ht.GetSize() != 2 {
		t.Errorf("Expected size 2 (full table), got %d", ht.GetSize())
	}

	_, ok := ht.Search('c')
	if ok {
		t.Error("Key 'c' should not be inserted when table is full")
	}
}

func TestHashTableToFromVector(t *testing.T) {
	ht := NewOpenAddressingHashTable(10)
	ht.Insert('a', 1)
	ht.Insert('b', 2)
	ht.Insert('c', 3)

	vec := ht.ToVector()
	if len(vec) != 3 {
		t.Errorf("Expected vector length 3, got %d", len(vec))
	}

	ht2 := NewOpenAddressingHashTable(5)
	ht2.FromVector(vec)

	if ht.GetSize() != ht2.GetSize() {
		t.Errorf("Size mismatch: %d vs %d", ht.GetSize(), ht2.GetSize())
	}

	// Проверка элементов
	keys := []rune{'a', 'b', 'c'}
	for _, key := range keys {
		val1, ok1 := ht.Search(key)
		val2, ok2 := ht2.Search(key)
		if ok1 != ok2 || val1 != val2 {
			t.Errorf("Key %c mismatch: (%d,%v) vs (%d,%v)", key, val1, ok1, val2, ok2)
		}
	}
}

// ==================== ТЕСТЫ ДЛЯ BINARY TREE ====================

func TestTreeBasic(t *testing.T) {
	tree := NewFullBinaryTree()

	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}
}

func TestTreeInsert(t *testing.T) {
	tree := NewFullBinaryTree()

	// Вставка корня
	tree.Insert(5)
	if tree.IsEmpty() {
		t.Error("Tree should not be empty after insert")
	}

	// Вставка левого и правого
	tree.Insert(3)
	tree.Insert(7)

	vec := tree.ToVector()
	expected := []int{3, 5, 7} // in-order обход
	for i, v := range vec {
		if v != expected[i] {
			t.Errorf("ToVector[%d] = %d, expected %d", i, v, expected[i])
		}
	}
}

func TestTreeInsertDuplicate(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.Insert(5)
	tree.Insert(5) // Должен пойти в правое поддерево

	vec := tree.ToVector()
	if len(vec) != 2 {
		t.Errorf("Expected size 2, got %d", len(vec))
	}
	if vec[0] != 5 || vec[1] != 5 {
		t.Errorf("Expected [5 5], got %v", vec)
	}
}

func TestTreeComplex(t *testing.T) {
	tree := NewFullBinaryTree()
	values := []int{5, 3, 7, 2, 4, 6, 8, 1, 9}
	for _, v := range values {
		tree.Insert(v)
	}

	vec := tree.ToVector()
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // in-order должен быть отсортирован
	for i, v := range vec {
		if v != expected[i] {
			t.Errorf("In-order: expected %v, got %v", expected, vec)
			break
		}
	}
}

func TestTreeToFromVector(t *testing.T) {
	original := NewFullBinaryTree()
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		original.Insert(v)
	}

	vec := original.ToVector()

	newTree := NewFullBinaryTree()
	newTree.FromVector(vec)

	origVec := original.ToVector()
	newVec := newTree.ToVector()

	if len(origVec) != len(newVec) {
		t.Errorf("Size mismatch: %d vs %d", len(origVec), len(newVec))
	}

	for i := 0; i < len(origVec); i++ {
		if origVec[i] != newVec[i] {
			t.Errorf("At index %d: %d vs %d", i, origVec[i], newVec[i])
		}
	}
}

func TestTreePrint(t *testing.T) {
	// Просто проверяем, что функции не паникуют
	tree := NewFullBinaryTree()
	tree.Print() // Пустое дерево
	tree.PrintZigZag() // Пустое дерево

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Print()
	tree.PrintZigZag()
}

func TestTreeIsEmpty(t *testing.T) {
	tree := NewFullBinaryTree()
	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}

	tree.Insert(5)
	if tree.IsEmpty() {
		t.Error("Tree with node should not be empty")
	}
}

// ==================== ТЕСТЫ ДЛЯ ГРАНИЧНЫХ СЛУЧАЕВ ====================

func TestEdgeCases(t *testing.T) {
	// Array с нулевой емкостью
	arr := NewArray("test", 0)
	defer arr.Free()
	arr.PushBack("test") // Должен вызвать resize с 0*2 = 0? Проверяем что не паникует
	if arr.capacity == 0 {
		t.Log("Note: capacity 0 after push?")
	}

	// HashTable с нулевой емкостью (должна стать 256)
	ht := NewOpenAddressingHashTable(0)
	if ht.capacity != 256 {
		t.Errorf("Expected default capacity 256, got %d", ht.capacity)
	}

	// Stack операции после очистки
	s := NewStack()
	s.Push("A")
	s.Clear()
	s.Push("B") // Должно работать
	val, _ := s.Top()
	if val != "B" {
		t.Errorf("Expected 'B', got '%s'", val)
	}
}