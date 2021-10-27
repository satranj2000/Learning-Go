package stringmanipulation_test

import (
	"testing"

	"sathish.com/gomod/stringmanipulation"
)

func TestReverseString(t *testing.T) {
	//case 1
	inputstr := [][]byte{
		{'h', 'e', 'l', 'l', 'o'},
		{'H', 'a', 'n', 'n', 'a', 'h'},
		{'h'},
		{'a', 'a'},
		{'a', 'b'},
		{'a', 'b', 'c'},
	}

	for i, byteArr := range inputstr {
		stringmanipulation.ReverseString(byteArr)
		if i == 0 && (string(byteArr) != "olleh") {
			t.Errorf("Expected olleh, but got something else")
		}
		if i == 1 && (string(byteArr) != "hannaH") {
			t.Errorf("Expected hannaH, but got something else")
		}
		if i == 2 && (string(byteArr) != "h") {
			t.Errorf("Expected h, but got something else (%v)", string(byteArr))
		}
		if i == 3 && (string(byteArr) != "aa") {
			t.Errorf("Expected h, but got something else (%v)", string(byteArr))
		}
		if i == 4 && (string(byteArr) != "ba") {
			t.Errorf("Expected ba, but got something else (%v)", string(byteArr))
		}
		if i == 5 && (string(byteArr) != "cba") {
			t.Errorf("Expected cba, but got something else (%v)", string(byteArr))
		}
	}

}

func TestReverseVowels(t *testing.T) {
	//case 1
	inputstr := []string{
		"hello",
		"Hannah",
		"h",
		"aa",
		"ab",
		"cba",
		"abe",
		"xyzei",
		"xyzoexyz",
		"Aa",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz",
		"aaaaA",
	}

	for i, s := range inputstr {
		//fmt.Printf("\n working %v", s)
		revs := stringmanipulation.ReverseVowels(s)
		if i == 0 && (revs != "holle") {
			t.Errorf("Expected holle, but got something else")
		}
		if i == 1 && (revs != "Hannah") {
			t.Errorf("Expected Hannah, but got something else")
		}
		if i == 2 && (revs != "h") {
			t.Errorf("Expected h, but got something else (%v)", revs)
		}
		if i == 3 && (revs != "aa") {
			t.Errorf("Expected h, but got something else (%v)", revs)
		}
		if i == 4 && (revs != "ab") {
			t.Errorf("Expected ab, but got something else (%v)", revs)
		}
		if i == 5 && (revs != "cba") {
			t.Errorf("Expected cba, but got something else (%v)", revs)
		}
		if i == 6 && (revs != "eba") {
			t.Errorf("Expected eba, but got something else (%v)", revs)
		}
		if i == 7 && (revs != "xyzie") {
			t.Errorf("Expected xyzie, but got something else (%v)", revs)
		}
		if i == 8 && (revs != "xyzeoxyz") {
			t.Errorf("Expected xyzeoxyz, but got something else (%v)", revs)
		}
		if i == 9 && (revs != "aA") {
			t.Errorf("Expected xyzeoxyz, but got something else (%v)", revs)
		}
	}

}

func TestReverseStrtillK(t *testing.T) {
	inputstr := []string{
		"abcdefg",
		"abcd",
		"ab",
	}

	for i, s := range inputstr {
		//fmt.Printf("\n working %v", s)
		retval := stringmanipulation.ReverseStrtillK(s, 2)
		if i == 0 && retval != "bacdefg" {
			t.Errorf("Expected bacdefg, but got something else (%v)", s)
		}
		if i == 1 && retval != "bacd" {
			t.Errorf("Expected bacd, but got something else (%v)", s)
		}
		if i == 2 && retval != "ba" {
			t.Errorf("Expected bacd, but got something else (%v)", s)
		}
	}
}

func TestMakeFancyString(t *testing.T) {
	inputstr := []string{
		"leeetcode",
		"aaabaaaa",
		"aab",
		"aaaaaaa",
		"a",
		"aaaabbbb",
		"aabaabaabaa",
	}

	outputstr := []string{
		"leetcode",
		"aabaa",
		"aab",
		"aa",
		"a",
		"aabb",
		"aabaabaabaa",
	}

	for i, s := range inputstr {
		//fmt.Printf("\n working %v", s)
		retval := stringmanipulation.MakeFancyString(s)
		if retval != outputstr[i] {
			t.Errorf("Expected (%v), but got something else (%v) for the input (%v)", outputstr[i], retval, inputstr[i])
		}
	}

}

func TestIsPrefixString(t *testing.T) {
	inputstr := []string{"i", "love", "leetcode", "apples"}
	s := "iloveleetcode"
	if !stringmanipulation.IsPrefixString(s, inputstr) {
		t.Errorf("Expected true")
	}

	inputstr = []string{"apples", "i", "love", "leetcode"}
	s = "iloveleetcode"
	if stringmanipulation.IsPrefixString(s, inputstr) {
		t.Errorf("Expected false")
	}

	inputstr = []string{"i", "love", "leetcode"}
	s = "iloveleetcode"
	if !stringmanipulation.IsPrefixString(s, inputstr) {
		t.Errorf("Expected true")
	}

	inputstr = []string{"aa", "aaaa", "banana"}
	s = "a"
	if stringmanipulation.IsPrefixString(s, inputstr) {
		t.Errorf("Expected true")
	}
}

