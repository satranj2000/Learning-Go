package datastructs

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
