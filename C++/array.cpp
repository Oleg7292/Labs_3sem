// array.cpp
#include "array.h"
#include <iostream>
using namespace std;

// Конструктор по умолчанию
Array::Array() {
    volume = 10;  // начальная вместимость
    arr = new string[volume];
    arr_size = 0;
}

// Деструктор
Array::~Array() {
    delete[] arr;
}

// Показать все элементы массива
void Array::ShowArray() const {
    cout << "Array elements: ";
    for (size_t i = 0; i < arr_size; i++) {
        cout << "\"" << arr[i] << "\" ";
    }
    cout << endl;
}

// Добавить элемент в конец
void Array::addToEnd(string value) {
    if (arr_size >= volume) {
        // Увеличиваем массив в 2 раза
        size_t newVolume = volume * 2;
        string* newArr = new string[newVolume];
        
        // Копируем старые элементы
        for (size_t i = 0; i < arr_size; i++) {
            newArr[i] = arr[i];
        }
        
        // Удаляем старый массив и заменяем новым
        delete[] arr;
        arr = newArr;
        volume = newVolume;
    }
    
    arr[arr_size++] = value;
}

// Добавить элемент по индексу
void Array::add(size_t index, string value) {
    if (index > arr_size) {
        cout << "Index out of range" << endl;
        return;
    }
    
    if (arr_size >= volume) {
        // Увеличиваем массив
        size_t newVolume = volume * 2;
        string* newArr = new string[newVolume];
        
        // Копируем элементы до индекса
        for (size_t i = 0; i < index; i++) {
            newArr[i] = arr[i];
        }
        
        // Вставляем новый элемент
        newArr[index] = value;
        
        // Копируем остальные элементы
        for (size_t i = index; i < arr_size; i++) {
            newArr[i + 1] = arr[i];
        }
        
        delete[] arr;
        arr = newArr;
        volume = newVolume;
        arr_size++;
    } else {
        // Сдвигаем элементы вправо
        for (size_t i = arr_size; i > index; i--) {
            arr[i] = arr[i - 1];
        }
        arr[index] = value;
        arr_size++;
    }
}

// Получить элемент по индексу
string Array::getIndex(size_t index) {
    if (index >= arr_size) {
        cout << "Index out of range" << endl;
        return "";
    }
    return arr[index];
}

// Удалить элемент по индексу
void Array::remove(size_t index) {
    if (index >= arr_size) {
        cout << "Index out of range" << endl;
        return;
    }
    
    // Сдвигаем элементы влево
    for (size_t i = index; i < arr_size - 1; i++) {
        arr[i] = arr[i + 1];
    }
    arr_size--;
}

// Заменить элемент по индексу
void Array::replace(size_t index, string value) {
    if (index >= arr_size) {
        cout << "Index out of range" << endl;
        return;
    }
    arr[index] = value;
}

// Получить текущий размер массива
size_t Array::getSize() const {
    return arr_size;
}

// Для сериализации - преобразовать в вектор
vector<string> Array::toVector() const {
    vector<string> result;
    for (size_t i = 0; i < arr_size; i++) {
        result.push_back(arr[i]);
    }
    return result;
}

// Для сериализации - загрузить из вектора
void Array::fromVector(const vector<string>& elements) {
    // Очищаем текущий массив
    delete[] arr;
    
    // Создаем новый массив нужного размера
    arr_size = elements.size();
    volume = arr_size + 5;  // небольшой запас
    
    arr = new string[volume];
    
    // Копируем элементы
    for (size_t i = 0; i < arr_size; i++) {
        arr[i] = elements[i];
    }
}