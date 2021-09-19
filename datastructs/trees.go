package datastructs

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodeVal := []int{root.Val}

	if root.Left != nil {
		nodeVal = append(nodeVal, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		nodeVal = append(nodeVal, preorderTraversal(root.Right)...)
	}

	return nodeVal
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodeVal := []int{}

	if root.Left != nil {
		nodeVal = append(nodeVal, inorderTraversal(root.Left)...)
	}

	nodeVal = append(nodeVal, root.Val)

	if root.Right != nil {
		nodeVal = append(nodeVal, inorderTraversal(root.Right)...)
	}

	return nodeVal
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodeVal := []int{}

	if root.Left != nil {
		nodeVal = append(nodeVal, postorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		nodeVal = append(nodeVal, postorderTraversal(root.Right)...)
	}

	nodeVal = append(nodeVal, root.Val)

	return nodeVal
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lsize := maxDepth(root.Left)
	rsize := maxDepth(root.Right)

	return int(math.Max(float64(lsize), float64(rsize))) + 1
}

//https://leetcode.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricNodes(root.Left, root.Right)
}

func isSymmetricNodes(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricNodes(left.Left, right.Right) && isSymmetricNodes(left.Right, right.Left)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	lnode := invertTree(root.Left)
	rnode := invertTree(root.Right)

	root.Left = rnode
	root.Right = lnode

	return root
}
