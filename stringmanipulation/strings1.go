package stringmanipulation

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"

	"sathish.com/gomod/stacks"
)

// https://leetcode.com/problems/permutation-in-string/
func CheckInclusion(s1 string, s2 string) bool {

	s1map := make(map[rune]int)
	for _, c := range s1 {
		_, ok := s1map[c]
		if ok {
			s1map[c] = s1map[c] + 1
		} else {
			s1map[c] = 1
		}
	}

	var s1mapcopy map[rune]int = make(map[rune]int)
	deepCopyMap(s1map, s1mapcopy)
	for i := 0; i < len(s2); i++ {
		for j := i; j < len(s2); j++ {
			_, ok := s1mapcopy[rune(s2[j])]
			if ok {
				s1mapcopy[rune(s2[j])]--
				if isMapempty(s1mapcopy) {
					return true
				}
			} else {
				break
			}
		}
		s1mapcopy = make(map[rune]int)
		deepCopyMap(s1map, s1mapcopy)
	}

	return false
}

// https://leetcode.com/problems/valid-parentheses/
func CheckValidParenthesis(s string) bool {
	var st stacks.RuneStack
	for i := 0; i < len(s); i++ {
		if st.IsEmpty() {
			if checkLeftParenthesis(rune(s[i])) {
				st.Push(rune(s[i]))
				continue
			} else {
				return false
			}
		}
		l, _ := st.Peek()
		if checkLeftParenthesis(rune(s[i])) {
			st.Push(rune(s[i]))
		} else {
			if checkEqualParenthesis(l, rune(s[i])) {
				st.Pop()
			} else {
				return false
			}
		}
	}

	for !st.IsEmpty() {
		r, _ := st.Pop()
		l, _ := st.Peek()
		if checkEqualParenthesis(l, r) {
			st.Pop()
		} else {
			return false
		}
	}
	return true
}

func checkEqualParenthesis(left rune, right rune) bool {
	if left == '{' && right == '}' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	if left == '(' && right == ')' {
		return true
	}
	return false
}

func checkLeftParenthesis(left rune) bool {
	if left == '{' || left == '(' || left == '[' {
		return true
	}
	return false
}

func isMapempty(mp map[rune]int) bool {
	bret := true
	for _, v := range mp {
		if v == 0 {
			continue
		} else {
			bret = false
			break
		}
	}
	return bret
}

func deepCopyMap(src map[rune]int, dst map[rune]int) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}
	if dst == nil {
		return errors.New("dst cannot be nil")
	}
	jsonStr, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonStr, &dst)
	if err != nil {
		return err
	}
	return nil
}

//https://leetcode.com/problems/shuffle-string/
func RestoreString(s string, indices []int) string {

	var restoredString []byte = make([]byte, len(s))

	for i, val := range s {
		restoredString[indices[i]] = byte(val)
	}

	return string(restoredString)
}

func BalancedStringSplit(s string) int {
	// find the balanced set of Ls and Rs
	balcnt := 0
	lsideval := s[0]
	lsidecnt := 1
	rsidecnt := 0
	for i := 1; i < len(s); i++ {
		if lsidecnt == 0 {
			lsideval = s[i]
			lsidecnt += 1
			continue
		}
		if s[i] == lsideval {
			lsidecnt++
		} else {
			rsidecnt++
		}
		if lsidecnt == rsidecnt {
			balcnt += 1
			lsidecnt = 0
			rsidecnt = 0
		}
	}
	return balcnt
}

func IsPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	if len(s) > 2 {
		return IsPalindrome(s[1 : len(s)-1])
	}
	return true
}

func IsPalindromeCaseInsensitive(s string) bool {
	return IsPalindrome(strings.ToLower(s))
}

func IsPalindromeCharsOnly(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	news := reg.ReplaceAllString(s, "")
	if news == "" {
		return false
	}
	return IsPalindrome(strings.ToLower(news))
}

//return the sub string that is a palindrome.
func AllStrings(s string) []string {
	if len(s) == 0 {
		var v []string = nil
		var v1 string
		v = append(v, v1)
		return v
	}
	firstElem := s[0]
	restElem := s[1:]

	var currCombos []string = nil
	// get an slice of slice from the rest of the elements.
	restCombo := AllStrings(restElem)

	//combine that with first element and create a new array of array
	for _, str := range restCombo {
		newarray := string(firstElem) + str
		currCombos = append(currCombos, string(newarray))
	}
	// return the combination -
	return append(restCombo, currCombos...)

}

// recursive and highly time consuming. not a good one.
func IsvalidPalindromeAfterOneDelete(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, _ := range s {
		news := s[:i] + s[i+1:]
		if IsPalindrome(news) {
			return true
		}
	}
	return false
}

func LengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	sarry := strings.SplitAfter(s, " ")
	return len(sarry[len(sarry)-1])
}

