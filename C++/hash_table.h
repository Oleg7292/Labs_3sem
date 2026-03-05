#ifndef HASH_TABLE_H
#define HASH_TABLE_H

#include <vector>
#include <utility>  // для pair
#include <string>
using namespace std;

class OpenAddressingHashTable {
private:
    vector<pair<char, int>> table;  // хеш-таблица
    vector<bool> occupied;           // флаги занятости
    vector<bool> deleted;            // **ОБЯЗАТЕЛЬНО ДОБАВИТЬ** флаги удаленных элементов
    int capacity;                     // вместимость
    int size;                         // текущий размер

    int hash(char key, int attempt) const;  // хеш-функция (const для использования в const-методах)

public:
    // Конструктор с параметром вместимости (по умолчанию 256)
    OpenAddressingHashTable(int cap = 256);
    
    // Основные операции
    void insert(char key, int value);
    bool search(char key, int& value) const;  // const метод
    bool remove(char key);
    
    // Вспомогательные методы
    void clear();
    bool isEmpty() const;
    int getSize() const;
    int getCapacity() const;
    
    // Для сериализации (используются в serialization.cpp)
    vector<pair<char, int>> toVector() const;
    void fromVector(const vector<pair<char, int>>& elements);
};

#endif