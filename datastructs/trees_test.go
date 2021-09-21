package datastructs

import "testing"

func TestPreorderTraversal(t *testing.T) {

	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Right = &TreeNode{30, nil, nil}
	root.Left.Left = &TreeNode{40, nil, nil}
	root.Left.Right = &TreeNode{50, nil, nil}
	root.Right.Right = &TreeNode{60, nil, nil}
	root.Left.Left.Left = &TreeNode{70, nil, nil}

	output := []int{10, 20, 40, 70, 50, 30, 60}

	res := preorderTraversal(root)

	for i := 0; i < len(output); i++ {
		if output[i] != res[i] {
			t.Errorf("Expected %v, got %v", output[i], res[i])
		}
	}

}

func TestPreorderTraversal2(t *testing.T) {

	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	output := []int{10, 20}

	res := preorderTraversal(root)

	for i := 0; i < len(output); i++ {
		if output[i] != res[i] {
			t.Errorf("Expected %v, got %v", output[i], res[i])
		}
	}

}

func TestPreorderTraversal3(t *testing.T) {

	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Left.Left = &TreeNode{40, nil, nil}
	root.Left.Left.Left = &TreeNode{70, nil, nil}

	output := []int{10, 20, 40, 70}

	res := preorderTraversal(root)

	for i := 0; i < len(output); i++ {
		if output[i] != res[i] {
			t.Errorf("Expected %v, got %v", output[i], res[i])
		}
	}

}

func TestPreorderTraversal4(t *testing.T) {

	res := preorderTraversal(nil)

	if res != nil {
		t.Errorf("Expected nil values")
	}

}

func TestInorderTraversal4(t *testing.T) {

	res := inorderTraversal(nil)

	if res != nil {
		t.Errorf("Expected nil values")
	}

}

func TestInorderTraversal2(t *testing.T) {

	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Left.Left = &TreeNode{40, nil, nil}
	root.Left.Left.Left = &TreeNode{70, nil, nil}

	output := []int{70, 40, 20, 10}

	res := inorderTraversal(root)

	for i := 0; i < len(output); i++ {
		if output[i] != res[i] {
			t.Errorf("Expected %v, got %v", output[i], res[i])
		}
	}

}

func TestInorderTraversal3(t *testing.T) {

	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	output := []int{20, 10}

	res := inorderTraversal(root)

	for i := 0; i < len(output); i++ {
		if output[i] != res[i] {
			t.Errorf("Expected %v, got %v", output[i], res[i])
		}
	}

}

func TestInorderTraversal5(t *testing.T) {

	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Right = &TreeNode{30, nil, nil}
	root.Left.Left = &TreeNode{40, nil, nil}
	root.Left.Right = &TreeNode{50, nil, nil}
	root.Right.Right = &TreeNode{60, nil, nil}
	root.Left.Left.Left = &TreeNode{70, nil, nil}

	output := []int{70, 40, 20, 50, 10, 30, 60}

	res := inorderTraversal(root)

	for i := 0; i < len(output); i++ {
		if output[i] != res[i] {
			t.Errorf("Expected %v, got %v", output[i], res[i])
		}
	}

}

func TestMaxDepth(t *testing.T) {

	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Right = &TreeNode{30, nil, nil}
	root.Left.Left = &TreeNode{40, nil, nil}
	root.Left.Right = &TreeNode{50, nil, nil}
	root.Right.Right = &TreeNode{60, nil, nil}
	root.Left.Left.Left = &TreeNode{70, nil, nil}

	res := maxDepth(root)

	if res != 4 {
		t.Errorf("Expected depth is 3, got %v", res)
	}
}

func TestMaxDepth2(t *testing.T) {

	root := &TreeNode{10, nil, nil}

	res := maxDepth(root)

	if res != 1 {
		t.Errorf("Expected depth is 1, got %v", res)
	}
}

func TestMaxDepth3(t *testing.T) {

	res := maxDepth(nil)

	if res != 0 {
		t.Errorf("Expected depth is 1, got %v", res)
	}
}

func TestMaxDepth4(t *testing.T) {

	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Left.Left = &TreeNode{40, nil, nil}
	root.Left.Left.Left = &TreeNode{70, nil, nil}

	res := maxDepth(root)

	if res != 4 {
		t.Errorf("Expected depth is 3, got %v", res)
	}
}

