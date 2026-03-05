#include "binary_tree.h"

FullBinaryTree::FullBinaryTree() : root(nullptr) {}

FullBinaryTree::~FullBinaryTree() { 
    clear(root); 
}

NodeT* FullBinaryTree::_insert(NodeT* node, int value) {
    if (!node) return new NodeT(value);
    if (value < node->data) {
        node->left = _insert(node->left, value);
    } else {
        node->right = _insert(node->right, value);
    }
    return node;
}

void FullBinaryTree::clear(NodeT* node) {
    if (!node) return;
    clear(node->left);
    clear(node->right);
    delete node;
}

void FullBinaryTree::insert(int value) { 
    root = _insert(root, value); 
}

void FullBinaryTree::printTree(NodeT* node, int depth) {
    if (!node) return;
    printTree(node->right, depth + 1);
    cout << setw(4 * depth) << " " << node->data << endl;
    printTree(node->left, depth + 1);
}

void FullBinaryTree::print() { 
    printTree(root, 0); 
}

void FullBinaryTree::printZigZag() {
    if (!root) return;
    
    stack<NodeT*> currentLevel;
    stack<NodeT*> nextLevel;
    bool leftToRight = true;
    
    currentLevel.push(root);
    
    while (!currentLevel.empty()) {
        NodeT* node = currentLevel.top();
        currentLevel.pop();
        
        if (node) {
            cout << node->data << " ";
            
            if (leftToRight) {
                if (node->left) nextLevel.push(node->left);
                if (node->right) nextLevel.push(node->right);
            } else {
                if (node->right) nextLevel.push(node->right);
                if (node->left) nextLevel.push(node->left);
            }
        }
        
        if (currentLevel.empty()) {
            cout << endl;
            swap(currentLevel, nextLevel);
            leftToRight = !leftToRight;
        }
    }
}

bool FullBinaryTree::isEmpty() const {
    return root == nullptr;
}

void FullBinaryTree::inOrderToVector(NodeT* node, vector<int>& result) const {
    if (!node) return;
    inOrderToVector(node->left, result);
    result.push_back(node->data);
    inOrderToVector(node->right, result);
}

vector<int> FullBinaryTree::toVector() const {
    vector<int> result;
    inOrderToVector(root, result);
    return result;
}

void FullBinaryTree::fromVector(const vector<int>& elements) {
    clear(root);
    root = nullptr;
    for (int value : elements) {
        insert(value);
    }
}