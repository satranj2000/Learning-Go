package datastructs

//A tree node with a pointer to right node
// reference image :- https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
type NodeL struct {
	Val   int
	Left  *NodeL
	Right *NodeL
	Next  *NodeL
}

//https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
// this is perfect binary tree.
func connect(root *NodeL) *NodeL {
	if root == nil {
		return nil
	}
	root.Next = nil
	connectTrees(root, nil)
	return root
}

func connectTrees(root *NodeL, right *NodeL) {
	if root == nil {
		return
	}
	if root != nil {
		root.Next = right
	}
	connectTrees(root.Left, root.Right)
	if root.Next != nil {
		connectTrees(root.Right, root.Next.Left)
	} else {
		connectTrees(root.Right, nil)
	}

}

//https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
// not yet completed.
func connectnotperfectbinary(root *NodeL) *NodeL {
	if root == nil {
		return nil
	}
	root.Next = nil
	connectTreesnotperfect(root, nil)
	return root
}

func connectTreesnotperfect(root *NodeL, right *NodeL) {
	if root == nil {
		return
	}
	if root != nil {
		root.Next = right
	}

	if root.Right != nil {
		if root.Next != nil {
			if root.Next.Left != nil {
				connectTreesnotperfect(root.Right, root.Next.Left)
			} else {
				connectTreesnotperfect(root.Right, root.Next.Right)
			}
		} else {
			connectTreesnotperfect(root.Right, nil)
		}
	}

	if root.Left != nil {
		if root.Right != nil {
			connectTreesnotperfect(root.Left, root.Right)
		} else {
			Next := root.Next
			for Next != nil && Next.Left == nil && Next.Right == nil {
				Next = Next.Next
			}
			if Next != nil {
				if Next.Left != nil {
					connectTreesnotperfect(root.Left, Next.Left)
				} else {
					connectTreesnotperfect(root.Left, Next.Right)
				}
			} else {
				connectTreesnotperfect(root.Left, nil)
			}
		}
	}

}
