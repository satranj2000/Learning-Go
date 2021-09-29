package datastructs

import (
	"math"
	"sort"
)

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

//https://leetcode.com/problems/path-sum/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)

}

//https://leetcode.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lsize := maxDepth(root.Left)
	rsize := maxDepth(root.Right)

	// check the depth is not more than 1 and call the isBalanced again for left and right
	return math.Abs(float64(lsize)-float64(rsize)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

// not yet fully working.. need to continue working on this.
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lsize := maxDepth(root.Left)
	rsize := maxDepth(root.Right)

	return int(math.Min(float64(lsize), float64(rsize)))

}

func searchBST(root *TreeNode, val int) *TreeNode {

	for root != nil {
		if val == root.Val {
			return root
		}
		if val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return nil
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		newroot := &TreeNode{Val: val, Left: nil, Right: nil}
		return newroot
	}

	if val > root.Val {
		if root.Right != nil {
			insertIntoBST(root.Right, val)
		} else {
			root.Right = &TreeNode{Val: val, Left: nil, Right: nil}
		}

	} else {
		if root.Left != nil {
			insertIntoBST(root.Left, val)
		} else {
			root.Left = &TreeNode{Val: val, Left: nil, Right: nil}
		}
	}
	return root
}

func isValidBST(root *TreeNode) bool {
	return isValidBinaryTree(root, math.MinInt64, math.MaxInt64)
}

func isValidBinaryTree(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return isValidBinaryTree(root.Left, min, root.Val) && isValidBinaryTree(root.Right, root.Val, max)
}

// find two numbers in BST that are equivalent to the sum
func findTarget(root *TreeNode, k int) bool {

	numlist := inorderTraversal(root)

	for i := 0; i < len(numlist); i++ {
		for j := i; j < len(numlist); j++ {
			if numlist[i]+numlist[j] == k {
				return true
			}
		}
	}
	return false
}

func findMode(root *TreeNode) []int {

	modes := make(map[int]int)
	findNodeBST(root, modes)
	// fill the map with node value and their count (key - node value; value - count)

	// sort them.
	valueslist := make([]int, len(modes))
	i := 0
	for _, v := range modes {
		valueslist[i] = v
		i++
	}
	sort.Ints(valueslist)

	// take the top
	topnodes := make([]int, 0, len(modes))

	for k, v := range modes {
		if v == valueslist[len(valueslist)-1] {
			topnodes = append(topnodes, k)
		}
	}
	return topnodes
}

func findNodeBST(root *TreeNode, modes map[int]int) {
	if root == nil {
		return
	}
	if root != nil {
		if _, ok := modes[root.Val]; ok {
			modes[root.Val] += 1
		} else {
			modes[root.Val] = 1
		}
	}

	findNodeBST(root.Left, modes)
	findNodeBST(root.Right, modes)
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftHeight, rightHeight int

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil {
		leftHeight = treeHeight(root.Left) + 1
	}

	if root.Right != nil {
		rightHeight = treeHeight(root.Right) + 1
	}

	if leftHeight > rightHeight {
		return leftHeight
	} else {
		return rightHeight
	}
}

func inorderTraversalNoRecursion(root *TreeNode) {

}

//https://leetcode.com/problems/merge-two-binary-trees/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	var root1val, root2val int

	if root1 == nil {
		root1val = 0
	} else {
		root1val = root1.Val
	}

	if root2 == nil {
		root2val = 0
	} else {
		root2val = root2.Val
	}
	mergedVals := root1val + root2val

	newRoot := &TreeNode{Val: mergedVals}

	if root1 == nil || root2 == nil {
		if root1 == nil {
			newRoot.Left = mergeTrees(nil, root2.Left)
			newRoot.Right = mergeTrees(nil, root2.Right)
		} else if root2 == nil {
			newRoot.Left = mergeTrees(root1.Left, nil)
			newRoot.Right = mergeTrees(root1.Right, nil)
		}
	} else {
		newRoot.Left = mergeTrees(root1.Left, root2.Left)
		newRoot.Right = mergeTrees(root1.Right, root2.Right)
	}
	return newRoot
}
