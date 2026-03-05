#include "stack.h"

Stack::Stack(size_t size) : capacity(size), head(-1) {
    data = new string[capacity];
}

Stack::~Stack() { 
    delete[] data; 
}

void Stack::push(string value) {
    if (head == (int)capacity - 1) {
        throw overflow_error("Stack is full");
    }
    data[++head] = value;
}

string Stack::pop() {
    if (head == -1) {
        throw underflow_error("Stack is empty");
    }
    return data[head--];
}

string Stack::peek() {
    if (head == -1) {
        throw underflow_error("Stack is empty");
    }
    return data[head];
}

bool Stack::isEmpty() { 
    return head == -1; 
}

size_t Stack::size() { 
    return head + 1; 
}

vector<string> Stack::toVector() const {
    vector<string> result;
    for (int i = 0; i <= head; ++i) {
        result.push_back(data[i]);
    }
    return result;
}

void Stack::fromVector(const vector<string>& elements) {
    delete[] data;
    capacity = elements.size() + 10;
    data = new string[capacity];
    head = -1;
    for (const auto& element : elements) {
        push(element);
    }
}