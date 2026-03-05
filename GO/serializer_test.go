package main

import (
	"encoding/binary"  // ← добавьте этот импорт
	"fmt"              // ← добавьте для fmt.Errorf
	"os"
	"testing"
)

// ==================== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ====================

func cleanup(files ...string) {
	for _, file := range files {
		os.Remove(file)
	}
}

// ==================== ТЕСТЫ ДЛЯ ARRAY ====================

func TestArrayBinarySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_array.bin"
	defer cleanup(filename)

	// Создание исходного массива
	original := NewArray("testArray", 2)
	defer original.Free()

	// Заполнение данными
	original.PushBack("hello")
	original.PushBack("world")
	original.PushBack("go")
	original.Insert(1, "beautiful")

	// Сохранение
	err := serializer.BinarySerializeArray(original, filename)
	if err != nil {
		t.Fatalf("Failed to save array: %v", err)
	}

	// Загрузка
	loaded := NewArray("", 1)
	defer loaded.Free()
	err = serializer.BinaryDeserializeArray(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load array: %v", err)
	}

	// Сравнение
	if !original.Equals(loaded) {
		t.Error("Binary serialization: arrays are not equal")
		t.Logf("Original: %v", original.toSlice())
		t.Logf("Loaded: %v", loaded.toSlice())
	}
}

func TestArrayTextSerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_array.txt"
	defer cleanup(filename)

	original := NewArray("testArray", 2)
	defer original.Free()

	original.PushBack("hello")
	original.PushBack("world")
	original.PushBack("go")

	err := serializer.TextSerializeArray(original, filename)
	if err != nil {
		t.Fatalf("Failed to save array: %v", err)
	}

	loaded := NewArray("", 1)
	defer loaded.Free()
	err = serializer.TextDeserializeArray(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load array: %v", err)
	}

	if !original.Equals(loaded) {
		t.Error("Text serialization: arrays are not equal")
	}
}

func TestArrayEmptySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_array_empty.bin"
	defer cleanup(filename)

	original := NewArray("empty", 5)
	defer original.Free()

	err := serializer.BinarySerializeArray(original, filename)
	if err != nil {
		t.Fatalf("Failed to save empty array: %v", err)
	}

	loaded := NewArray("", 1)
	defer loaded.Free()
	err = serializer.BinaryDeserializeArray(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load empty array: %v", err)
	}

	if !original.Equals(loaded) {
		t.Error("Empty array serialization failed")
	}
}

func TestArrayLargeData(t *testing.T) {
    serializer := NewSerializer()
    filename := "test_array_large.bin"
    defer cleanup(filename)

    original := NewArray("large", 10)
    defer original.Free()

    // Заполнение большим количеством данных
    for i := 0; i < 1000; i++ {
        original.PushBack("test string with number")
    }

    err := serializer.BinarySerializeArray(original, filename)
    if err != nil {
        t.Fatalf("Failed to save large array: %v", err)
    }

    loaded := NewArray("", 1)
    defer loaded.Free()
    err = serializer.BinaryDeserializeArray(loaded, filename)
    if err != nil {
        t.Fatalf("Failed to load large array: %v", err)
    }

    // ИСПРАВЛЕНО: используем поля, а не методы
    if original.size != loaded.size {
        t.Errorf("Size mismatch: original=%d, loaded=%d", original.size, loaded.size)
    }
}

// ==================== ТЕСТЫ ДЛЯ STACK ====================

func TestStackBinarySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_stack.bin"
	defer cleanup(filename)

	original := NewStack()
	original.Push("third")
	original.Push("second")
	original.Push("first")

	err := serializer.BinarySerializeStack(original, filename)
	if err != nil {
		t.Fatalf("Failed to save stack: %v", err)
	}

	loaded := NewStack()
	err = serializer.BinaryDeserializeStack(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load stack: %v", err)
	}

	if !original.Equals(loaded) {
		t.Error("Binary serialization: stacks are not equal")
	}
}

func TestStackTextSerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_stack.txt"
	defer cleanup(filename)

	original := NewStack()
	original.Push("A")
	original.Push("B")
	original.Push("C")

	err := serializer.TextSerializeStack(original, filename)
	if err != nil {
		t.Fatalf("Failed to save stack: %v", err)
	}

	loaded := NewStack()
	err = serializer.TextDeserializeStack(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load stack: %v", err)
	}

	if !original.Equals(loaded) {
		t.Error("Text serialization: stacks are not equal")
	}
}

func TestStackSimplifiedSerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_stack_simple.txt"
	defer cleanup(filename)

	original := NewStack()
	original.Push("X")
	original.Push("Y")
	original.Push("Z")

	err := serializer.SimplifiedTextSerializeStack(original, filename)
	if err != nil {
		t.Fatalf("Failed to save stack: %v", err)
	}

	loaded := NewStack()
	err = serializer.SimplifiedTextDeserializeStack(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load stack: %v", err)
	}

	if !original.Equals(loaded) {
		t.Error("Simplified serialization: stacks are not equal")
	}
}

func TestStackCompactBinarySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_stack_compact.bin"
	defer cleanup(filename)

	original := NewStack()
	original.Push("one")
	original.Push("two")
	original.Push("three")

	err := serializer.CompactBinarySerializeStack(original, filename)
	if err != nil {
		t.Fatalf("Failed to save stack: %v", err)
	}

	loaded := NewStack()
	err = serializer.CompactBinaryDeserializeStack(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load stack: %v", err)
	}

	if !original.Equals(loaded) {
		t.Error("Compact binary serialization: stacks are not equal")
	}
}

func TestStackEmptySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_stack_empty.bin"
	defer cleanup(filename)

	original := NewStack()

	err := serializer.BinarySerializeStack(original, filename)
	if err != nil {
		t.Fatalf("Failed to save empty stack: %v", err)
	}

	loaded := NewStack()
	err = serializer.BinaryDeserializeStack(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load empty stack: %v", err)
	}

	if original.Size() != loaded.Size() {
		t.Errorf("Size mismatch: original=%d, loaded=%d", original.Size(), loaded.Size())
	}
}

func TestStackLIFOOrder(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_stack_order.bin"
	defer cleanup(filename)

	original := NewStack()
	original.Push("bottom")
	original.Push("middle")
	original.Push("top")

	err := serializer.BinarySerializeStack(original, filename)
	if err != nil {
		t.Fatalf("Failed to save stack: %v", err)
	}

	loaded := NewStack()
	err = serializer.BinaryDeserializeStack(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load stack: %v", err)
	}

	// Проверка порядка LIFO
	top1, _ := original.Top()
	top2, _ := loaded.Top()
	if top1 != top2 {
		t.Errorf("Top element mismatch: original=%s, loaded=%s", top1, top2)
	}
}

// ==================== ТЕСТЫ ДЛЯ SINGLY LINKED LIST ====================

func TestSinglyLinkedListBinarySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_sll.bin"
	defer cleanup(filename)

	original := NewSinglyLinkedList()
	original.PushBack("first")
	original.PushBack("second")
	original.PushBack("third")
	original.PushFront("zero")

	err := serializer.BinarySerializeSinglyLinkedList(original, filename)
	if err != nil {
		t.Fatalf("Failed to save singly linked list: %v", err)
	}

	loaded := NewSinglyLinkedList()
	err = serializer.BinaryDeserializeSinglyLinkedList(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load singly linked list: %v", err)
	}

	// Сравнение через ToVector
	origVec := original.ToVector()
	loadVec := loaded.ToVector()
	
	if len(origVec) != len(loadVec) {
		t.Errorf("Size mismatch: original=%d, loaded=%d", len(origVec), len(loadVec))
	}
	
	for i := 0; i < len(origVec); i++ {
		if origVec[i] != loadVec[i] {
			t.Errorf("Element %d mismatch: %s vs %s", i, origVec[i], loadVec[i])
		}
	}
}

func TestSinglyLinkedListTextSerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_sll.txt"
	defer cleanup(filename)

	original := NewSinglyLinkedList()
	original.PushBack("A")
	original.PushBack("B")
	original.PushBack("C")

	err := serializer.TextSerializeSinglyLinkedList(original, filename)
	if err != nil {
		t.Fatalf("Failed to save singly linked list: %v", err)
	}

	loaded := NewSinglyLinkedList()
	err = serializer.TextDeserializeSinglyLinkedList(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load singly linked list: %v", err)
	}

	origVec := original.ToVector()
	loadVec := loaded.ToVector()
	
	for i := 0; i < len(origVec); i++ {
		if origVec[i] != loadVec[i] {
			t.Errorf("Element %d mismatch", i)
		}
	}
}

func TestSinglyLinkedListEmptySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_sll_empty.bin"
	defer cleanup(filename)

	original := NewSinglyLinkedList()

	err := serializer.BinarySerializeSinglyLinkedList(original, filename)
	if err != nil {
		t.Fatalf("Failed to save empty list: %v", err)
	}

	loaded := NewSinglyLinkedList()
	err = serializer.BinaryDeserializeSinglyLinkedList(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load empty list: %v", err)
	}

	if original.Size() != loaded.Size() {
		t.Errorf("Size mismatch: original=%d, loaded=%d", original.Size(), loaded.Size())
	}
}

