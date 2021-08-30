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
