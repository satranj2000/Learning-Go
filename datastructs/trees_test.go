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
