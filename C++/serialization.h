#ifndef SERIALIZATION_H
#define SERIALIZATION_H

#include "data_structures.h"
#include "stack.h"
#include "array.h"
#include "hash_table.h"
#include "singly_linked_list.h"
#include "doubly_linked_list.h"
#include <fstream>
#include <sstream>

class Serializer {
public:
    // Для STACK (уже есть)
    static void binarySerialize(const Stack& obj, const std::string& filename);
    static void binaryDeserialize(Stack& obj, const std::string& filename);
    static void textSerialize(const Stack& obj, const std::string& filename);
    static void textDeserialize(Stack& obj, const std::string& filename);
    
    // Для ARRAY
    static void binarySerialize(const Array& obj, const std::string& filename);
    static void binaryDeserialize(Array& obj, const std::string& filename);
    static void textSerialize(const Array& obj, const std::string& filename);
    static void textDeserialize(Array& obj, const std::string& filename);
    
    // Для HASH TABLE
    static void binarySerialize(const OpenAddressingHashTable& obj, const std::string& filename);
    static void binaryDeserialize(OpenAddressingHashTable& obj, const std::string& filename);
    static void textSerialize(const OpenAddressingHashTable& obj, const std::string& filename);
    static void textDeserialize(OpenAddressingHashTable& obj, const std::string& filename);
    
    // Для SINGLY LINKED LIST
    static void binarySerialize(const SinglyLinkedList& obj, const std::string& filename);
    static void binaryDeserialize(SinglyLinkedList& obj, const std::string& filename);
    static void textSerialize(const SinglyLinkedList& obj, const std::string& filename);
    static void textDeserialize(SinglyLinkedList& obj, const std::string& filename);
    
    // Для DOUBLY LINKED LIST
    static void binarySerialize(const DoublyLinkedList& obj, const std::string& filename);
    static void binaryDeserialize(DoublyLinkedList& obj, const std::string& filename);
    static void textSerialize(const DoublyLinkedList& obj, const std::string& filename);
    static void textDeserialize(DoublyLinkedList& obj, const std::string& filename);
};

#endif