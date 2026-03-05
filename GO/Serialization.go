package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// Serializer - универсальный сериализатор для всех структур
type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

// ==================== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ====================

func (s *Serializer) writeString(file *os.File, str string) error {
	strBytes := []byte(str)
	strLen := uint32(len(strBytes))
	
	if err := binary.Write(file, binary.LittleEndian, strLen); err != nil {
		return err
	}
	
	_, err := file.Write(strBytes)
	return err
}

func (s *Serializer) readString(file *os.File) (string, error) {
	var strLen uint32
	if err := binary.Read(file, binary.LittleEndian, &strLen); err != nil {
		return "", err
	}
	
	if strLen > 1000000 {
		return "", fmt.Errorf("string too long: %d", strLen)
	}
	
	strBytes := make([]byte, strLen)
	if _, err := io.ReadFull(file, strBytes); err != nil {
		return "", err
	}
	
	return string(strBytes), nil
}

func (s *Serializer) writeInt(file *os.File, val int) error {
	return binary.Write(file, binary.LittleEndian, uint64(val))
}

func (s *Serializer) readInt(file *os.File) (int, error) {
	var val uint64
	if err := binary.Read(file, binary.LittleEndian, &val); err != nil {
		return 0, err
	}
	return int(val), nil
}

// ==================== ARRAY ====================

func (s *Serializer) BinarySerializeArray(arr *Array, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Версия (опционально)
	if err := binary.Write(file, binary.LittleEndian, uint32(1)); err != nil {
		return err
	}

	// Имя массива
	if err := s.writeString(file, arr.name); err != nil {
		return err
	}

	// Емкость и размер
	if err := s.writeInt(file, arr.capacity); err != nil {
		return err
	}
	if err := s.writeInt(file, arr.size); err != nil {
		return err
	}

	// Элементы
	data := arr.toSlice()
	if err := s.writeInt(file, len(data)); err != nil {
		return err
	}

	for _, str := range data {
		if err := s.writeString(file, str); err != nil {
			return err
		}
	}

	return nil
}

func (s *Serializer) BinaryDeserializeArray(arr *Array, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Версия
	var version uint32
	if err := binary.Read(file, binary.LittleEndian, &version); err != nil {
		return err
	}
	if version != 1 {
		return fmt.Errorf("unsupported version: %d", version)
	}

	// Имя
	name, err := s.readString(file)
	if err != nil {
		return err
	}

	// Емкость и размер (пропускаем)
	if _, err := s.readInt(file); err != nil {
		return err
	}
	if _, err := s.readInt(file); err != nil {
		return err
	}

	// Количество элементов
	elemCount, err := s.readInt(file)
	if err != nil {
		return err
	}

	// Чтение элементов
	data := make([]string, 0, elemCount)
	for i := 0; i < elemCount; i++ {
		str, err := s.readString(file)
		if err != nil {
			return err
		}
		data = append(data, str)
	}

	arr.Free()
	arr.name = name
	return arr.fromSlice(data)
}

func (s *Serializer) TextSerializeArray(arr *Array, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "NAME:%s\n", arr.name)
	fmt.Fprintf(file, "CAPACITY:%d\n", arr.capacity)
	fmt.Fprintf(file, "SIZE:%d\n", arr.size)

	data := arr.toSlice()
	fmt.Fprintf(file, "ELEMENTS:%d\n", len(data))
	
	for _, str := range data {
		fmt.Fprintf(file, "%s\n", str)
	}

	return nil
}

func (s *Serializer) TextDeserializeArray(arr *Array, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var name string
	fmt.Fscanf(file, "NAME:%s\n", &name)

	var capacity, size int
	fmt.Fscanf(file, "CAPACITY:%d\n", &capacity)
	fmt.Fscanf(file, "SIZE:%d\n", &size)

	var elemCount int
	fmt.Fscanf(file, "ELEMENTS:%d\n", &elemCount)

	data := make([]string, 0, elemCount)
	for i := 0; i < elemCount; i++ {
		var str string
		fmt.Fscanf(file, "%s\n", &str)
		data = append(data, str)
	}

	arr.Free()
	arr.name = name
	return arr.fromSlice(data)
}

// ==================== SINGLY LINKED LIST ====================

