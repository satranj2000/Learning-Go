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
