package datastructs

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	l1.Next = &node2
	node2.Next = &node3

	l2 := ListNode{Val: 1, Next: nil}
	node21 := ListNode{Val: 2, Next: nil}
	node31 := ListNode{Val: 3, Next: nil}
	l2.Next = &node21
	node21.Next = &node31

	res := addTwoNumbers(&l1, &l2)
	outputRes := res.AsArray()
	output := []int{2, 4, 6}

	if !CompareArrays(output, outputRes) {
		t.Errorf("Expected %v, got %v", output, outputRes)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := ListNode{Val: 0, Next: nil}

	l2 := ListNode{Val: 1, Next: nil}

	res := addTwoNumbers(&l1, &l2)
	outputRes := res.AsArray()
	output := []int{1}

	if !CompareArrays(output, outputRes) {
		t.Errorf("Expected %v, got %v", output, outputRes)
	}
}

func TestAddTwoNumbers3(t *testing.T) {
	l1 := ListNode{Val: 2, Next: nil}
	node2 := ListNode{Val: 4, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	l1.Next = &node2
	node2.Next = &node3

	l2 := ListNode{Val: 5, Next: nil}
	node21 := ListNode{Val: 6, Next: nil}
	node31 := ListNode{Val: 4, Next: nil}
	l2.Next = &node21
	node21.Next = &node31

	res := addTwoNumbers(&l1, &l2)
	outputRes := res.AsArray()
	output := []int{7, 0, 8}

	if !CompareArrays(output, outputRes) {
		t.Errorf("Expected %v, got %v", output, outputRes)
	}
}

func TestAddTwoNumbers4(t *testing.T) {
	l1 := ListNode{Val: 9, Next: nil}
	node2 := ListNode{Val: 9, Next: nil}
	node3 := ListNode{Val: 9, Next: nil}
	l1.Next = &node2
	node2.Next = &node3

	l2 := ListNode{Val: 9, Next: nil}
	node21 := ListNode{Val: 9, Next: nil}
	node31 := ListNode{Val: 9, Next: nil}
	l2.Next = &node21
	node21.Next = &node31

	res := addTwoNumbers(&l1, &l2)
	outputRes := res.AsArray()
	output := []int{8, 9, 9, 1}

	if !CompareArrays(output, outputRes) {
		t.Errorf("Expected %v, got %v", output, outputRes)
	}
}

func TestAddTwoNumbers5(t *testing.T) {
	l1 := ListNode{Val: 9, Next: nil}
	node2 := ListNode{Val: 9, Next: nil}
	node3 := ListNode{Val: 9, Next: nil}
	node4 := ListNode{Val: 9, Next: nil}
	node5 := ListNode{Val: 9, Next: nil}
	node6 := ListNode{Val: 9, Next: nil}
	node7 := ListNode{Val: 9, Next: nil}
	l1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7

	l2 := ListNode{Val: 9, Next: nil}
	node21 := ListNode{Val: 9, Next: nil}
	node31 := ListNode{Val: 9, Next: nil}
	node41 := ListNode{Val: 9, Next: nil}
	l2.Next = &node21
	node21.Next = &node31
	node31.Next = &node41

	res := addTwoNumbers(&l1, &l2)
	outputRes := res.AsArray()
	output := []int{8, 9, 9, 9, 0, 0, 0, 1}

	if !CompareArrays(output, outputRes) {
		t.Errorf("Expected %v, got %v", output, outputRes)
	}
}

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
	// not yet done.
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

func TestHasCycle(t *testing.T) {
	head := ListNode{Val: 0, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 2, Next: nil}
	node4 := ListNode{Val: 3, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node3

	if !HasCycle(&head) {
		t.Errorf("Expected true, but got false")
	}
}

func TestHasCycle2(t *testing.T) {
	head := ListNode{Val: 0, Next: nil}

	if HasCycle(&head) {
		t.Errorf("Expected false, but got tru")
	}
}

func TestHasCycle3(t *testing.T) {
	head := ListNode{Val: 0, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 2, Next: nil}
	node4 := ListNode{Val: 3, Next: nil}
	node5 := ListNode{Val: 4, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	if HasCycle(&head) {
		t.Errorf("Expected false, but got true")
	}
}

func TestMergeTwoLists(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 4, Next: nil}
	node4 := ListNode{Val: 8, Next: nil}
	node5 := ListNode{Val: 10, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	head2 := ListNode{Val: 1, Next: nil}
	node22 := ListNode{Val: 3, Next: nil}
	node23 := ListNode{Val: 4, Next: nil}
	node24 := ListNode{Val: 9, Next: nil}
	node25 := ListNode{Val: 11, Next: nil}

	head2.Next = &node22
	node22.Next = &node23
	node23.Next = &node24
	node24.Next = &node25

	head3 := MergeTwoLists(&head, &head2)

	fmt.Println(head3.AsArray())
}

func TestRemoveElements(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 4, Next: nil}
	node4 := ListNode{Val: 8, Next: nil}
	node5 := ListNode{Val: 10, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	outputs := []int{1, 2, 8, 10}

	res := RemoveElements(&head, 4)
	resArr := res.AsArray()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] != resArr[i] {
			t.Errorf("Expected %v, got %v", outputs, resArr)
		}
	}
}

func TestRemoveElements2(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 4, Next: nil}
	node4 := ListNode{Val: 8, Next: nil}
	node5 := ListNode{Val: 10, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	outputs := []int{4, 8, 10}

	res := RemoveElements(&head, 1)
	resArr := res.AsArray()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] != resArr[i] {
			t.Errorf("Expected %v, got %v", outputs, resArr)
		}
	}
}

