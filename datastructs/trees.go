package datastructs

import (
	"errors"
	"log"
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

// min depth is same as max depth with a slight difference. Send the min back and not the max value.
// We send min in all cases expect if one of the sides is zero, then we send the other one
//https://leetcode.com/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftHeight, rightHeight int

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil {
		leftHeight = minDepth(root.Left) + 1
	}

	if root.Right != nil {
		rightHeight = minDepth(root.Right) + 1
	}

	// if the left side height is zero, send the right side.
	if leftHeight == 0 {
		return rightHeight
	}
	// if right side height is zero, send the left side.
	if rightHeight == 0 {
		return leftHeight
	}
	// if both heights are provided, send the min back.
	if leftHeight > rightHeight {
		return rightHeight
	} else {
		return leftHeight
	}
}

// level order traversal. go through the tree level by level and return array of those values.
func levelOrderValues(root *TreeNode) []int {
	q := QueueTrees{}
	levelVal := []int{}
	q.Enqueue(root)

	for !q.IsEmpty() {
		currTree, _ := q.Dequeue()
		levelVal = append(levelVal, currTree.Val)
		if currTree.Left != nil {
			q.Enqueue(currTree.Left)
		}
		if currTree.Right != nil {
			q.Enqueue(currTree.Right)
		}
	}
	return levelVal
}

// in this case return an slice of slice. Each slice contains the values in that level.
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := QueueTrees{}
	Levels := [][]int{}
	q.Enqueue(root)

	for !q.IsEmpty() {
		sz := q.Size()
		currLevel := []int{}
		for i := 0; i < sz; i++ {
			currTree, _ := q.Dequeue()
			currLevel = append(currLevel, currTree.Val)
			if currTree.Left != nil {
				q.Enqueue(currTree.Left)
			}
			if currTree.Right != nil {
				q.Enqueue(currTree.Right)
			}
		}
		if len(currLevel) > 0 {
			Levels = append(Levels, currLevel)
		}
	}
	return Levels
}

// level order but from leaf to root
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := QueueTrees{}
	Levels := [][]int{}
	q.Enqueue(root)

	for !q.IsEmpty() {
		sz := q.Size()
		currLevel := []int{}
		for i := 0; i < sz; i++ {
			currTree, _ := q.Dequeue()
			currLevel = append(currLevel, currTree.Val)
			if currTree.Left != nil {
				q.Enqueue(currTree.Left)
			}
			if currTree.Right != nil {
				q.Enqueue(currTree.Right)
			}
		}
		if len(currLevel) > 0 {
			Levels = append(Levels, currLevel)
		}
	}
	// create a new slice of slices and store all values in reverse
	revLevels := [][]int{}
	for i := len(Levels) - 1; i >= 0; i-- {
		revLevels = append(revLevels, Levels[i])
	}
	return revLevels
}

// return average at each level
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	q := QueueTrees{}
	AvgLvl := []float64{}
	q.Enqueue(root)

	for !q.IsEmpty() {
		sz := q.Size()
		currLevel := []int{}
		for i := 0; i < sz; i++ {
			currTree, _ := q.Dequeue()
			currLevel = append(currLevel, currTree.Val)
			if currTree.Left != nil {
				q.Enqueue(currTree.Left)
			}
			if currTree.Right != nil {
				q.Enqueue(currTree.Right)
			}
		}
		if len(currLevel) > 0 {
			sum := 0.0
			for i := 0; i < len(currLevel); i++ {
				sum += float64(currLevel[i])
			}
			AvgLvl = append(AvgLvl, sum/float64(len(currLevel)))
		}
	}
	return AvgLvl
}

type QueueTrees struct {
	items       []*TreeNode
	bInitialize bool
	maxsz       int
	currmax     int
	currmin     int
}

func (q *QueueTrees) Enqueue(d *TreeNode) error {
	if !q.bInitialize {
		q.bInitialize = true
		q.maxsz = 4096
		q.currmax = 0
		q.currmin = 0
		q.items = make([]*TreeNode, q.maxsz)
	}
	if q.currmax > q.maxsz {
		log.Println("Cannot allocate more than 4096 values in the queue")
		return errors.New("cannot allocate more than 4096 values in the queue")
	}
	q.items[q.currmax] = d
	q.currmax++
	return nil
}

func (q *QueueTrees) Dequeue() (*TreeNode, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	val := q.items[q.currmin]
	q.currmin++
	return val, nil
}

func (q *QueueTrees) Size() int {
	return q.currmax - q.currmin
}

func (q *QueueTrees) IsEmpty() bool {
	return (q.currmax <= q.currmin)
}
