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

	var result []int
	queue := []*Node{root}

	for len(queue) > 0 {
		processingNode := queue[0]
		queue = queue[1:]

		result = append(result, processingNode.Value)

		if processingNode.Left != nil {
			queue = append(queue, processingNode.Left)
		}
		if processingNode.Right != nil {
			queue = append(queue, processingNode.Right)
		}
	}

	return result
}