func (s *Serializer) BinarySerializeSinglyLinkedList(list *SinglyLinkedList, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Версия
	if err := binary.Write(file, binary.LittleEndian, uint32(1)); err != nil {
		return err
	}

	// Размер
	if err := s.writeInt(file, list.size); err != nil {
		return err
	}

	// Элементы
	data := list.ToVector()
	for _, str := range data {
		if err := s.writeString(file, str); err != nil {
			return err
		}
	}

	return nil
}

func (s *Serializer) BinaryDeserializeSinglyLinkedList(list *SinglyLinkedList, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Версия
	var version uint32
	if err := binary.Read(file, binary.LittleEndian, &version); err != nil {
		return err
	}
	if version != 1 {
		return fmt.Errorf("unsupported version: %d", version)
	}

	// Размер
	size, err := s.readInt(file)
	if err != nil {
		return err
	}

	// Чтение элементов
	data := make([]string, 0, size)
	for i := 0; i < size; i++ {
		str, err := s.readString(file)
		if err != nil {
			return err
		}
		data = append(data, str)
	}

	list.FromVector(data)
	return nil
}

func (s *Serializer) TextSerializeSinglyLinkedList(list *SinglyLinkedList, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "SIZE:%d\n", list.size)

	data := list.ToVector()
	for _, str := range data {
		fmt.Fprintf(file, "%s\n", str)
	}

	return nil
}

func (s *Serializer) TextDeserializeSinglyLinkedList(list *SinglyLinkedList, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var size int
	fmt.Fscanf(file, "SIZE:%d\n", &size)

	data := make([]string, 0, size)
	for i := 0; i < size; i++ {
		var str string
		fmt.Fscanf(file, "%s\n", &str)
		data = append(data, str)
	}

	list.FromVector(data)
	return nil
}

// ==================== DOUBLY LINKED LIST ====================

func (s *Serializer) BinarySerializeDoublyLinkedList(list *DoublyLinkedList, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Версия
	if err := binary.Write(file, binary.LittleEndian, uint32(1)); err != nil {
		return err
	}

	// Размер
	if err := s.writeInt(file, list.size); err != nil {
		return err
	}

	// Элементы
	data := list.ToVector()
	for _, str := range data {
		if err := s.writeString(file, str); err != nil {
			return err
		}
	}

	return nil
}

func (s *Serializer) BinaryDeserializeDoublyLinkedList(list *DoublyLinkedList, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Версия
	var version uint32
	if err := binary.Read(file, binary.LittleEndian, &version); err != nil {
		return err
	}
	if version != 1 {
		return fmt.Errorf("unsupported version: %d", version)
	}

	// Размер
	size, err := s.readInt(file)
	if err != nil {
		return err
	}

	// Чтение элементов
	data := make([]string, 0, size)
	for i := 0; i < size; i++ {
		str, err := s.readString(file)
		if err != nil {
			return err
		}
		data = append(data, str)
	}

	list.FromVector(data)
	return nil
}

func (s *Serializer) TextSerializeDoublyLinkedList(list *DoublyLinkedList, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "SIZE:%d\n", list.size)

	data := list.ToVector()
	for _, str := range data {
		fmt.Fprintf(file, "%s\n", str)
	}

	return nil
}

func (s *Serializer) TextDeserializeDoublyLinkedList(list *DoublyLinkedList, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var size int
	fmt.Fscanf(file, "SIZE:%d\n", &size)

	data := make([]string, 0, size)
	for i := 0; i < size; i++ {
		var str string
		fmt.Fscanf(file, "%s\n", &str)
		data = append(data, str)
	}

	list.FromVector(data)
	return nil
}

// ==================== HASH TABLE ====================

func (s *Serializer) BinarySerializeHashTable(ht *OpenAddressingHashTable, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Версия
	if err := binary.Write(file, binary.LittleEndian, uint32(1)); err != nil {
		return err
	}

	// Емкость и размер
	if err := s.writeInt(file, ht.capacity); err != nil {
		return err
	}
	if err := s.writeInt(file, ht.size); err != nil {
		return err
	}

	// Элементы
	data := ht.ToVector()
	if err := s.writeInt(file, len(data)); err != nil {
		return err
	}

	for _, pair := range data {
		// Запись ключа (rune) как int32
		if err := binary.Write(file, binary.LittleEndian, int32(pair.key)); err != nil {
			return err
		}
		// Запись значения
		if err := binary.Write(file, binary.LittleEndian, int64(pair.value)); err != nil {
			return err
		}
	}

	return nil
}

