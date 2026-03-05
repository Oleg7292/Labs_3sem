#include "singly_linked_list.h"

SinglyLinkedList::SinglyLinkedList() : head(nullptr), tail(nullptr), list_size(0) {}

SinglyLinkedList::~SinglyLinkedList() {
    clear();
}

void SinglyLinkedList::push_back(const string& value) {
    Node* new_node = new Node(value);
    if (!head) {
        head = tail = new_node;
    } else {
        tail->next = new_node;
        tail = new_node;
    }
    list_size++;
}

void SinglyLinkedList::push_front(const string& value) {
    Node* new_node = new Node(value);
    if (!head) {
        head = tail = new_node;
    } else {
        new_node->next = head;
        head = new_node;
    }
    list_size++;
}

void SinglyLinkedList::pop_front() {
    if (!head) return;
    Node* temp = head;
    head = head->next;
    delete temp;
    list_size--;
    if (!head) tail = nullptr;
}

void SinglyLinkedList::insert(size_t index, const string& value) {
    if (index > list_size) return;
    
    if (index == 0) {
        push_front(value);
        return;
    }
    if (index == list_size) {
        push_back(value);
        return;
    }

    Node* current = head;
    for (size_t i = 0; i < index - 1; ++i) {
        current = current->next;
    }
    
    Node* new_node = new Node(value);
    new_node->next = current->next;
    current->next = new_node;
    list_size++;
}

void SinglyLinkedList::remove(size_t index) {
    if (index >= list_size) return;
    
    if (index == 0) {
        pop_front();
        return;
    }

    Node* current = head;
    for (size_t i = 0; i < index - 1; ++i) {
        current = current->next;
    }
    
    Node* to_delete = current->next;
    current->next = to_delete->next;
    
    if (to_delete == tail) {
        tail = current;
    }
    
    delete to_delete;
    list_size--;
}

string& SinglyLinkedList::get(size_t index) {
    if (index >= list_size) throw out_of_range("Index out of range");
    
    Node* current = head;
    for (size_t i = 0; i < index; ++i) {
        current = current->next;
    }
    return current->data;
}

bool SinglyLinkedList::isEmpty() const {
    return list_size == 0;
}

size_t SinglyLinkedList::size() const {
    return list_size;
}

void SinglyLinkedList::clear() {
    while (!isEmpty()) {
        pop_front();
    }
}

vector<string> SinglyLinkedList::toVector() const {
    vector<string> result;
    Node* current = head;
    while (current) {
        result.push_back(current->data);
        current = current->next;
    }
    return result;
}

void SinglyLinkedList::fromVector(const vector<string>& elements) {
    clear();
    for (const auto& element : elements) {
        push_back(element);
    }
}