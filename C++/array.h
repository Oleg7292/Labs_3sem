// array.h
#ifndef ARRAY_H
#define ARRAY_H

#include <string>
#include <vector>
using namespace std;

class Array {
private:
    string *arr;        // динамический массив строк
    size_t volume;       // вместимость массива (выделенная память)
    size_t arr_size;     // текущий размер массива (количество элементов)

public:
    // Конструктор и деструктор
    Array();                    // конструктор по умолчанию
    ~Array();                   // деструктор
    
    // Основные методы работы с массивом
    void ShowArray() const;                     // показать все элементы
    void addToEnd(string value);                 // добавить в конец
    void add(size_t index, string value);        // добавить по индексу
    string getIndex(size_t index);                // получить элемент по индексу
    void remove(size_t index);                    // удалить по индексу
    void replace(size_t index, string value);     // заменить по индексу
    size_t getSize() const;                        // получить текущий размер
    
    // Методы для сериализации
    vector<string> toVector() const;              // преобразовать в вектор
    void fromVector(const vector<string>& elements); // загрузить из вектора
};

#endif