package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return nil
	}
	var q []*TreeNode
	var traversed = make(map[string]interface{})
	var results []*TreeNode
	var resultsAsString = make(map[string]interface{})
	constMapVal := &struct{}{}

	if root.Left != nil {
		q = append(q, root.Left)
	}
	if root.Right != nil {
		q = append(q, root.Right)
	}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		encoded := encodeTree(node)

		_, isDup := traversed[encoded]
		if !isDup {
			traversed[encoded] = constMapVal
			continue
		}
		// isDup == true
		_, isDupInResults := resultsAsString[encoded]
		if !isDupInResults {
			results = append(results, node)
			resultsAsString[encoded] = constMapVal
		}
	}
	return results
}

const null = "null"

var mapNodeToEncodedString = make(map[*TreeNode]string)

func encodeTree(tree *TreeNode) string {
	if tree == nil {
		return null
	}
	if ret, ok := mapNodeToEncodedString[tree]; ok {
		return ret
	}
	ret := fmt.Sprintf("%d,%s,%s", tree.Val, encodeTree(tree.Left), encodeTree(tree.Right))
	mapNodeToEncodedString[tree] = ret
	return ret
}
