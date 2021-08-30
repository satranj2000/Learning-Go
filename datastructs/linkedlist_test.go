package datastructs

import "testing"

func TestIsPalindrome(t *testing.T) {
	head := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 2, Next: nil}
	node4 := ListNode{Val: 1, Next: nil}
	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	res := IsPalindrome(&head)
	if res != true {
		t.Errorf("Wrongly identified as not a palindrome for the values [1 2 2 1]")
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
