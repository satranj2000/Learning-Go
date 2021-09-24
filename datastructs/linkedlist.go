package datastructs

import "fmt"

type Node struct {
	data int
	next *Node
}

func setupNode(val int) Node {
	var n Node = Node{data: val, next: nil}
	return n
}

func (head *Node) Add(val int) {
	if head == nil {
		*head = setupNode(val)
	} else {
		for head.next != nil {
			head = head.next
		}
		newnode := setupNode(val)
		head.next = &newnode
	}
}

func (head *Node) ListAll() {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

//https://leetcode.com/problems/palindrome-linked-list/
//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.com/problems/reverse-linked-list/
func ReverseList(head *ListNode) *ListNode {

	var tempnode *ListNode
	var tempnodeprev *ListNode
	tempnodeprev = nil

	for head != nil {
		tempnode = head.Next
		head.Next = tempnodeprev
		tempnodeprev = head
		head = tempnode
	}
	return tempnodeprev
}

// return both head and last node
func ReverseList2(head *ListNode) (*ListNode, *ListNode) {
	newtail := head
	var tempnode *ListNode
	var tempnodeprev *ListNode
	tempnodeprev = nil

	for head != nil {
		tempnode = head.Next
		head.Next = tempnodeprev
		tempnodeprev = head
		head = tempnode
	}
	return tempnodeprev, newtail
}

//https://leetcode.com/problems/reverse-linked-list-ii/
func ReverseBetween(head *ListNode, left int, right int) *ListNode {

	runner := 1
	originhead := head // store the pointer to original head

	//left most node is pointer to the first node in the inner list (between left and right)
	var leftmostnode *ListNode
	// left node minus one - keeps track of node where we need to link back to
	var leftmostnodeminusone *ListNode
	leftmostnodeminusone = nil

	var rightmostnode *ListNode
	var rightmostnodeplusone *ListNode
	rightmostnodeplusone = nil

	for head != nil {
		if runner == left {
			leftmostnode = head
		}
		if runner == left-1 {
			leftmostnodeminusone = head
		}
		if runner == right {
			rightmostnode = head
		}
		if runner == right+1 {
			rightmostnodeplusone = head
		}
		runner++
		head = head.Next
	}
	if leftmostnodeminusone != nil {
		leftmostnodeminusone.Next = nil
	}
	rightmostnode.Next = nil

	revlisthead, revlisttail := ReverseList2(leftmostnode)

	if leftmostnodeminusone != nil {
		leftmostnodeminusone.Next = revlisthead
	}
	revlisttail.Next = rightmostnodeplusone

	if leftmostnodeminusone != nil {
		return originhead
	} else {
		return revlisthead
	}
}

//https://leetcode.com/problems/odd-even-linked-list/
func OddEvenList(head *ListNode) *ListNode {
	oddlist := &ListNode{}
	bfirstOddList := true
	bfirstEvenList := true
	runner := 0
	if head == nil {
		return head
	}
	evenlist := &ListNode{}
	oddlisttail := &ListNode{}
	evenlisttail := &ListNode{}
	for head != nil {
		if runner%2 != 0 {
			if bfirstOddList {
				oddlist = head
				oddlisttail = head
				bfirstOddList = false
			} else {
				oddlisttail.Next = evenlisttail.Next
				oddlisttail = evenlisttail.Next
			}
		} else {
			if bfirstEvenList {
				evenlist = head
				evenlisttail = head
				bfirstEvenList = false
			} else {
				evenlisttail.Next = oddlisttail.Next
				evenlisttail = oddlisttail.Next
			}
		}
		runner++
		head = head.Next
	}
	oddlisttail.Next = nil
	evenlisttail.Next = nil
	if runner != 1 {
		evenlisttail.Next = oddlist
	}

	return evenlist
}

func (node *ListNode) ListAll() {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func (node *ListNode) AsArray() []int {
	iArr := []int{}
	for node != nil {
		//fmt.Println(node.Val)
		iArr = append(iArr, node.Val)
		node = node.Next
	}
	return iArr
}

//https://leetcode.com/problems/rotate-list/
func RotateRight(head *ListNode, k int) *ListNode {

	// first find the length of the list.
	var tail *ListNode
	var newhead *ListNode

	if k == 0 {
		return head
	}

	originhead := head
	len := 0
	runner := 0
	// calculate the length of the linked list.
	for head != nil {
		if head.Next == nil {
			tail = head
			runner++
			break
		}
		head = head.Next
		runner++
	}

	// if the input linked list length is zero, return back the head.
	if runner == 0 {
		return head
	}

	len = runner
	posToCut := 0
	//set the position where to cut.
	// if it is greater than length, take reminder
	if k < len {
		posToCut = len - k
	} else {
		posToCut = len - (k % len)
	}

	// after calculation if it ends up that position to cut is zero or length of linked list,
	// return back the head
	if posToCut == 0 || posToCut == len {
		return originhead
	}

	// start from one in this case.
	runner = 1
	head = originhead
	for head != nil {
		// as soon as we arrive at the position to cut, make the next to nil and set next as new head.
		if runner == posToCut {
			newhead = head.Next
			head.Next = nil
			break
		}
		head = head.Next
		runner++
	}

	tail.Next = originhead
	head = newhead
	return newhead
}

//https://leetcode.com/problems/design-browser-history/
type BrowserHistory struct {
	url  string
	Next *BrowserHistory
	Prev *BrowserHistory
}

func Constructor(homepage string) BrowserHistory {
	home := BrowserHistory{}
	home.url = homepage
	home.Next = nil
	home.Prev = nil
	return home
}

func (this *BrowserHistory) Visit(url string) {
	nexturl := Constructor(url)

	for this.Next != nil {
		this = this.Next
	}
	this.Next = &nexturl
	nexturl.Prev = this
}

func (this *BrowserHistory) Back(steps int) string {
	i := 0
	for i < steps {
		if this != nil && this.Prev != nil {
			this = this.Prev
			i++
		} else {
			break
		}
	}
	if this != nil {
		return this.url
	} else {
		return ""
	}

}

func (this *BrowserHistory) Forward(steps int) string {
	i := 0
	for i < steps {
		if this != nil && this.Next != nil {
			this = this.Next
			i++
		} else {
			break
		}

	}
	if this != nil {
		return this.url
	} else {
		return ""
	}
}

//https://leetcode.com/problems/merge-in-between-linked-lists/
// merge linked list
func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	run := 0
	originallist1 := list1
	var aNode *ListNode
	var bNode *ListNode
	aNode, bNode = nil, nil

	for list1 != nil {
		if run == a-1 {
			aNode = list1
		}
		if run == b {
			bNode = list1
		}
		list1 = list1.Next
		run++
	}

	aNode.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = bNode.Next

	return originallist1
}

// check for a loop in linked list. - https://leetcode.com/problems/linked-list-cycle/
func HasCycle(head *ListNode) bool {
	// create 2 pointers, slow pointer and fast pointer.
	slowp := head
	fastp := head

	for slowp != nil && fastp != nil && fastp.Next != nil {
		slowp = slowp.Next
		fastp = fastp.Next.Next
		if slowp == fastp {
			return true
		}
	}

	return false

}

//https://leetcode.com/problems/merge-two-sorted-lists/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newhead *ListNode
	var currhead *ListNode
	var node *ListNode
	bfirst := true
	for l1 != nil && l2 != nil {
		var node *ListNode
		if l1.Val < l2.Val {
			node = &ListNode{Val: l1.Val, Next: nil}
			l1 = l1.Next
		} else {
			node = &ListNode{Val: l2.Val, Next: nil}
			l2 = l2.Next

		}
		if bfirst {
			newhead = node
			currhead = node
			bfirst = false
		} else {
			currhead.Next = node
			currhead = currhead.Next
		}
	}

	for l1 != nil {
		node = &ListNode{Val: l1.Val, Next: nil}
		if bfirst {
			newhead = node
			currhead = node
			bfirst = false
		} else {
			currhead.Next = node
			currhead = currhead.Next
		}
		l1 = l1.Next
	}

	for l2 != nil {
		node = &ListNode{Val: l2.Val, Next: nil}
		if bfirst {
			newhead = node
			currhead = node
			bfirst = false
		} else {
			currhead.Next = node
			currhead = currhead.Next
		}
		l2 = l2.Next
	}
	return newhead
}

