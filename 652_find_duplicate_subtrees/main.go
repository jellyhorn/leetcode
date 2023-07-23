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
	var traversed []string
	var results []*TreeNode
	var resultsAsString []string

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

		var isDup bool
		for _, traversedNode := range traversed {
			if encoded == traversedNode {
				isDup = true
				break
			}
		}
		if !isDup {
			traversed = append(traversed, encoded)
			continue
		}
		// isDup == true
		var isDupInResults bool
		for _, resultNode := range resultsAsString {
			if encoded == resultNode {
				isDupInResults = true
				break
			}
		}
		if !isDupInResults {
			results = append(results, node)
			resultsAsString = append(resultsAsString, encoded)
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
