package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type avgData struct {
	sum         int
	nodesNumber int
}

func (r *avgData) Add(other avgData) {
	r.sum += other.sum
	r.nodesNumber += other.nodesNumber
}

func recursiveAvg(node *TreeNode) (avgData, int) {
	d := avgData{node.Val, 1}
	matchingNumber := 0

	if node.Left != nil {
		stData, stMatchNo := recursiveAvg(node.Left)
		d.Add(stData)
		matchingNumber += stMatchNo
	}
	if node.Right != nil {
		stData, stMatchNo := recursiveAvg(node.Right)
		d.Add(stData)
		matchingNumber += stMatchNo
	}

	if d.sum/d.nodesNumber == node.Val {
		matchingNumber++
	}

	return d, matchingNumber
}

func averageOfSubtree(root *TreeNode) int {
	_, result := recursiveAvg(root)
	return result
}
