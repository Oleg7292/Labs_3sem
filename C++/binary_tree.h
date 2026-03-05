#ifndef BINARY_TREE_H
#define BINARY_TREE_H

#include "data_structures.h"

class NodeT {
public:
    int data;
    NodeT* left;
    NodeT* right;
    NodeT(int value) : data(value), left(nullptr), right(nullptr) {}
};

class FullBinaryTree {
private:
    NodeT* root;

    NodeT* _insert(NodeT* node, int value);
    void clear(NodeT* node);
    void printTree(NodeT* node, int depth);
    void inOrderToVector(NodeT* node, vector<int>& result) const;

public:
    FullBinaryTree();
    ~FullBinaryTree();
    void insert(int value);
    void print();
    void printZigZag();
    bool isEmpty() const;
    
    // Для сериализации
    vector<int> toVector() const;
    void fromVector(const vector<int>& elements);
};

#endif