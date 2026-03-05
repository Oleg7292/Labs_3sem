#ifndef STACK_H
#define STACK_H

#include "data_structures.h"

class Stack {
private:
    string* data;
    size_t capacity;
    int head;

public:
    Stack(size_t size = 30);
    ~Stack();
    void push(string value);
    string pop();
    string peek();
    bool isEmpty();
    size_t size();
    
    // Для сериализации
    vector<string> toVector() const;
    void fromVector(const vector<string>& elements);
};

#endif