func TestMaxDepth5(t *testing.T) {

	root := &TreeNode{10, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Right.Right = &TreeNode{40, nil, nil}
	root.Right.Right.Right = &TreeNode{70, nil, nil}

	res := maxDepth(root)

	if res != 4 {
		t.Errorf("Expected depth is 3, got %v", res)
	}
}

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{10, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Right.Right = &TreeNode{40, nil, nil}
	root.Right.Right.Right = &TreeNode{70, nil, nil}

	res := isSymmetric(root)

	if res {
		t.Errorf("Expected it to be false. But, got true")
	}
}

func TestIsSymmetric2(t *testing.T) {
	root := &TreeNode{10, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Left = &TreeNode{20, nil, nil}

	res := isSymmetric(root)

	if !res {
		t.Errorf("Expected it to be true. But, got false")
	}
}

func TestIsSymmetric3(t *testing.T) {
	root := &TreeNode{10, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Right.Right = &TreeNode{30, nil, nil}
	root.Left.Left = &TreeNode{30, nil, nil}
	root.Left.Right = &TreeNode{40, nil, nil}
	root.Right.Left = &TreeNode{40, nil, nil}

	res := isSymmetric(root)

	if !res {
		t.Errorf("Expected it to be true. But, got false")
	}
}

func TestIsSymmetric4(t *testing.T) {
	root := &TreeNode{10, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Right.Right = &TreeNode{30, nil, nil}
	root.Left.Left = &TreeNode{30, nil, nil}
	root.Left.Right = &TreeNode{41, nil, nil}
	root.Right.Left = &TreeNode{40, nil, nil}

	res := isSymmetric(root)

	if res {
		t.Errorf("Expected it to be False. But, got true")
	}
}

func TestInvertNode(t *testing.T) {
	root := &TreeNode{10, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Left = &TreeNode{3, nil, nil}

	resNode := invertTree(root)

	expectedOutputs := []int{10, 2, 3}
	resNodeVals := preorderTraversal(resNode)

	for i, _ := range expectedOutputs {
		if expectedOutputs[i] != resNodeVals[i] {
			t.Errorf("Expected %v, got %v", expectedOutputs[i], resNodeVals[i])
		}
	}
}

func TestInvertNode2(t *testing.T) {
	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{31, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Right.Right = &TreeNode{21, nil, nil}

	resNode := invertTree(root)

	expectedOutputs := []int{10, 2, 21, 3, 31}
	resNodeVals := preorderTraversal(resNode)

	for i, _ := range expectedOutputs {
		if expectedOutputs[i] != resNodeVals[i] {
			t.Errorf("Expected %v, got %v", expectedOutputs[i], resNodeVals[i])
		}
	}
}

func TestHasPathSum(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{9, nil, nil}

	res := hasPathSum(root, 8)
	if !res {
		t.Errorf("Expected true, got false for sum as 8 ")
	}

	res = hasPathSum(root, 9)
	if res {
		t.Errorf("Expected false, got true for sum as 9 ")
	}

	res = hasPathSum(root, 14)
	if !res {
		t.Errorf("Expected true, got false for sum as 14 ")
	}

	res = hasPathSum(root, 5)
	if res {
		t.Errorf("Expected false, got true for sum as 5 ")
	}

}

func TestHasPathSum2(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{9, nil, nil}

	res := hasPathSum(root, 17)
	if !res {
		t.Errorf("Expected true, got false for sum as 17 ")
	}

	res = hasPathSum(root, 8)
	if res {
		t.Errorf("Expected false, got true for sum as 8 ")
	}

	res = hasPathSum(root, 12)
	if res {
		t.Errorf("Expected false, got true for sum as 12 ")
	}

}

func TestIsBalanced(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{9, nil, nil}

	res := isBalanced(root)
	if !res {
		t.Error("Expected true, got false ")
	}
}

func TestIsBalanced2(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{9, nil, nil}

	res := isBalanced(root)
	if res {
		t.Error("Expected false, got true ")
	}
}

func TestIsBalanced3(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Right.Left = &TreeNode{9, nil, nil}

	res := isBalanced(root)
	if res {
		t.Error("Expected false, got true ")
	}
}

func TestIsBalanced4(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Right.Right = &TreeNode{9, nil, nil}

	res := isBalanced(root)
	if res {
		t.Error("Expected false, got true ")
	}
}

func TestSearchBST(t *testing.T) {

	root := &TreeNode{5, nil, nil}
	root.Right = &TreeNode{8, nil, nil}
	root.Right.Right = &TreeNode{9, nil, nil}

	res := searchBST(root, 8)
	if res.Val != 8 {
		t.Errorf("wrong node returned. ")
	}

	res = searchBST(root, 5)
	if res.Val != 5 {
		t.Errorf("wrong node returned. ")
	}

	res = searchBST(root, 15)
	if res != nil {
		t.Errorf("wrong node returned. ")
	}
}

func TestInsertBST(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Right = &TreeNode{8, nil, nil}
	root.Right.Right = &TreeNode{9, nil, nil}

	result := insertIntoBST(root, 3)

	output := []int{3, 5, 8, 9}

	reslist := inorderTraversal(result)

	for i, _ := range output {
		if output[i] != reslist[i] {
			t.Errorf("Expected %v, got %v", output[i], reslist[i])
		}
	}

}

func TestInsertBST2(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Right = &TreeNode{8, nil, nil}
	root.Right.Right = &TreeNode{9, nil, nil}

	result := insertIntoBST(root, 18)

	output := []int{5, 8, 9, 18}

	reslist := inorderTraversal(result)

	for i, _ := range output {
		if output[i] != reslist[i] {
			t.Errorf("Expected %v, got %v", output[i], reslist[i])
		}
	}

}

func TestInsertBST3(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{8, nil, nil}

	result := insertIntoBST(root, 4)

	output := []int{3, 4, 5, 8}

	reslist := inorderTraversal(result)

	for i, _ := range output {
		if output[i] != reslist[i] {
			t.Errorf("Expected %v, got %v", output[i], reslist[i])
		}
	}

}

func TestInsertBST4(t *testing.T) {

	root := insertIntoBST(nil, 10)

	out := inorderTraversal(root)

	if out[0] != 10 && len(out) != 1 {
		t.Errorf("Expected a single root node")
	}

}

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{8, nil, nil}

	res := isValidBST(root)

	if !res {
		t.Error("Expected it to be a valid BST. But, got a false result")
	}

}

func TestIsValidBST2(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{13, nil, nil}
	root.Right = &TreeNode{8, nil, nil}

	res := isValidBST(root)

	if res {
		t.Error("Expected it to be a invalid BST. But, got a valid result")
	}

}

func TestIsValidBST3(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{-3, nil, nil}
	root.Left.Left.Left = &TreeNode{-5, nil, nil}

	res := isValidBST(root)

	if !res {
		t.Error("Expected it to be a valid BST. But, got a invalid result")
	}

}

func TestIsValidBST4(t *testing.T) {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{4, nil, nil}
	root.Right = &TreeNode{6, nil, nil}
	root.Right.Left = &TreeNode{3, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}

	res := isValidBST(root)

	if res {
		t.Error("Expected it to be a invalid BST. But, got a valid result")
	}

}

func TestIsValidBST6(t *testing.T) {
	root := &TreeNode{2, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{2, nil, nil}

	res := isValidBST(root)

	if res {
		t.Error("Expected it to be a invalid BST. But, got a valid result")
	}

}

func TestFindTarget(t *testing.T) {
	root := &TreeNode{2, nil, nil}
	root.Left = &TreeNode{1, nil, nil}
	root.Right = &TreeNode{4, nil, nil}

	res := findTarget(root, 3)
	if !res {
		t.Error("Expected to find the target 3, but did not get it")
	}

	res = findTarget(root, 6)
	if !res {
		t.Error("Expected to find the target 6, but did not get it")
	}

	res = findTarget(root, 5)
	if !res {
		t.Error("Expected to find the target 5, but did not get it")
	}

	res = findTarget(root, 7)
	if res {
		t.Error("not expected to find target 7. but, got it. ")
	}
}

func TestFindMode(t *testing.T) {
	root := &TreeNode{1, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{2, nil, nil}

	res := findMode(root)

	if res[0] != 2 || len(res) != 1 {
		t.Errorf("Expected the mode to be 2; but %v", res[0])
	}
}

func TestFindMode2(t *testing.T) {
	root := &TreeNode{0, nil, nil}

	res := findMode(root)

	if res[0] != 0 || len(res) != 1 {
		t.Errorf("Expected the mode to be 0; but %v", res[0])
	}
}

func TestFindMode3(t *testing.T) {
	root := &TreeNode{2, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Left.Left = &TreeNode{1, nil, nil}
	root.Left.Left.Left = &TreeNode{1, nil, nil}

	res := findMode(root)

	output := []int{2, 1}

	for i := 0; i < 2; i++ {
		if output[i] != res[i] {
			t.Errorf("not expected mode values were received. Expected %v. but got %v", output, res)
		}
	}
}
