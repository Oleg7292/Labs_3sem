package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type FullBinaryTree struct {
	root *TreeNode
}

func NewFullBinaryTree() *FullBinaryTree {
	return &FullBinaryTree{}
}

func (b *FullBinaryTree) Insert(value int) {
	b.root = b.insert(b.root, value)
}

func (b *FullBinaryTree) insert(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{data: value}
	}
	if value < node.data {
		node.left = b.insert(node.left, value)
	} else {
		node.right = b.insert(node.right, value)
	}
	return node
}

func (b *FullBinaryTree) Print() {
	b.printTree(b.root, 0)
}

func (b *FullBinaryTree) printTree(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	b.printTree(node.right, depth+1)
	fmt.Printf("%s%d\n", strings.Repeat("    ", depth), node.data)
	b.printTree(node.left, depth+1)
}

func (b *FullBinaryTree) PrintZigZag() {
	if b.root == nil {
		return
	}
	currentLevel := []*TreeNode{b.root}
	nextLevel := []*TreeNode{}
	leftToRight := true

	for len(currentLevel) > 0 {
		node := currentLevel[len(currentLevel)-1]
		currentLevel = currentLevel[:len(currentLevel)-1]

		if node != nil {
			fmt.Printf("%d ", node.data)
			if leftToRight {
				if node.left != nil {
					nextLevel = append(nextLevel, node.left)
				}
				if node.right != nil {
					nextLevel = append(nextLevel, node.right)
				}
			} else {
				if node.right != nil {
					nextLevel = append(nextLevel, node.right)
				}
				if node.left != nil {
					nextLevel = append(nextLevel, node.left)
				}
			}
		}

		if len(currentLevel) == 0 {
			fmt.Println()
			currentLevel, nextLevel = nextLevel, currentLevel
			leftToRight = !leftToRight
		}
	}
}

func (b *FullBinaryTree) IsEmpty() bool {
	return b.root == nil
}

func (b *FullBinaryTree) ToVector() []int {
	result := []int{}
	b.inOrderToVector(b.root, &result)
	return result
}

func (b *FullBinaryTree) inOrderToVector(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	b.inOrderToVector(node.left, result)
	*result = append(*result, node.data)
	b.inOrderToVector(node.right, result)
}

func (b *FullBinaryTree) FromVector(elements []int) {
	b.root = nil
	for _, value := range elements {
		b.Insert(value)
	}
}
