package main

import (
	"fmt"
	"reflect"
)

func main() {
	root := NewNode(0)
	root.Left = NewNode(1)
	root.Right = NewNode(2)
	root.Left.Left = NewNode(3)
	root.Left.Right = NewNode(4)
	root.Right.Left = NewNode(5)
	root.Right.Right = NewNode(6)

	expectedBreadthFirstTraversalResult := []int{0, 1, 2, 3, 4, 5, 6}
	expectedDepthFirstTraversalResult := []int{3, 1, 4, 0, 5, 2, 6}

	if reflect.DeepEqual(expectedBreadthFirstTraversalResult, TraverseBreadthFirst(root)) {
		fmt.Println("Breadth-first tree traversal succeeded.")
	}

	if reflect.DeepEqual(expectedDepthFirstTraversalResult, TraverseDepthFirst(root)) {
		fmt.Println("Depth-first tree traversal succeeded.")
	}
}
