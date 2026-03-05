#include "doubly_linked_list.h"

DoublyLinkedList::DoublyLinkedList() : head(nullptr), tail(nullptr), list_size(0) {}

DoublyLinkedList::~DoublyLinkedList() {
    clear();
}

void DoublyLinkedList::push_back(const string& value) {
    Node* new_node = new Node(value);
    if (!head) {
        head = tail = new_node;
    } else {
        tail->next = new_node;
        new_node->prev = tail;
        tail = new_node;
    }
    list_size++;
}

void DoublyLinkedList::push_front(const string& value) {
    Node* new_node = new Node(value);
    if (!head) {
        head = tail = new_node;
    } else {
        new_node->next = head;
        head->prev = new_node;
        head = new_node;
    }
    list_size++;
}

void DoublyLinkedList::pop_front() {
    if (!head) return;
    
    Node* temp = head;
    head = head->next;
    if (head) {
        head->prev = nullptr;
    } else {
        tail = nullptr;
    }
    delete temp;
    list_size--;
}

void DoublyLinkedList::pop_back() {
    if (!tail) return;
    
    Node* temp = tail;
    tail = tail->prev;
    if (tail) {
        tail->next = nullptr;
    } else {
        head = nullptr;
    }
    delete temp;
    list_size--;
}

void DoublyLinkedList::insert(size_t index, const string& value) {
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
    new_node->prev = current;
    new_node->next = current->next;
    
    if (current->next) {
        current->next->prev = new_node;
    }
    current->next = new_node;
    
    // Если вставляем в конец, обновляем tail
    if (new_node->next == nullptr) {
        tail = new_node;
    }
    
    list_size++;
}

void DoublyLinkedList::remove(size_t index) {
    if (index >= list_size) return;
    
    if (index == 0) {
        pop_front();
        return;
    }
    if (index == list_size - 1) {
        pop_back();
        return;
    }

    // Находим элемент для удаления
    Node* current = head;
    for (size_t i = 0; i < index; ++i) {
        current = current->next;
    }
    
    // Удаляем current
    if (current->prev) {
        current->prev->next = current->next;
    }
    if (current->next) {
        current->next->prev = current->prev;
    }
    
    delete current;
    list_size--;
}

string& DoublyLinkedList::get(size_t index) {
    if (index >= list_size) throw out_of_range("Index out of range");
    
    Node* current = head;
    for (size_t i = 0; i < index; ++i) {
        current = current->next;
    }
    return current->data;
}

string& DoublyLinkedList::get_front() {
    if (!head) throw out_of_range("List is empty");
    return head->data;
}

string& DoublyLinkedList::get_back() {
    if (!tail) throw out_of_range("List is empty");
    return tail->data;
}

bool DoublyLinkedList::isEmpty() const {
    return list_size == 0;
}

size_t DoublyLinkedList::size() const {
    return list_size;
}

void DoublyLinkedList::clear() {
    while (!isEmpty()) {
        pop_front();
    }
}

vector<string> DoublyLinkedList::toVector() const {
    vector<string> result;
    Node* current = head;
    while (current) {
        result.push_back(current->data);
        current = current->next;
    }
    return result;
}

void DoublyLinkedList::fromVector(const vector<string>& elements) {
    clear();
    for (const auto& element : elements) {
        push_back(element);
    }
}