// ==================== ТЕСТЫ ДЛЯ DOUBLY LINKED LIST ====================

func TestDoublyLinkedListBinarySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_dll.bin"
	defer cleanup(filename)

	original := NewDoublyLinkedList()
	original.PushBack("first")
	original.PushBack("second")
	original.PushBack("third")
	original.PushFront("zero")

	err := serializer.BinarySerializeDoublyLinkedList(original, filename)
	if err != nil {
		t.Fatalf("Failed to save doubly linked list: %v", err)
	}

	loaded := NewDoublyLinkedList()
	err = serializer.BinaryDeserializeDoublyLinkedList(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load doubly linked list: %v", err)
	}

	origVec := original.ToVector()
	loadVec := loaded.ToVector()
	
	if len(origVec) != len(loadVec) {
		t.Errorf("Size mismatch: original=%d, loaded=%d", len(origVec), len(loadVec))
	}
	
	for i := 0; i < len(origVec); i++ {
		if origVec[i] != loadVec[i] {
			t.Errorf("Element %d mismatch: %s vs %s", i, origVec[i], loadVec[i])
		}
	}
}

func TestDoublyLinkedListTextSerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_dll.txt"
	defer cleanup(filename)

	original := NewDoublyLinkedList()
	original.PushBack("X")
	original.PushBack("Y")
	original.PushBack("Z")

	err := serializer.TextSerializeDoublyLinkedList(original, filename)
	if err != nil {
		t.Fatalf("Failed to save doubly linked list: %v", err)
	}

	loaded := NewDoublyLinkedList()
	err = serializer.TextDeserializeDoublyLinkedList(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load doubly linked list: %v", err)
	}

	origVec := original.ToVector()
	loadVec := loaded.ToVector()
	
	for i := 0; i < len(origVec); i++ {
		if origVec[i] != loadVec[i] {
			t.Errorf("Element %d mismatch", i)
		}
	}
}

func TestDoublyLinkedListBidirectionalCheck(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_dll_check.bin"
	defer cleanup(filename)

	original := NewDoublyLinkedList()
	original.PushBack("first")
	original.PushBack("second")
	original.PushBack("third")

	err := serializer.BinarySerializeDoublyLinkedList(original, filename)
	if err != nil {
		t.Fatalf("Failed to save list: %v", err)
	}

	loaded := NewDoublyLinkedList()
	err = serializer.BinaryDeserializeDoublyLinkedList(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load list: %v", err)
	}

	// Проверка первого элемента
	front1, _ := original.GetFront()
	front2, _ := loaded.GetFront()
	if front1 != front2 {
		t.Errorf("Front element mismatch: %s vs %s", front1, front2)
	}

	// Проверка последнего элемента
	back1, _ := original.GetBack()
	back2, _ := loaded.GetBack()
	if back1 != back2 {
		t.Errorf("Back element mismatch: %s vs %s", back1, back2)
	}
}

// ==================== ТЕСТЫ ДЛЯ HASH TABLE ====================

func TestHashTableBinarySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_ht.bin"
	defer cleanup(filename)

	original := NewOpenAddressingHashTable(10)
	original.Insert('a', 1)
	original.Insert('b', 2)
	original.Insert('c', 3)
	original.Insert('d', 4)

	err := serializer.BinarySerializeHashTable(original, filename)
	if err != nil {
		t.Fatalf("Failed to save hash table: %v", err)
	}

	loaded := NewOpenAddressingHashTable(1)
	err = serializer.BinaryDeserializeHashTable(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load hash table: %v", err)
	}

	// Проверка размера
	if original.GetSize() != loaded.GetSize() {
		t.Errorf("Size mismatch: original=%d, loaded=%d", original.GetSize(), loaded.GetSize())
	}

	// Проверка элементов
	keys := []rune{'a', 'b', 'c', 'd'}
	for _, key := range keys {
		val1, ok1 := original.Search(key)
		val2, ok2 := loaded.Search(key)
		
		if ok1 != ok2 {
			t.Errorf("Key %c existence mismatch", key)
		}
		if val1 != val2 {
			t.Errorf("Key %c value mismatch: %d vs %d", key, val1, val2)
		}
	}
}

func TestHashTableTextSerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_ht.txt"
	defer cleanup(filename)

	original := NewOpenAddressingHashTable(10)
	original.Insert('x', 10)
	original.Insert('y', 20)
	original.Insert('z', 30)

	err := serializer.TextSerializeHashTable(original, filename)
	if err != nil {
		t.Fatalf("Failed to save hash table: %v", err)
	}

	loaded := NewOpenAddressingHashTable(1)
	err = serializer.TextDeserializeHashTable(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load hash table: %v", err)
	}

	keys := []rune{'x', 'y', 'z'}
	for _, key := range keys {
		val1, ok1 := original.Search(key)
		val2, ok2 := loaded.Search(key)
		
		if ok1 != ok2 {
			t.Errorf("Key %c existence mismatch", key)
		}
		if val1 != val2 {
			t.Errorf("Key %c value mismatch", key)
		}
	}
}

func TestHashTableEmptySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_ht_empty.bin"
	defer cleanup(filename)

	original := NewOpenAddressingHashTable(10)

	err := serializer.BinarySerializeHashTable(original, filename)
	if err != nil {
		t.Fatalf("Failed to save empty hash table: %v", err)
	}

	loaded := NewOpenAddressingHashTable(1)
	err = serializer.BinaryDeserializeHashTable(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load empty hash table: %v", err)
	}

	if original.GetSize() != loaded.GetSize() {
		t.Errorf("Size mismatch: original=%d, loaded=%d", original.GetSize(), loaded.GetSize())
	}
}

func TestHashTableOverwriteSerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_ht_overwrite.bin"
	defer cleanup(filename)

	original := NewOpenAddressingHashTable(10)
	original.Insert('a', 1)
	original.Insert('a', 100) // Перезапись

	err := serializer.BinarySerializeHashTable(original, filename)
	if err != nil {
		t.Fatalf("Failed to save hash table: %v", err)
	}

	loaded := NewOpenAddressingHashTable(1)
	err = serializer.BinaryDeserializeHashTable(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load hash table: %v", err)
	}

	val, ok := loaded.Search('a')
	if !ok || val != 100 {
		t.Errorf("Overwritten value not preserved: got %d, expected 100", val)
	}
}

// ==================== ТЕСТЫ ДЛЯ BINARY TREE ====================

func TestTreeBinarySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_tree.bin"
	defer cleanup(filename)

	original := NewFullBinaryTree()
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		original.Insert(v)
	}

	err := serializer.BinarySerializeTree(original, filename)
	if err != nil {
		t.Fatalf("Failed to save tree: %v", err)
	}

	loaded := NewFullBinaryTree()
	err = serializer.BinaryDeserializeTree(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load tree: %v", err)
	}

	// Сравнение через ToVector (in-order обход)
	origVec := original.ToVector()
	loadVec := loaded.ToVector()

	if len(origVec) != len(loadVec) {
		t.Errorf("Size mismatch: original=%d, loaded=%d", len(origVec), len(loadVec))
	}

	for i := 0; i < len(origVec); i++ {
		if origVec[i] != loadVec[i] {
			t.Errorf("Element %d mismatch: %d vs %d", i, origVec[i], loadVec[i])
		}
	}
}

func TestTreeTextSerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_tree.txt"
	defer cleanup(filename)

	original := NewFullBinaryTree()
	values := []int{10, 5, 15, 3, 7, 12, 18}
	for _, v := range values {
		original.Insert(v)
	}

	err := serializer.TextSerializeTree(original, filename)
	if err != nil {
		t.Fatalf("Failed to save tree: %v", err)
	}

	loaded := NewFullBinaryTree()
	err = serializer.TextDeserializeTree(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load tree: %v", err)
	}

	origVec := original.ToVector()
	loadVec := loaded.ToVector()

	if len(origVec) != len(loadVec) {
		t.Errorf("Size mismatch: original=%d, loaded=%d", len(origVec), len(loadVec))
	}
}

func TestTreeEmptySerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_tree_empty.bin"
	defer cleanup(filename)

	original := NewFullBinaryTree()

	err := serializer.BinarySerializeTree(original, filename)
	if err != nil {
		t.Fatalf("Failed to save empty tree: %v", err)
	}

	loaded := NewFullBinaryTree()
	err = serializer.BinaryDeserializeTree(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load empty tree: %v", err)
	}

	if len(original.ToVector()) != len(loaded.ToVector()) {
		t.Error("Empty tree serialization failed")
	}
}

