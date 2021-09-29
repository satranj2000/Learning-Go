package datastructs

import "testing"

func TestConnectTrees(t *testing.T) {

	root := &NodeL{1, nil, nil, nil}
	root.Left = &NodeL{3, nil, nil, nil}
	root.Right = &NodeL{2, nil, nil, nil}

	out := connect(root)

	if out.Left.Next.Val != 2 {
		t.Errorf("Next val at first level not set properly. ")
	}

}

func TestConnectTrees2(t *testing.T) {

	root := &NodeL{1, nil, nil, nil}
	root.Left = &NodeL{2, nil, nil, nil}
	root.Right = &NodeL{3, nil, nil, nil}
	root.Left.Left = &NodeL{4, nil, nil, nil}
	root.Left.Right = &NodeL{5, nil, nil, nil}
	root.Right.Left = &NodeL{6, nil, nil, nil}
	root.Right.Right = &NodeL{7, nil, nil, nil}

	out := connect(root)

	if out.Left.Next.Val != 3 {
		t.Errorf("Next val at first level not set properly. ")
	}
	if out.Left.Right.Next.Val != 6 {
		t.Errorf("Next val at first level not set properly. ")
	}
}

func TestConnectTreesnotPerfect(t *testing.T) {

	root := &NodeL{1, nil, nil, nil}
	root.Left = &NodeL{2, nil, nil, nil}
	root.Right = &NodeL{3, nil, nil, nil}
	root.Left.Left = &NodeL{4, nil, nil, nil}
	root.Left.Right = &NodeL{5, nil, nil, nil}
	root.Right.Right = &NodeL{7, nil, nil, nil}

	out := connectnotperfectbinary(root)

	if out.Left.Next.Val != 3 {
		t.Errorf("Next val at first level not set properly. ")
	}

	if out.Left.Right.Next.Val != 7 {
		t.Errorf("Next val at 2nd level not set properly. ")
	}
}

func TestConnectTreesnotPerfect2(t *testing.T) {

	root := &NodeL{1, nil, nil, nil}
	root.Left = &NodeL{2, nil, nil, nil}

	out := connectnotperfectbinary(root)

	if out.Left.Next != nil {
		t.Errorf("Next val at first level not set properly. ")
	}
}

func TestConnectTreesnotPerfect3(t *testing.T) {

	root := &NodeL{1, nil, nil, nil}
	root.Left = &NodeL{2, nil, nil, nil}
	root.Right = &NodeL{3, nil, nil, nil}
	root.Left.Left = &NodeL{4, nil, nil, nil}
	root.Left.Right = &NodeL{5, nil, nil, nil}
	root.Right.Right = &NodeL{6, nil, nil, nil}
	root.Left.Left.Left = &NodeL{7, nil, nil, nil}
	root.Right.Right.Right = &NodeL{8, nil, nil, nil}

	out := connectnotperfectbinary(root)

	if out.Left.Next.Val != 3 {
		t.Errorf("Next val at first level not set properly. ")
	}

	if out.Left.Right.Next.Val != 6 {
		t.Errorf("Next val at 2nd level not set properly. ")
	}

	if out.Left.Left.Left.Next.Val != 8 {
		t.Errorf("Next val at 3rd level not set properly. ")
	}
}