func (s *Serializer) BinaryDeserializeHashTable(ht *OpenAddressingHashTable, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Версия
	var version uint32
	if err := binary.Read(file, binary.LittleEndian, &version); err != nil {
		return err
	}
	if version != 1 {
		return fmt.Errorf("unsupported version: %d", version)
	}

	// Емкость и размер
	capacity, err := s.readInt(file)
	if err != nil {
		return err
	}
	_, err = s.readInt(file)
	if err != nil {
		return err
	}

	// Количество элементов
	elemCount, err := s.readInt(file)
	if err != nil {
		return err
	}

	// Чтение элементов
	data := make([]Pair, 0, elemCount)
	for i := 0; i < elemCount; i++ {
		var key int32
		if err := binary.Read(file, binary.LittleEndian, &key); err != nil {
			return err
		}
		
		var value int64
		if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
			return err
		}
		
		data = append(data, Pair{key: rune(key), value: int(value)})
	}

	ht.Clear()
	*ht = *NewOpenAddressingHashTable(capacity)
	for _, pair := range data {
		ht.Insert(pair.key, pair.value)
	}

	return nil
}

func (s *Serializer) TextSerializeHashTable(ht *OpenAddressingHashTable, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "CAPACITY:%d\n", ht.capacity)
	fmt.Fprintf(file, "SIZE:%d\n", ht.size)

	data := ht.ToVector()
	fmt.Fprintf(file, "ELEMENTS:%d\n", len(data))
	
	for _, pair := range data {
		fmt.Fprintf(file, "%c:%d\n", pair.key, pair.value)
	}

	return nil
}

func (s *Serializer) TextDeserializeHashTable(ht *OpenAddressingHashTable, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var capacity, size int
	fmt.Fscanf(file, "CAPACITY:%d\n", &capacity)
	fmt.Fscanf(file, "SIZE:%d\n", &size)

	var elemCount int
	fmt.Fscanf(file, "ELEMENTS:%d\n", &elemCount)

	data := make([]Pair, 0, elemCount)
	for i := 0; i < elemCount; i++ {
		var keyChar byte
		var value int
		fmt.Fscanf(file, "%c:%d\n", &keyChar, &value)
		data = append(data, Pair{key: rune(keyChar), value: value})
	}

	ht.Clear()
	*ht = *NewOpenAddressingHashTable(capacity)
	for _, pair := range data {
		ht.Insert(pair.key, pair.value)
	}

	return nil
}

// ==================== BINARY TREE ====================

func (s *Serializer) BinarySerializeTree(tree *FullBinaryTree, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Версия
	if err := binary.Write(file, binary.LittleEndian, uint32(1)); err != nil {
		return err
	}

	// Элементы
	data := tree.ToVector()
	if err := s.writeInt(file, len(data)); err != nil {
		return err
	}

	for _, val := range data {
		if err := binary.Write(file, binary.LittleEndian, int64(val)); err != nil {
			return err
		}
	}

	return nil
}

func (s *Serializer) BinaryDeserializeTree(tree *FullBinaryTree, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Версия
	var version uint32
	if err := binary.Read(file, binary.LittleEndian, &version); err != nil {
		return err
	}
	if version != 1 {
		return fmt.Errorf("unsupported version: %d", version)
	}

	// Количество элементов
	elemCount, err := s.readInt(file)
	if err != nil {
		return err
	}

	// Чтение элементов
	data := make([]int, 0, elemCount)
	for i := 0; i < elemCount; i++ {
		var val int64
		if err := binary.Read(file, binary.LittleEndian, &val); err != nil {
			return err
		}
		data = append(data, int(val))
	}

	tree.FromVector(data)
	return nil
}

func (s *Serializer) TextSerializeTree(tree *FullBinaryTree, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	data := tree.ToVector()
	fmt.Fprintf(file, "ELEMENTS:%d\n", len(data))
	
	for i, val := range data {
		if i > 0 {
			fmt.Fprintf(file, " ")
		}
		fmt.Fprintf(file, "%d", val)
	}
	fmt.Fprintf(file, "\n")

	return nil
}

func (s *Serializer) TextDeserializeTree(tree *FullBinaryTree, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var elemCount int
	fmt.Fscanf(file, "ELEMENTS:%d\n", &elemCount)

	data := make([]int, 0, elemCount)
	for i := 0; i < elemCount; i++ {
		var val int
		fmt.Fscanf(file, "%d", &val)
		data = append(data, val)
	}

	tree.FromVector(data)
	return nil
}

