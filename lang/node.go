package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Print(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("nil")
	}
	node.value = value
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	root.setValue(4)
	fmt.Println(root.value)
	fmt.Println(nodes)
	//var pNode *treeNode
	//pNode.setValue(9)
	root.traverse()

	nodeCount := 0
	root.traverseFunc(func(n *treeNode) {
		nodeCount++
	})
	fmt.Println("node count: ", nodeCount)
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	fmt.Println(node.value)
	node.right.traverse()
}

func (node *treeNode) traverseFunc(f func(*treeNode)) {
	if node == nil {
		return
	}
	node.left.traverseFunc(f)
	f(node)
	node.right.traverseFunc(f)
}
