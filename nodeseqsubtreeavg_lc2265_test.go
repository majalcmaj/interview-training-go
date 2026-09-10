package main

import (
	"testing"
)

func TestAverageOfSubtreeYieldsExpectedResults(t *testing.T) {
	for _, d := range []struct {
		nodes    []int
		expected int
	}{
		{[]int{4, 8, 5, 0, 1, -1, 6}, 5},
		{[]int{1}, 1},
		{[]int{1, 1, 1}, 2},
		{[]int{4, -1, 5, -1, -1, -1, 6}, 2}, // Verify numbers
	} {
		root := arrayToGraph(d.nodes)

		result := averageOfSubtree(root)

		if result != d.expected {
			t.Errorf("Expected avg %d for graph %v, got %d", d.expected, d.nodes, result)
		}
	}
}

func arrayToGraph(values []int) *TreeNode {
	nodes := make([]*TreeNode, len(values))
	for idx, val := range values {
		if val != -1 {
			nodes[idx] = &TreeNode{Val: val}
		}
	}
	for idx, node := range nodes {
		if node == nil {
			continue
		}
		if idx*2+1 < len(nodes) {
			node.Left = nodes[idx*2+1]
		} else {
			break
		}
		if idx*2+2 < len(nodes) {
			node.Right = nodes[idx*2+2]
		} else {
			break
		}
	}
	return nodes[0]
}
