#include "hash_table.h"
#include "serialization.h"  // Добавляем для Serializer
#include <fstream>
#include <iostream>
#include <vector>
#include <cmath>
using namespace std;

OpenAddressingHashTable::OpenAddressingHashTable(int cap) : capacity(cap), size(0) {
    table.resize(capacity);
    occupied.resize(capacity, false);
    deleted.resize(capacity, false);  
}

vector<pair<char, int>> OpenAddressingHashTable::toVector() const {
    vector<pair<char, int>> result;
    for (int i = 0; i < capacity; i++) {
        if (occupied[i] && !deleted[i]) {  // Учитываем флаг deleted
            result.push_back(table[i]);
        }
    }
    return result;
}

void OpenAddressingHashTable::fromVector(const vector<pair<char, int>>& elements) {
    clear();
    for (const auto& element : elements) {
        insert(element.first, element.second);
    }
}

void OpenAddressingHashTable::clear() {
    for (int i = 0; i < capacity; i++) {
        occupied[i] = false;
        deleted[i] = false;  // Сбрасываем флаг deleted
    }
    size = 0;
}

bool OpenAddressingHashTable::isEmpty() const {
    return size == 0;
}

int OpenAddressingHashTable::hash(char key, int attempt) const {
    // Используем const_cast для совместимости с const-методом
    // Но лучше сделать метод не const, либо передавать capacity как параметр
    int h1 = abs(static_cast<int>(key)) % capacity;
    
    // Для двойного хеширования нужен второй хеш
    int h2 = 1 + (abs(static_cast<int>(key)) % (capacity - 1));
    
    return (h1 + attempt * h2) % capacity;
}

void OpenAddressingHashTable::insert(char key, int value) {
    // Проверяем, не нужно ли увеличить таблицу (коэффициент загрузки 0.7)
    if (size >= capacity * 0.7) {
        // Здесь можно вызвать resize(), если он реализован
        // resize();
    }
    
    // Сначала ищем, есть ли уже такой ключ
    for (int attempt = 0; attempt < capacity; attempt++) {
        int index = hash(key, attempt);
        
        // Если нашли такой же ключ и он не удален
        if (occupied[index] && !deleted[index] && table[index].first == key) {
            table[index].second = value;  // Обновляем значение
            return;
        }
        
        // Если нашли пустое место или удаленную ячейку
        if (!occupied[index] || deleted[index]) {
            table[index] = {key, value};
            occupied[index] = true;
            deleted[index] = false;
            size++;
            return;
        }
    }
    
    // Если дошли сюда, значит таблица переполнена
    // Можно увеличить размер
    // resize();
    // insert(key, value); // Повторная попытка
}

bool OpenAddressingHashTable::search(char key, int& value) const {
    for (int attempt = 0; attempt < capacity; attempt++) {
        int index = hash(key, attempt);
        
        if (occupied[index] && !deleted[index] && table[index].first == key) {
            value = table[index].second;
            return true;
        }
        
        if (!occupied[index]) {
            return false; // Нашли пустую ячейку - ключа нет
        }
    }
    return false;
}

bool OpenAddressingHashTable::remove(char key) {
    for (int attempt = 0; attempt < capacity; attempt++) {
        int index = hash(key, attempt);
        
        if (occupied[index] && !deleted[index] && table[index].first == key) {
            deleted[index] = true; // Помечаем как удаленную
            size--;
            return true;
        }
        
        if (!occupied[index]) {
            return false; // Нашли пустую ячейку - ключа нет
        }
    }
    return false;
}

int OpenAddressingHashTable::getSize() const {
    return size;
}

int OpenAddressingHashTable::getCapacity() const {
    return capacity;
}