func TestGetDecimalValue(t *testing.T) {
	head := stringmanipulation.ListNode{Val: 1, Next: nil}
	node2 := stringmanipulation.ListNode{Val: 0, Next: nil}
	node3 := stringmanipulation.ListNode{Val: 1, Next: nil}
	node4 := stringmanipulation.ListNode{Val: 1, Next: nil}
	node5 := stringmanipulation.ListNode{Val: 0, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	res := stringmanipulation.GetDecimalValue(&head)
	if res != 22 {
		t.Errorf("Error. Got wrong output. Expected 22, got %v", res)
	}
}

func TestGetDecimalValue2(t *testing.T) {
	head := stringmanipulation.ListNode{Val: 1, Next: nil}
	node2 := stringmanipulation.ListNode{Val: 0, Next: nil}

	head.Next = &node2

	res := stringmanipulation.GetDecimalValue(&head)
	if res != 2 {
		t.Errorf("Error. Got wrong output. Expected 2, got %v", res)
	}
}

func TestGetDecimalValue3(t *testing.T) {
	head := stringmanipulation.ListNode{Val: 0, Next: nil}
	node2 := stringmanipulation.ListNode{Val: 0, Next: nil}

	head.Next = &node2

	res := stringmanipulation.GetDecimalValue(&head)
	if res != 3 {
		t.Errorf("Error. Got wrong output. Expected 0, got %v", res)
	}
}

func TestGetDecimalValue4(t *testing.T) {
	head := stringmanipulation.ListNode{Val: 1, Next: nil}
	node2 := stringmanipulation.ListNode{Val: 1, Next: nil}
	node3 := stringmanipulation.ListNode{Val: 1, Next: nil}
	node4 := stringmanipulation.ListNode{Val: 1, Next: nil}
	node5 := stringmanipulation.ListNode{Val: 1, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	res := stringmanipulation.GetDecimalValue(&head)
	if res != 31 {
		t.Errorf("Error. Got wrong output. Expected 31, got %v", res)
	}
}

func TestGetDecimalValue5(t *testing.T) {
	head := stringmanipulation.ListNode{Val: 0, Next: nil}
	node2 := stringmanipulation.ListNode{Val: 0, Next: nil}
	node3 := stringmanipulation.ListNode{Val: 0, Next: nil}
	node4 := stringmanipulation.ListNode{Val: 0, Next: nil}
	node5 := stringmanipulation.ListNode{Val: 0, Next: nil}

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	res := stringmanipulation.GetDecimalValue(&head)
	if res != 0 {
		t.Errorf("Error. Got wrong output. Expected 0, got %v", res)
	}
}

func TestReverseWords(t *testing.T) {
	inputstrs := []string{"God Ding", "let us write code"}
	outputsstr := []string{"doG gniD", "tel su etirw edoc"}

	for i, s := range inputstrs {
		res := stringmanipulation.ReverseWords(s)
		if res != outputsstr[i] {
			t.Errorf("String is not reversed as required. Expected %v, got %v", outputsstr[i], res)
		}
	}

}

func TestCountValidWords(t *testing.T) {
	inputStrs := []string{
		"cat and  dog",
		"!this  1-s b8d!",
		"alice and  bob are playing stone-game10",
		"he bought 2 pencils, 3 erasers, and 1  pencil-sharpener.",
		" 62   nvtk0wr4f  8 qt3r! w1ph 1l ,e0d 0n 2v 7c.  n06huu2n9 s9   ui4 nsr!d7olr  q-, vqdo!btpmtmui.bb83lf g .!v9-lg 2fyoykex uy5a 8v whvu8 .y sc5 -0n4 zo pfgju 5u 4 3x,3!wl  fv4   s  aig cf j1 a i  8m5o1  !u n!.1tz87d3 .9    n a3  .xb1p9f  b1i a j8s2 cugf l494cx1! hisceovf3 8d93 sg 4r.f1z9w   4- cb r97jo hln3s h2 o .  8dx08as7l!mcmc isa49afk i1 fk,s e !1 ln rt2vhu 4ks4zq c w  o- 6  5!.n8ten0 6mk 2k2y3e335,yj  h p3 5 -0  5g1c  tr49, ,qp9 -v p  7p4v110926wwr h x wklq u zo 16. !8  u63n0c l3 yckifu 1cgz t.i   lh w xa l,jt   hpi ng-gvtk8 9 j u9qfcd!2  kyu42v dmv.cst6i5fo rxhw4wvp2 1 okc8!  z aribcam0  cp-zp,!e x  agj-gb3 !om3934 k vnuo056h g7 t-6j! 8w8fncebuj-lq    inzqhw v39,  f e 9. 50 , ru3r  mbuab  6  wz dw79.av2xp . gbmy gc s6pi pra4fo9fwq k   j-ppy -3vpf   o k4hy3 -!..5s ,2 k5 j p38dtd   !i   b!fgj,nx qgif ",
	}
	outputs := []int{3, 0, 5, 6, 49}

	for i, s := range inputStrs {
		res := stringmanipulation.CountValidWords(s)
		if res != outputs[i] {
			t.Errorf("Expected %v, got %v", outputs[i], res)
		}
	}
}