func RemoveElements(head *ListNode, val int) *ListNode {

	// remove values in front.
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
	}

	if head == nil {
		return head
	}
	loop := head.Next
	prev := head
	for loop != nil {
		if loop.Val != val {
			loop = loop.Next
			prev = prev.Next
		} else {
			prev.Next = loop.Next
			loop = loop.Next
		}

	}

	return head
}

//https://leetcode.com/problems/remove-duplicates-from-sorted-list/
func DeleteDuplicates(head *ListNode) *ListNode {
	dupl := make(map[int]int)
	orighead := head
	prev := &ListNode{}

	for head != nil {
		if _, ok := dupl[head.Val]; ok {
			// time to delete the node
			prev.Next = head.Next
			head = head.Next
		} else {
			dupl[head.Val] = 1
			prev = head
			head = head.Next
		}
	}
	return orighead
}

//https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
// different from above - delete all instances of duplicates
func DeleteAllDuplicates(head *ListNode) *ListNode {
	// node values can be betweeen -100 to 100
	// should work on this.
	return nil
}

func MiddleNode(head *ListNode) *ListNode {

	cnt := 0
	curr := head

	for curr != nil {
		cnt++
		curr = curr.Next
	}

	curr = head

	for i := 0; i < cnt/2; i++ {
		curr = curr.Next
	}
	return curr
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	cnt := 0
	curr := head

	for curr != nil {
		cnt++
		curr = curr.Next
	}

	curr = head
	pos := cnt - n

	for i := 0; i < pos-1; i++ {
		curr = curr.Next
	}
	if curr == head {
		if pos <= 0 {
			head = head.Next
		} else {
			head.Next = head.Next.Next
		}
	} else {
		curr.Next = curr.Next.Next
	}
	return head
}
