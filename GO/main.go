package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Демонстрация работы структур данных и сериализации ===")

	// Создаем сериализатор
	serializer := NewSerializer()

	// ===== ДЕМОНСТРАЦИЯ ARRAY =====
	fmt.Println("--- Array ---")
	arr := NewArray("myArray", 2)
	defer arr.Free()
	
	arr.PushBack("hello")
	arr.PushBack("world")
	arr.PushBack("go")
	arr.Insert(1, "beautiful")
	
	fmt.Println("Исходный массив:")
	arr.Print()
	
	// Сохраняем
	serializer.BinarySerializeArray(arr, "array.bin")
	serializer.TextSerializeArray(arr, "array.txt")
	fmt.Println("Массив сохранен в array.bin и array.txt")
	
	// Загружаем
	loadedArr := NewArray("loaded", 1)
	defer loadedArr.Free()
	serializer.BinaryDeserializeArray(loadedArr, "array.bin")
	fmt.Println("Загруженный массив:")
	loadedArr.Print()
	fmt.Println()

	// ===== ДЕМОНСТРАЦИЯ STACK =====
	fmt.Println("--- Stack ---")
	stack := NewStack()
	stack.Push("третий")
	stack.Push("второй")
	stack.Push("первый")
	
	fmt.Println("Исходный стек:")
	stack.Print()
	
	serializer.BinarySerializeStack(stack, "stack.bin")
	fmt.Println("Стек сохранен в stack.bin")
	
	loadedStack := NewStack()
	serializer.BinaryDeserializeStack(loadedStack, "stack.bin")
	fmt.Println("Загруженный стек:")
	loadedStack.Print()
	fmt.Println()

	// ===== ДЕМОНСТРАЦИЯ СПИСКОВ =====
	fmt.Println("--- Связные списки ---")
	
	// Односвязный список
	sll := NewSinglyLinkedList()
	sll.PushBack("A")
	sll.PushBack("B")
	sll.PushBack("C")
	
	fmt.Println("Односвязный список:", sll.ToVector())
	serializer.BinarySerializeSinglyLinkedList(sll, "sll.bin")
	
	// Двусвязный список
	dll := NewDoublyLinkedList()
	dll.PushBack("X")
	dll.PushBack("Y")
	dll.PushBack("Z")
	
	fmt.Println("Двусвязный список:", dll.ToVector())
	serializer.BinarySerializeDoublyLinkedList(dll, "dll.bin")
	fmt.Println()

	// ===== ДЕМОНСТРАЦИЯ ХЕШ-ТАБЛИЦЫ =====
	fmt.Println("--- Хеш-таблица ---")
	ht := NewOpenAddressingHashTable(10)
	ht.Insert('a', 100)
	ht.Insert('b', 200)
	ht.Insert('c', 300)
	
	fmt.Printf("Хеш-таблица: размер=%d\n", ht.GetSize())
	val, _ := ht.Search('b')
	fmt.Printf("Поиск 'b': %d\n", val)
	
	serializer.BinarySerializeHashTable(ht, "ht.bin")
	fmt.Println("Хеш-таблица сохранена в ht.bin")
	fmt.Println()

	// ===== ДЕМОНСТРАЦИЯ ДЕРЕВА =====
	fmt.Println("--- Бинарное дерево ---")
	tree := NewFullBinaryTree()
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		tree.Insert(v)
	}
	
	fmt.Println("Дерево (in-order):", tree.ToVector())
	fmt.Println("Zig-zag обход:")
	tree.PrintZigZag()
	
	serializer.BinarySerializeTree(tree, "tree.bin")
	fmt.Println("Дерево сохранено в tree.bin")

	fmt.Println("\n=== Все данные успешно сохранены ===")
}