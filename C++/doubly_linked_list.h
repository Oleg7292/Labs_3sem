#ifndef DOUBLY_LINKED_LIST_H
#define DOUBLY_LINKED_LIST_H

#include "data_structures.h"

class DoublyLinkedList {
private:
    struct Node {
        string data;
        Node* prev;
        Node* next;
        Node(const string& value) : data(value), prev(nullptr), next(nullptr) {}
    };
    
    Node* head;
    Node* tail;
    size_t list_size;

public:
    DoublyLinkedList();
    ~DoublyLinkedList();
    void push_back(const string& value);
    void push_front(const string& value);
    void pop_front();
    void pop_back();
    void insert(size_t index, const string& value);
    void remove(size_t index);
    string& get(size_t index);
    string& get_front();
    string& get_back();
    bool isEmpty() const;
    size_t size() const;
    void clear();
    
    // Для сериализации
    vector<string> toVector() const;
    void fromVector(const vector<string>& elements);
};

#endif