func ConvertToTitle(columnNumber int) string {
	alpha := "ZABCDEFGHIJKLMNOPQRSTUVWXY"
	title := ""
	num := columnNumber
	for num != 0 {
		lastdigit := num % 26
		/* 		if lastdigit == 0 {
		   			title = "z" + title
		   		} else {
		   			title = string(alpha[lastdigit-1]) + title
		   		}
		*/
		title = string(alpha[lastdigit]) + title
		if lastdigit == 0 {
			num = num/26 - 1
		} else {
			num = num / 26
		}
	}
	return title
}

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap := make(map[rune]int)
	for _, c := range s {
		_, ok := smap[c]
		if ok {
			smap[c] = smap[c] + 1
		} else {
			smap[c] = 1
		}
	}

	tmap := make(map[rune]int)
	for _, c := range t {
		_, ok := tmap[c]
		if ok {
			tmap[c] = tmap[c] + 1
		} else {
			tmap[c] = 1
		}
	}

	for k, _ := range tmap {
		_, ok := smap[k]
		if ok {
			if smap[k] != tmap[k] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func AreOccurrencesEqual(s string) bool {
	smap := make(map[rune]int)
	for _, c := range s {
		_, ok := smap[c]
		if ok {
			smap[c] = smap[c] + 1
		} else {
			smap[c] = 1
		}
	}

	checkOccur := 0
	for _, v := range smap {
		checkOccur = v
		break
	}

	for _, v := range smap {
		if checkOccur != v {
			return false
		}
	}
	return true
}

//https://leetcode.com/problems/maximum-number-of-words-you-can-type/
func CanBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	wc := len(words)
	wcbroken := 0
	for _, w := range words {
		for _, c := range brokenLetters {
			if strings.Contains(w, string(c)) {
				wcbroken += 1
				break
			}
		}
	}
	return wc - wcbroken
}

//https://leetcode.com/problems/largest-odd-number-in-string/submissions/
func LargestOddNumber(num string) string {

	inum, _ := strconv.Atoi(num)

	for inum != 0 {
		if inum%2 == 1 {
			break
		} else {
			inum = inum / 10
		}
	}

	if inum == 0 {
		return ""
	}

	return strconv.Itoa(inum)

}

func LargestOddNumber2(num string) string {
	l := len(num) - 1
	for len(num) > 0 {
		inum, _ := strconv.Atoi(string(num[l]))
		if inum%2 != 0 {
			return num
		} else {
			num = num[0:l]
			l = len(num) - 1
		}
	}
	return ""
}

func ReverseString(s []byte) {
	j := len(s) - 1
	half := len(s) / 2
	for i := 0; i < half; i++ {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		j--
	}
}

func ReverseStr(s []byte) string {
	j := len(s) - 1
	half := len(s) / 2
	for i := 0; i < half; i++ {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		j--
	}
	return string(s)
}

func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	var reverseStr string
	for _, w := range words {
		revWord := ReverseStr([]byte(w))
		reverseStr = reverseStr + " " + revWord
	}
	return strings.TrimLeft(reverseStr, " ")
}

//https://leetcode.com/problems/reverse-vowels-of-a-string/
func ReverseVowels(s string) string {
	// assumption is that all string values are in ascii. Therefore converted to byte and not rune
	rstr := []byte(s)
	l := len(rstr)
	vowels := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1}
	i := 0
	j := l - 1

	for i < j && i < l-1 {
		_, ok := vowels[rstr[i]]
		if ok {
			_, ok2 := vowels[rstr[j]]
			if ok2 {
				rstr[i], rstr[j] = rstr[j], rstr[i]
				i++
				j--
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return string(rstr)
}

func ReverseStrtillK(s string, k int) string {
	pos := 0
	remainingstr := ""
	if k < len(s)-1 {
		pos = k
		remainingstr = s[pos:]
	} else {
		pos = len(s)
	}
	revstr := []byte(s[:pos])
	ReverseString(revstr)
	fullstr := string(revstr) + remainingstr
	return fullstr
}

//https://leetcode.com/problems/delete-characters-to-make-fancy-string/
func MakeFancyString(s string) string {
	var lastchar rune
	var cnt int
	var retstr string
	for i, c := range s {
		if i == 0 {
			lastchar = c
			cnt = 1
			retstr = string(c)
		} else {
			if lastchar == c {
				cnt = cnt + 1
			} else {
				if cnt >= 2 {
					retstr = retstr + string(lastchar)
				}
				retstr = retstr + string(c)
				lastchar = c
				cnt = 1
			}
		}
	}
	if cnt >= 2 {
		retstr = retstr + string(lastchar)
	}
	return retstr
}

func IsPrefixString(s string, words []string) bool {
	combs := ""
	start := 0
	for i := 0; i < len(words); i++ {
		if (start + len(words[i])) > len(s) {
			return false
		}
		if s[start:start+len(words[i])] != words[i] {
			return false
		}
		start = start + len(words[i])
		combs = combs + words[i]
		if s == combs {
			return true
		}
		if start > len(s) {
			return false
		}
	}
	return false

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
func GetDecimalValue(head *ListNode) int {
	boolStr := ""
	for head != nil {
		if head.Val == 1 {
			boolStr = boolStr + "1"
		} else {
			boolStr = boolStr + "0"
		}

		head = head.Next
	}
	val, _ := strconv.ParseInt(boolStr, 2, 0)
	return int(val)
}

//https://leetcode.com/problems/number-of-valid-words-in-a-sentence/
func CountValidWords(sentence string) int {
	words := strings.Split(sentence, " ")
	cnt := 0
	for _, w := range words {
		if isvalidToken(w) {
			log.Print(w)
			cnt++
		}
	}
	return cnt
}

func isvalidToken(token string) bool {
	l := len(token)
	cntPunct := 0
	cntHyphen := 0
	if l == 0 {
		return false
	}
	for i, c := range token {
		if c == '!' || c == '.' || c == ',' {
			cntPunct++
			if l-1 != i {
				return false // only last character in the token can be punctuation mark
			}
			if cntPunct > 1 {
				return false
			}
		}
		if c >= 48 && c <= 57 {
			return false
		}
		if c == '-' {
			cntHyphen++
			if cntHyphen > 1 {
				return false
			}
		}
		if (i == 0 && c == '-') || (i == l-1 && c == '-') {
			return false
		}
		if c == '-' {
			if !(token[i-1] >= 'a' && token[i-1] <= 'z') {
				return false
			}
			if !(token[i+1] >= 'a' && token[i+1] <= 'z') {
				return false
			}
		}
	}
	return true
}
