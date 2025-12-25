package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func TraverseBreadthFirst(root *Node) []int {
	if root == nil {
		return nil
	}
}
