package datastructs

import "testing"

func TestOddEvenList(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	res := OddEvenList(&head)
	res.ListAll()
}

func TestOddEvenList2(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}
	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	res := OddEvenList(&head)
	outarr := res.AsArray()
	output := []int{1, 3, 5, 2, 4}
	for i := 0; i < len(output); i++ {
		if outarr[i] != output[i] {
			t.Errorf("Expected %v, got %v", outarr[i], output[i])
		}
	}
}

func TestOddEvenList3(t *testing.T) {
	head := ListNode{Val: 2, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 5, Next: nil}
	node5 := ListNode{Val: 6, Next: nil}
	node6 := ListNode{Val: 4, Next: nil}
	node7 := ListNode{Val: 7, Next: nil}
	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7

	res := OddEvenList(&head)

	outarr := res.AsArray()
	output := []int{2, 3, 6, 7, 1, 5, 4}
	for i := 0; i < len(output); i++ {
		if outarr[i] != output[i] {
			t.Errorf("Expected %v, got %v", outarr[i], output[i])
		}
	}
}

func TestOddEvenList4(t *testing.T) {
	head := ListNode{Val: 2, Next: nil}
	res := OddEvenList(&head)

	outarr := res.AsArray()
	output := []int{2}
	for i := 0; i < len(output); i++ {
		if outarr[i] != output[i] {
			t.Errorf("Expected %v, got %v", outarr[i], output[i])
		}
	}
}

func TestOddEvenList5(t *testing.T) {
	var head *ListNode = nil
	res := OddEvenList(head)

	outarr := res.AsArray()
	if len(outarr) > 0 {
		t.Errorf("Expected the length of array to be zero. Got more than that. ")
	}

}

func TestBrowserHistory(t *testing.T) {
	obj := Constructor("leetcode.com")
	obj.Visit("google.com")
	obj.Visit("facebook.com")
	obj.Visit("youtube.com")

	param_2 := obj.Back(1)
	if param_2 != "facebook.com" {
		t.Errorf("Expected facebook.com, got %v", param_2)
	}
	param_2 = obj.Back(1)
	if param_2 != "google.com" {
		t.Errorf("Expected google.com, got %v", param_2)
	}
}

func TestReverseList(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}

	outvalues := []int{3, 2, 1}

	head.Next = &node2
	node2.Next = &node3

	revhead := ReverseList(&head)
	out := revhead.AsArray()

	for i := 0; i < len(outvalues); i++ {
		if out[i] != outvalues[i] {
			t.Errorf("Expected %v, got %v", outvalues[i], out[i])
		}
	}
}

func TestReverseList2(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	outvalues := []int{1}

	revhead := ReverseList(&head)
	out := revhead.AsArray()

	for i := 0; i < len(outvalues); i++ {
		if out[i] != outvalues[i] {
			t.Errorf("Expected %v, got %v", outvalues[i], out[i])
		}
	}
}

func TestReverseList3(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}

	outvalues := []int{2, 1}

	head.Next = &node2

	revhead := ReverseList(&head)
	out := revhead.AsArray()

	for i := 0; i < len(outvalues); i++ {
		if out[i] != outvalues[i] {
			t.Errorf("Expected %v, got %v", outvalues[i], out[i])
		}
	}
}

func TestReverseBetween(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}

	outvalues := []int{1, 4, 3, 2, 5}
	left := 2
	right := 4

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	retlist := ReverseBetween(&head, left, right)

	retvalues := retlist.AsArray()

	for i := 0; i < len(outvalues); i++ {
		if retvalues[i] != outvalues[i] {
			t.Errorf("Expected %v, got %v", outvalues[i], retvalues[i])
		}
	}

}

func TestReverseBetween2(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}

	outvalues := []int{4, 3, 2, 1, 5}
	left := 1 // left value is same as origin
	right := 4

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	retlist := ReverseBetween(&head, left, right)

	retvalues := retlist.AsArray()

	for i := 0; i < len(outvalues); i++ {
		if retvalues[i] != outvalues[i] {
			t.Errorf("Expected %v, got %v", outvalues[i], retvalues[i])
		}
	}
}

func TestReverseBetween3(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}

	outvalues := []int{5, 4, 3, 2, 1}
	left := 1  //left is same as origin
	right := 5 // right is same as tail

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	retlist := ReverseBetween(&head, left, right)

	retvalues := retlist.AsArray()

	for i := 0; i < len(outvalues); i++ {
		if retvalues[i] != outvalues[i] {
			t.Errorf("Expected %v, got %v", outvalues[i], retvalues[i])
		}
	}
}

