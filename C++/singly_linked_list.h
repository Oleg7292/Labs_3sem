#ifndef SINGLY_LINKED_LIST_H
#define SINGLY_LINKED_LIST_H

#include "data_structures.h"

class SinglyLinkedList {
private:
    struct Node {
        string data;
        Node* next;
        Node(const string& value) : data(value), next(nullptr) {}
    };
    
    Node* head;
    Node* tail;
    size_t list_size;

public:
    SinglyLinkedList();
    ~SinglyLinkedList();
    void push_back(const string& value);
    void push_front(const string& value);
    void pop_front();
    void insert(size_t index, const string& value);
    void remove(size_t index);
    string& get(size_t index);
    bool isEmpty() const;
    size_t size() const;
    void clear();
    
    // Для сериализации
    vector<string> toVector() const;
    void fromVector(const vector<string>& elements);
};

#endif