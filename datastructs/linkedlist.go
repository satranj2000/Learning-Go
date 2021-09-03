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