// ==================== СТЕК СЕРИАЛИЗАЦИЯ ====================

// BinarySerializeStack сохраняет Stack в бинарный файл
func (s *Serializer) BinarySerializeStack(stack *Stack, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Версия
	if err := binary.Write(file, binary.LittleEndian, uint32(1)); err != nil {
		return err
	}

	// Размер
	if err := s.writeInt(file, stack.Size()); err != nil {
		return err
	}

	// Элементы от вершины к основанию
	elements := stack.ToVector() // уже от вершины
	for i := 0; i < len(elements); i++ {
		if err := s.writeString(file, elements[i]); err != nil {
			return err
		}
	}

	return nil
}

// BinaryDeserializeStack загружает Stack из бинарного файла
func (s *Serializer) BinaryDeserializeStack(stack *Stack, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Версия
	var version uint32
	if err := binary.Read(file, binary.LittleEndian, &version); err != nil {
		return err
	}
	if version != 1 {
		return fmt.Errorf("unsupported version: %d", version)
	}

	// Размер
	size, err := s.readInt(file)
	if err != nil {
		return err
	}

	// Чтение элементов
	elements := make([]string, 0, size)
	for i := 0; i < size; i++ {
		str, err := s.readString(file)
		if err != nil {
			return err
		}
		elements = append(elements, str)
	}

	// Загружаем
	stack.FromVector(elements)
	return nil
}

// TextSerializeStack сохраняет Stack в текстовый файл
func (s *Serializer) TextSerializeStack(stack *Stack, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "STACK\n")
	fmt.Fprintf(file, "SIZE:%d\n", stack.Size())

	elements := stack.ToVector() // от вершины к основанию
	fmt.Fprintf(file, "ELEMENTS:%d\n", len(elements))
	
	for i := 0; i < len(elements); i++ {
		fmt.Fprintf(file, "%s\n", elements[i])
	}

	return nil
}

// TextDeserializeStack загружает Stack из текстового файла
func (s *Serializer) TextDeserializeStack(stack *Stack, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var typeStr string
	fmt.Fscanf(file, "%s\n", &typeStr)
	if typeStr != "STACK" {
		return fmt.Errorf("invalid file type: %s", typeStr)
	}

	var size int
	fmt.Fscanf(file, "SIZE:%d\n", &size)

	var elemCount int
	fmt.Fscanf(file, "ELEMENTS:%d\n", &elemCount)

	elements := make([]string, 0, elemCount)
	for i := 0; i < elemCount; i++ {
		var str string
		fmt.Fscanf(file, "%s\n", &str)
		elements = append(elements, str)
	}

	stack.FromVector(elements)
	return nil
}

// SimplifiedTextSerializeStack - упрощенная версия
func (s *Serializer) SimplifiedTextSerializeStack(stack *Stack, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	elements := stack.ToVector()
	fmt.Fprintf(file, "%d\n", len(elements))
	
	for i := 0; i < len(elements); i++ {
		fmt.Fprintf(file, "%s\n", elements[i])
	}

	return nil
}

// SimplifiedTextDeserializeStack - упрощенная версия загрузки
func (s *Serializer) SimplifiedTextDeserializeStack(stack *Stack, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var size int
	fmt.Fscanf(file, "%d\n", &size)

	elements := make([]string, 0, size)
	for i := 0; i < size; i++ {
		var str string
		fmt.Fscanf(file, "%s\n", &str)
		elements = append(elements, str)
	}

	stack.FromVector(elements)
	return nil
}

// CompactBinarySerializeStack - компактная бинарная сериализация
func (s *Serializer) CompactBinarySerializeStack(stack *Stack, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	if err := s.writeInt(file, stack.Size()); err != nil {
		return err
	}

	elements := stack.ToVector()
	for i := 0; i < len(elements); i++ {
		if err := s.writeString(file, elements[i]); err != nil {
			return err
		}
	}

	return nil
}

// CompactBinaryDeserializeStack - компактная бинарная загрузка
func (s *Serializer) CompactBinaryDeserializeStack(stack *Stack, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	size, err := s.readInt(file)
	if err != nil {
		return err
	}

	elements := make([]string, 0, size)
	for i := 0; i < size; i++ {
		str, err := s.readString(file)
		if err != nil {
			return err
		}
		elements = append(elements, str)
	}

	stack.FromVector(elements)
	return nil
}