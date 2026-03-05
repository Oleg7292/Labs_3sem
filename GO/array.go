package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Array struct {
	name     string
	capacity int
	size     int
	data     uintptr
}

func malloc(size int) uintptr {
	ptr := C.malloc(C.size_t(size))
	if ptr == nil {
		panic("malloc failed")
	}
	return uintptr(ptr)
}

func free(ptr uintptr) {
	if ptr != 0 {
		C.free(unsafe.Pointer(ptr))
	}
}

func NewArray(name string, capacity int) *Array {
	sizeOfString := unsafe.Sizeof("")
	data := malloc(int(sizeOfString) * capacity)

	return &Array{
		name:     name,
		capacity: capacity,
		size:     0,
		data:     data,
	}
}

func (a *Array) Free() {
	if a.data != 0 {
		free(a.data)
		a.data = 0
		a.size = 0
		a.capacity = 0
	}
}

func (a *Array) getElementPtr(index int) unsafe.Pointer {
	sizeOfString := unsafe.Sizeof("")
	return unsafe.Pointer(a.data + uintptr(index)*sizeOfString)
}

func (a *Array) getString(index int) string {
	if index < 0 || index >= a.size {
		return ""
	}
	ptr := a.getElementPtr(index)
	// Безопасное приведение через промежуточный указатель
	return *(*string)(ptr)
}

func (a *Array) setString(index int, value string) {
	if index < 0 || index >= a.capacity {
		return
	}
	ptr := a.getElementPtr(index)
	// Безопасное приведение через промежуточный указатель
	*(*string)(ptr) = value
}

func (a *Array) resize() {
	newCapacity := a.capacity * 2
	sizeOfString := unsafe.Sizeof("")

	newData := malloc(int(sizeOfString) * newCapacity)

	// Копируем существующие элементы
	for i := 0; i < a.size; i++ {
		oldPtr := a.getElementPtr(i)
		newPtr := unsafe.Pointer(newData + uintptr(i)*sizeOfString)
		// Копируем через присваивание
		*(*string)(newPtr) = *(*string)(oldPtr)
	}

	free(a.data)
	a.data = newData
	a.capacity = newCapacity
}

func (a *Array) PushBack(value string) {
	if a.size >= a.capacity {
		a.resize()
	}
	a.setString(a.size, value)
	a.size++
}

func (a *Array) Insert(index int, value string) bool {
	if index < 0 || index > a.size {
		fmt.Println("incorrect")
		return false
	}

	if a.size >= a.capacity {
		a.resize()
	}

	// Сдвигаем элементы вправо
	for i := a.size; i > index; i-- {
		prev := a.getString(i - 1)
		a.setString(i, prev)
	}

	a.setString(index, value)
	a.size++
	return true
}

func (a *Array) Get(index int) string {
	if index < 0 || index >= a.size {
		fmt.Println("incorrect")
		return ""
	}
	return a.getString(index)
}

func (a *Array) Remove(index int) bool {
	if index < 0 || index >= a.size {
		fmt.Println("incorrect")
		return false
	}

	// Сдвигаем элементы влево
	for i := index; i < a.size-1; i++ {
		next := a.getString(i + 1)
		a.setString(i, next)
	}

	// Очищаем последний элемент
	a.setString(a.size-1, "")
	a.size--
	return true
}

func (a *Array) Replace(index int, value string) bool {
	if index < 0 || index >= a.size {
		fmt.Println("incorrect")
		return false
	}
	a.setString(index, value)
	return true
}

func (a *Array) Length() int {
	return a.size
}

func (a *Array) Size() int {
	return a.size
}

func (a *Array) Print() {
	fmt.Printf("%s = [ ", a.name)
	for i := 0; i < a.size; i++ {
		fmt.Printf("\"%s\"", a.getString(i))
		if i < a.size-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(" ]")
}

// Методы для тестирования
func (a *Array) Equals(other *Array) bool {
	if a.name != other.name || a.size != other.size {
		return false
	}

	for i := 0; i < a.size; i++ {
		if a.getString(i) != other.getString(i) {
			return false
		}
	}

	return true
}

func (a *Array) Copy() *Array {
	newArr := NewArray(a.name, a.capacity)
	for i := 0; i < a.size; i++ {
		newArr.PushBack(a.getString(i))
	}
	return newArr
}

func (a *Array) toSlice() []string {
	result := make([]string, a.size)
	for i := 0; i < a.size; i++ {
		result[i] = a.getString(i)
	}
	return result
}

func (a *Array) fromSlice(data []string) error {
	a.Free()
	
	a.name = "" // или сохранить старое имя?
	
	capacity := len(data)
	if capacity == 0 {
		capacity = 1
	}

	sizeOfString := unsafe.Sizeof("")
	a.data = malloc(int(sizeOfString) * capacity)
	a.capacity = capacity
	a.size = 0

	for i := 0; i < len(data); i++ {
		ptr := a.getElementPtr(i)
		*(*string)(ptr) = data[i]
		a.size++
	}
	
	// Если нужно сохранить имя:
	// a.name = oldName
	
	return nil
}