func TestReverseBetween4(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}

	outvalues := []int{1, 2, 3, 4, 5}
	left := 3 //left and right as same value
	right := 3

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	retlist := ReverseBetween(&head, left, right)

	retvalues := retlist.AsArray()

	for i := 0; i < len(outvalues); i++ {
		if retvalues[i] != outvalues[i] {
			t.Errorf("Expected %v, got %v", outvalues[i], retvalues[i])
		}
	}
}

func TestRotateRight(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	outputs := []int{4, 5, 1, 2, 3}
	rotate := 2

	result := RotateRight(&head, rotate)
	resultArr := result.AsArray()

	for j := 0; j < len(outputs); j++ {
		if outputs[j] != resultArr[j] {
			t.Errorf("Expected %v, got %v ", resultArr[j], outputs[j])
		}
	}

}

func TestRotateRight2(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	outputs := []int{2, 3, 4, 5, 1}
	rotate := 4

	result := RotateRight(&head, rotate)
	resultArr := result.AsArray()

	for j := 0; j < len(outputs); j++ {
		if outputs[j] != resultArr[j] {
			t.Errorf("Expected %v, got %v ", resultArr[j], outputs[j])
		}
	}

}

func TestRotateRight3(t *testing.T) {
	head := ListNode{Val: 0, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}

	head.Next = &node2

	outputs := []int{0, 1}
	rotate := 1000

	result := RotateRight(&head, rotate)
	resultArr := result.AsArray()

	for j := 0; j < len(outputs); j++ {
		if outputs[j] != resultArr[j] {
			t.Errorf("Expected %v, got %v ", resultArr[j], outputs[j])
		}
	}

}

func TestRotateRight4(t *testing.T) {
	head := ListNode{Val: 0, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 2, Next: nil}

	head.Next = &node2
	node2.Next = &node3

	outputs := []int{2, 0, 1}
	rotate := 4

	result := RotateRight(&head, rotate)
	resultArr := result.AsArray()

	for j := 0; j < len(outputs); j++ {
		if outputs[j] != resultArr[j] {
			t.Errorf("Expected %v, got %v ", outputs[j], resultArr[j])
		}
	}

}

func TestMergeInBetween(t *testing.T) {
	//list 1
	head := ListNode{Val: 0, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 2, Next: nil}
	node4 := ListNode{Val: 3, Next: nil}
	node5 := ListNode{Val: 4, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	//list 2
	ahead := ListNode{Val: 10, Next: nil}
	anode2 := ListNode{Val: 11, Next: nil}
	anode3 := ListNode{Val: 21, Next: nil}
	anode4 := ListNode{Val: 31, Next: nil}
	anode5 := ListNode{Val: 41, Next: nil}

	ahead.Next = &anode2
	anode2.Next = &anode3
	anode3.Next = &anode4
	anode4.Next = &anode5

	res := MergeInBetween(&head, 2, 3, &ahead)

	outvalues := res.AsArray()

	expectedOuts := []int{0, 1, 10, 11, 21, 31, 41, 4}

	for i := 0; i < len(expectedOuts); i++ {
		if expectedOuts[i] != outvalues[i] {
			t.Errorf("Expected %v, got %v", expectedOuts[i], outvalues[i])
		}
	}
}

func TestMergeInBetween2(t *testing.T) {
	//list 1
	head := ListNode{Val: 0, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 2, Next: nil}
	node4 := ListNode{Val: 3, Next: nil}
	node5 := ListNode{Val: 4, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	//list 2
	ahead := ListNode{Val: 10, Next: nil}
	anode2 := ListNode{Val: 11, Next: nil}
	anode3 := ListNode{Val: 21, Next: nil}
	anode4 := ListNode{Val: 31, Next: nil}
	anode5 := ListNode{Val: 41, Next: nil}

	ahead.Next = &anode2
	anode2.Next = &anode3
	anode3.Next = &anode4
	anode4.Next = &anode5

	res := MergeInBetween(&head, 3, 4, &ahead)

	outvalues := res.AsArray()

	expectedOuts := []int{0, 1, 2, 10, 11, 21, 31, 41}

	for i := 0; i < len(expectedOuts); i++ {
		if expectedOuts[i] != outvalues[i] {
			t.Errorf("Expected %v, got %v", expectedOuts[i], outvalues[i])
		}
	}
}