func TestRemoveElements3(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 10, Next: nil}
	node3 := ListNode{Val: 4, Next: nil}
	node4 := ListNode{Val: 8, Next: nil}
	node5 := ListNode{Val: 10, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	outputs := []int{1, 4, 8}

	res := RemoveElements(&head, 10)
	resArr := res.AsArray()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] != resArr[i] {
			t.Errorf("Expected %v, got %v", outputs, resArr)
		}
	}
}

func TestRemoveElements4(t *testing.T) {
	head := ListNode{Val: 10, Next: nil}
	node2 := ListNode{Val: 10, Next: nil}
	node3 := ListNode{Val: 10, Next: nil}
	node4 := ListNode{Val: 10, Next: nil}
	node5 := ListNode{Val: 10, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	outputs := []int{}

	res := RemoveElements(&head, 10)
	resArr := res.AsArray()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] != resArr[i] {
			t.Errorf("Expected %v, got %v", outputs, resArr)
		}
	}
}

func TestDeleteDuplicates(t *testing.T) {
	head := ListNode{Val: 10, Next: nil}
	node2 := ListNode{Val: 10, Next: nil}
	node3 := ListNode{Val: 10, Next: nil}
	node4 := ListNode{Val: 10, Next: nil}
	node5 := ListNode{Val: 10, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	outputs := []int{10}

	res := DeleteDuplicates(&head)
	resArr := res.AsArray()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] != resArr[i] {
			t.Errorf("Expected %v, got %v", outputs, resArr)
		}
	}

}

func TestDeleteDuplicates2(t *testing.T) {
	head := ListNode{Val: 10, Next: nil}
	node2 := ListNode{Val: 101, Next: nil}
	node3 := ListNode{Val: 102, Next: nil}
	node4 := ListNode{Val: 103, Next: nil}
	node5 := ListNode{Val: 101, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	outputs := []int{10, 101, 102, 103}

	res := DeleteDuplicates(&head)
	resArr := res.AsArray()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] != resArr[i] {
			t.Errorf("Expected %v, got %v", outputs, resArr)
		}
	}

}

func TestDeleteDuplicates3(t *testing.T) {
	head := ListNode{Val: 10, Next: nil}
	node2 := ListNode{Val: 101, Next: nil}
	node3 := ListNode{Val: 102, Next: nil}
	node4 := ListNode{Val: 102, Next: nil}
	node5 := ListNode{Val: 101, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	outputs := []int{10, 101}

	res := DeleteDuplicates(&head)
	resArr := res.AsArray()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] != resArr[i] {
			t.Errorf("Expected %v, got %v", outputs, resArr)
		}
	}

}

func TestDeleteDuplicates4(t *testing.T) {
	head := ListNode{Val: 10, Next: nil}

	outputs := []int{10}

	res := DeleteDuplicates(&head)
	resArr := res.AsArray()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] != resArr[i] {
			t.Errorf("Expected %v, got %v", outputs, resArr)
		}
	}

}

func TestDeleteDuplicates5(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 3, Next: nil}
	node5 := ListNode{Val: 4, Next: nil}
	node6 := ListNode{Val: 4, Next: nil}
	node7 := ListNode{Val: 5, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7

	outputs := []int{1, 2, 3, 4, 5}

	res := DeleteDuplicates(&head)
	resArr := res.AsArray()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] != resArr[i] {
			t.Errorf("Expected %v, got %v", outputs, resArr)
		}
	}

}

func TestMiddleNode(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	res := MiddleNode(&head)

	if res.Val != 3 {
		t.Errorf("Expected output as 3. but got %v", res.Val)
	}

}

func TestMiddleNode2(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}
	node6 := ListNode{Val: 6, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6

	res := MiddleNode(&head)

	if res.Val != 4 {
		t.Errorf("Expected output as 4. but got %v", res.Val)
	}

}

func TestMiddleNode3(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	res := MiddleNode(&head)

	if res.Val != 1 {
		t.Errorf("Expected output as 1. but got %v", res.Val)
	}

}

func TestMiddleNode4(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node1 := ListNode{Val: 2, Next: nil}
	head.Next = &node1
	res := MiddleNode(&head)

	if res.Val != 2 {
		t.Errorf("Expected output as 1. but got %v", res.Val)
	}

}

func TestRemoveNthFromEnd(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	res := RemoveNthFromEnd(&head, 1)

	if res != nil {
		t.Errorf("Expected nil head but got something else - %v", res.Val)
	}
}

func TestRemoveNthFromEnd1(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}
	node6 := ListNode{Val: 6, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6

	res := RemoveNthFromEnd(&head, 2)
	output := []int{1, 2, 3, 4, 6}
	outputRes := res.AsArray()

	if !CompareArrays(output, outputRes) {
		t.Errorf("Expected %v, got %v", output, outputRes)
	}

}

func TestRemoveNthFromEnd2(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}
	node6 := ListNode{Val: 6, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6

	res := RemoveNthFromEnd(&head, 6)
	output := []int{2, 3, 4, 5, 6}
	outputRes := res.AsArray()

	if !CompareArrays(output, outputRes) {
		t.Errorf("Expected %v, got %v", output, outputRes)
	}

}

func TestRemoveNthFromEnd4(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node1 := ListNode{Val: 2, Next: nil}
	head.Next = &node1

	res := RemoveNthFromEnd(&head, 1)
	output := []int{1}
	outputRes := res.AsArray()

	if !CompareArrays(output, outputRes) {
		t.Errorf("Expected %v, got %v", output, outputRes)
	}

}

func CompareArrays(a1 []int, a2 []int) bool {

	l1 := len(a1)
	l2 := len(a2)
	if l1 != l2 {
		return false
	}

	for i := 0; i < l1; i++ {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}
