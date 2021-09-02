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
