#include <iostream>
#include "data_structures.h"
#include "stack.h"
#include "array.h"
#include "singly_linked_list.h"
#include "doubly_linked_list.h"
#include "binary_tree.h"
#include "hash_table.h"
#include "serialization.h"

using namespace std;

void demoStack() {
    cout << "=== Stack Demo ===" << endl;
    Stack stack(5);
    
    stack.push("first");
    stack.push("second");
    stack.push("third");
    
    cout << "Stack size: " << stack.size() << endl;
    cout << "Top element: " << stack.peek() << endl;
    
    while (!stack.isEmpty()) {
        cout << "Popped: " << stack.pop() << endl;
    }
    cout << endl;
}

void demoArray() {
    cout << "=== Array Demo ===" << endl;
    Array arr;
    
    arr.addToEnd("apple");
    arr.addToEnd("banana");
    arr.addToEnd("cherry");
    arr.add(1, "blueberry");
    
    cout << "Array contents: ";
    arr.ShowArray();
    cout << "Array size: " << arr.getSize() << endl << endl;
}

void demoSinglyLinkedList() {
    cout << "=== Singly Linked List Demo ===" << endl;
    SinglyLinkedList list;
    
    list.push_back("first");
    list.push_back("third");
    list.insert(1, "second");
    
    cout << "List contents: ";
    for (size_t i = 0; i < list.size(); ++i) {
        cout << list.get(i) << " ";
    }
    cout << endl << "List size: " << list.size() << endl << endl;
}

void demoBinaryTree() {
    cout << "=== Binary Tree Demo ===" << endl;
    FullBinaryTree tree;
    
    tree.insert(5);
    tree.insert(3);
    tree.insert(7);
    tree.insert(1);
    tree.insert(4);
    tree.insert(6);
    tree.insert(8);
    
    cout << "Tree structure:" << endl;
    tree.print();
    
    cout << "Zigzag traversal:" << endl;
    tree.printZigZag();
    cout << endl;
}

int main() {
    try {
        demoStack();
        demoArray();
        demoSinglyLinkedList();
        demoBinaryTree();
        
        cout << "All demos completed successfully!" << endl;
        
    } catch (const exception& e) {
        cerr << "Error: " << e.what() << endl;
        return 1;
    }
    
    return 0;
}