func TestTreeSingleNodeSerialization(t *testing.T) {
	serializer := NewSerializer()
	filename := "test_tree_single.bin"
	defer cleanup(filename)

	original := NewFullBinaryTree()
	original.Insert(42)

	err := serializer.BinarySerializeTree(original, filename)
	if err != nil {
		t.Fatalf("Failed to save single node tree: %v", err)
	}

	loaded := NewFullBinaryTree()
	err = serializer.BinaryDeserializeTree(loaded, filename)
	if err != nil {
		t.Fatalf("Failed to load single node tree: %v", err)
	}

	origVec := original.ToVector()
	loadVec := loaded.ToVector()

	if len(origVec) != 1 || len(loadVec) != 1 {
		t.Errorf("Wrong size: original=%d, loaded=%d", len(origVec), len(loadVec))
	}

	if origVec[0] != loadVec[0] {
		t.Errorf("Value mismatch: %d vs %d", origVec[0], loadVec[0])
	}
}

// ==================== ТЕСТЫ НА ОШИБКИ И ГРАНИЧНЫЕ СЛУЧАИ ====================

func TestInvalidFileFormat(t *testing.T) {
	serializer := NewSerializer()
	filename := "invalid.bin"
	
	// Создаем некорректный файл
	f, _ := os.Create(filename)
	f.Write([]byte("invalid data"))
	f.Close()
	defer cleanup(filename)

	arr := NewArray("", 1)
	defer arr.Free()
	
	err := serializer.BinaryDeserializeArray(arr, filename)
	if err == nil {
		t.Error("Expected error for invalid file format, got nil")
	}
}

func TestNonExistentFile(t *testing.T) {
	serializer := NewSerializer()
	filename := "nonexistent.bin"

	arr := NewArray("", 1)
	defer arr.Free()
	
	err := serializer.BinaryDeserializeArray(arr, filename)
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestCorruptedBinaryFile(t *testing.T) {
	serializer := NewSerializer()
	filename := "corrupted.bin"
	
	// Создаем файл с неполными данными
	f, _ := os.Create(filename)
	binary.Write(f, binary.LittleEndian, uint32(1)) // версия
	binary.Write(f, binary.LittleEndian, uint64(5)) // размер 5, но данных нет
	f.Close()
	defer cleanup(filename)

	stack := NewStack()
	err := serializer.BinaryDeserializeStack(stack, filename)
	if err == nil {
		t.Error("Expected error for corrupted file, got nil")
	}
}

// ==================== ТЕСТЫ НА ПРОИЗВОДИТЕЛЬНОСТЬ ====================

func BenchmarkArraySerialization(b *testing.B) {
	serializer := NewSerializer()
	filename := "bench_array.bin"
	defer cleanup(filename)

	arr := NewArray("bench", 1000)
	defer arr.Free()
	
	for i := 0; i < 1000; i++ {
		arr.PushBack("test string")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		serializer.BinarySerializeArray(arr, filename)
		loaded := NewArray("", 1)
		serializer.BinaryDeserializeArray(loaded, filename)
		loaded.Free()
	}
}

func BenchmarkStackSerialization(b *testing.B) {
	serializer := NewSerializer()
	filename := "bench_stack.bin"
	defer cleanup(filename)

	stack := NewStack()
	for i := 0; i < 1000; i++ {
		stack.Push("test string")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		serializer.BinarySerializeStack(stack, filename)
		loaded := NewStack()
		serializer.BinaryDeserializeStack(loaded, filename)
	}
}

func BenchmarkHashTableSerialization(b *testing.B) {
	serializer := NewSerializer()
	filename := "bench_ht.bin"
	defer cleanup(filename)

	ht := NewOpenAddressingHashTable(1000)
	for i := 0; i < 100; i++ {
		ht.Insert(rune(i), i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		serializer.BinarySerializeHashTable(ht, filename)
		loaded := NewOpenAddressingHashTable(1)
		serializer.BinaryDeserializeHashTable(loaded, filename)
	}
}

// ==================== ТЕСТЫ НА ПАРАЛЛЕЛЬНОСТЬ ====================

func TestConcurrentSerialization(t *testing.T) {
	serializer := NewSerializer()
	
	// Запускаем несколько горутин для параллельной сериализации
	done := make(chan bool)
	
	for i := 0; i < 10; i++ {
		go func(id int) {
			filename := fmt.Sprintf("concurrent_%d.bin", id)
			defer cleanup(filename)
			
			arr := NewArray("test", 10)
			defer arr.Free()
			
			arr.PushBack("data")
			
			err := serializer.BinarySerializeArray(arr, filename)
			if err != nil {
				t.Errorf("Concurrent serialization failed: %v", err)
			}
			
			loaded := NewArray("", 1)
			defer loaded.Free()
			
			err = serializer.BinaryDeserializeArray(loaded, filename)
			if err != nil {
				t.Errorf("Concurrent deserialization failed: %v", err)
			}
			
			done <- true
		}(i)
	}
	
	// Ожидаем завершения всех горутин
	for i := 0; i < 10; i++ {
		<-done
	}
}