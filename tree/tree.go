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
	root.right.left.setValue(4)

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}

	//fmt.Println(root.value)
	//fmt.Println(nodes)
	//var pNode *treeNode
	//pNode.setValue(9)
	root.traverse()

	nodeCount := 0
	root.traverseFunc(func(n *treeNode) {
		nodeCount++
	})
	fmt.Println("node count: ", nodeCount)

	c := root.traverseWithChannel()
	maxValue := 0
	for node := range c {
		fmt.Print(node.value, " ")
		if node.value >= maxValue {
			maxValue = node.value
		}
	}
	fmt.Println("maxValue: ", maxValue)
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	fmt.Print(node.value, " ")
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

func (node *treeNode) traverseWithChannel() chan *treeNode {
	out := make(chan *treeNode)
	go func() {
		node.traverseFunc(func(node *treeNode) {
			out <- node
		})
		close(out)
	}()
	return out
}
