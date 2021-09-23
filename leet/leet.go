package leet

import (
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func grayCode(n int) []int {
	if n < 1 {
		return nil
	}
	var list []string = make([]string, 0, int(math.Pow(2, float64(n))))
	list = append(list, "0")
	list = append(list, "1")
	for i := 2; i <= n; i++ {

		for m := len(list) - 1; m >= 0; m-- {
			list = append(list, list[m])
		}
		half := len(list) / 2
		for k := 0; k < half; k++ {
			val := list[k]
			val = val + "0"
			list[k] = val
		}
		for k := half; k < len(list); k++ {
			val := list[k]
			val = val + "1"
			list[k] = val
		}
	}
	// convert the binary strings to integer list
	ilist, _ := ConvertBinaryStringListtoIntList(list)
	return ilist
}

func CircularPermutation(n int, start int) []int {
	var ilist []int = make([]int, 0, int(math.Pow(2, float64(n))))
	ilisttemp := grayCode(n)
	// find the starting position in the new integer list.
	var startpos int
	for i, v := range ilisttemp {
		if v == start {
			startpos = i
			break
		}
	}
	//copy from the startpos till end of the list.
	for i := startpos; i < len(ilisttemp); i++ {
		ilist = append(ilist, ilisttemp[i])
	}
	// start from zero to startpos
	for i := 0; i < startpos; i++ {
		ilist = append(ilist, ilisttemp[i])
	}

	return ilist
}

func ConvertBinaryStringListtoIntList(bstrlist []string) ([]int, bool) {
	var bIntList []int = make([]int, 0, len(bstrlist))
	for _, v := range bstrlist {
		if i, err := strconv.ParseInt(v, 2, 64); err == nil {
			bIntList = append(bIntList, int(i))
		} else {
			break
		}
	}
	if len(bIntList) < len(bstrlist) {
		return bIntList, false
	}
	return bIntList, true
}

// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/submissions/
func GetLucky(s string, k int) int {
	alphabetArray := make([]string, 0, len(s))
	for _, v := range s {
		bValue := byte(v - 96)
		alphabetArray = append(alphabetArray, strconv.Itoa(int(bValue)))
	}
	combinedSum := 0
	for _, alphabet := range alphabetArray {
		for _, individualNum := range alphabet {
			val := int(individualNum - 48)
			combinedSum += int(val)
		}
	}
	if k > 1 {
		combinedSum = combinedSumNtimes(combinedSum, k-1)
	}
	return combinedSum
}

func combinedSumNtimes(v int, k int) int {
	newv := 0
	for i := 1; i <= k; i++ {
		strValue := strconv.Itoa(v)
		for _, v := range strValue {
			val := int(v - 48)
			newv += int(val)
		}
		v = newv
		newv = 0
	}
	return v
}

func AddDigits(num int) int {
	s := strconv.Itoa(num)
	var newSum int
	if len(s) == 1 {
		return int(s[0] - 48)
	}
	for len(s) >= 2 {
		newSum = 0
		for _, v := range s {
			val := int(v - 48)
			newSum += int(val)
		}
		s = strconv.Itoa(newSum)

	}
	return newSum
}

//https://leetcode.com/problems/implement-strstr/
func StrStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}

	if haystack == "" {
		return -1
	}
	lenHaystack := len(haystack)
	for i, h := range haystack {
		if byte(h) == needle[0] {
			var bFound bool
			for j, n := range needle {
				bFound = true
				if i+j > lenHaystack-1 {
					bFound = false
					break
				}
				if byte(n) != haystack[i+j] {
					bFound = false
					break
				}
			}
			if bFound {
				return i
			}

		}
	}
	return -1
}

//https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/
func AreAlmostEqual(s1 string, s2 string) bool {

	diffCnt := 0

	if len(s1) != len(s2) {
		return false
	}
	s1vals := make([]byte, 0, 2)
	s2vals := make([]byte, 0, 2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCnt++
			s1vals = append(s1vals, s1[i])
			s2vals = append(s2vals, s2[i])
		}
		if diffCnt > 2 {
			return false
		}
	}
	if diffCnt == 0 {
		return true
	}
	if diffCnt == 2 {
		if s1vals[0] == s2vals[1] && s1vals[1] == s2vals[0] {
			return true
		} else {
			return false
		}

	}
	return false
}

func BuddyStrings(s string, goal string) bool {
	diffCnt := 0
	if len(s) != len(goal) {
		return false
	}
	s1vals := make([]byte, 0, 2)
	s2vals := make([]byte, 0, 2)
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diffCnt++
			s1vals = append(s1vals, s[i])
			s2vals = append(s2vals, goal[i])
		}
		if diffCnt > 2 {
			return false
		}
	}
	if diffCnt == 0 {
		charmap := make(map[byte]int)
		for k := 0; k < len(s); k++ {
			_, ok := charmap[byte(s[k])]
			if ok {
				return true
			} else {
				charmap[s[k]] = 1
			}
		}
		return false
	}
	if diffCnt == 2 {
		if s1vals[0] == s2vals[1] && s1vals[1] == s2vals[0] {
			return true
		} else {
			return false
		}

	}
	return false

}

//https://leetcode.com/problems/binary-subarrays-with-sum/
func NumSubarraysWithSum(nums []int, goal int) int {
	// generate a prefix sum array
	sumnums := make([]int, len(nums))
	sumnums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sumnums[i] = nums[i] + sumnums[i-1]
	}
	cnt := 0

	for i := 1; i <= len(sumnums); i++ {
		for j := i - 1; j < len(sumnums); j++ {
			diff := 0
			if (j - i) < 0 {
				diff = sumnums[j]
			} else {
				diff = sumnums[j] - sumnums[j-i]
			}
			if diff == goal {
				cnt += 1
			}
		}
	}
	return cnt
}

//https://leetcode.com/problems/truncate-sentence/
func TruncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
	var ts string
	for i, w := range words {
		if i == 0 {
			ts = w
		} else {
			if i < k {
				ts = ts + " " + w
			} else {
				break
			}
		}
	}
	return ts
}

//https://leetcode.com/problems/subarray-sums-divisible-by-k/
func SubarraysDivByK(nums []int, k int) int {
	if k < 1 {
		return 0
	}
	// generate a prefix sum array
	sumnums := make([]int, len(nums))
	sumnums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sumnums[i] = nums[i] + sumnums[i-1]
	}
	cnt := 0

	for i := 1; i <= len(sumnums); i++ {
		for j := i - 1; j < len(sumnums); j++ {
			diff := 0
			if (j - i) < 0 {
				diff = sumnums[j]
			} else {
				diff = sumnums[j] - sumnums[j-i]
			}
			if diff%k == 0 {
				cnt += 1
			}
		}
	}
	return cnt
}

//https://leetcode.com/problems/search-a-2d-matrix/
func SearchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)    // rows
	n := len(matrix[0]) // columns
	for i := 0; i < m; i++ {
		if target > matrix[i][n-1] {
			continue // switch to the next row
		}
		// if less ; check within the row
		idx := sort.SearchInts(matrix[i], target)
		if matrix[i][idx] == target {
			return true
		} else {
			return false
		}
	}
	return false
}

//https://leetcode.com/problems/search-a-2d-matrix-ii/
func SearchMatrix2(matrix [][]int, target int) bool {
	m := len(matrix)    // rows
	n := len(matrix[0]) // columns
	colidx := 0
	for i := 0; i < m; i++ {
		colidx = sort.SearchInts(matrix[i], target)
		if colidx >= n {
			continue
		}
		if matrix[i][colidx] == target {
			return true
		}
	}
	return false

}

//https://leetcode.com/problems/longest-common-prefix/
func LongestCommonPrefix(strs []string) string {
	lenarr := make([]int, len(strs))
	l := len(strs)
	min := 200 // max length of string.
	for i, s := range strs {
		lenarr[i] = len(s)
		if lenarr[i] < min {
			min = lenarr[i]
		}
	}

	if min == 0 {
		return ""
	}

	prefix := ""
	for i := 0; i < min; i++ {
		char1 := strs[0][i]
		j := 1
		for j < l {
			if char1 != strs[j][i] {
				break
			}
			j = j + 1
		}
		if j == l {
			prefix = prefix + string(char1)
		} else {
			break
		}
	}
	return prefix
}

func DetectCapitalUse(word string) bool {

	// small letters - 97 to 122 (A to Z)
	// capital letters - 65 to 90 ( a to z )

	wordlen := len(word)
	capcount := 0
	smallcount := 0
	firstcap := false

	for i, w := range word {
		if i == 0 {
			if w >= 65 && w <= 90 {
				firstcap = true
				capcount = 1
			} else {
				smallcount = 1
			}
			continue
		}
		if w >= 65 && w <= 90 {
			capcount = capcount + 1
		} else {
			smallcount = smallcount + 1
		}
	}

	if firstcap {
		if wordlen == (1 + smallcount) {
			return true
		} else {
			if wordlen == capcount {
				return true
			} else {
				return false
			}
		}
	}

	if wordlen == smallcount {
		return true
	} else {
		if wordlen == capcount {
			return true
		} else {
			return false
		}
	}

}

// https://leetcode.com/problems/word-break/ (**INTERESTING**)
func WordBreak(s string, wordDict []string) bool {
	// set a slice same as the s string and store 0 or 1 - based on if the word ending
	// matches the last character of one of the words in the wordDict
	bArr := make([]bool, len(s))
	// all values are by default false

	// loop through individual characters in the main string.
	lastendpos := 0
	for i, c := range s {
		for _, w := range wordDict {
			wl := len(w)
			// if current char is same as last character of the current word from the wordDict
			// and if the current word from wordDict is same as substring
			lastendpos = i - wl + 1
			if lastendpos <= 0 {
				//lastendpos = 0
				if byte(c) == w[wl-1] && s[0:i+1] == w {
					bArr[i] = true
				}
			} else {
				if byte(c) == w[wl-1] && s[lastendpos:i+1] == w && bArr[lastendpos-1] {
					bArr[i] = true
				}
			}

		}
	}
	return bArr[len(bArr)-1]
}

func checkIfPresent(a []int, val int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == val {
			return true
		}
	}
	return false
}

//https://leetcode.com/problems/second-largest-digit-in-a-string/
func SecondHighest(s string) int {
	iArr := make([]int, 0, 10)
	for _, c := range s {
		// in number range
		if c >= 48 && c <= 57 {
			i, _ := strconv.Atoi(string(c))
			if !checkIfPresent(iArr, i) {
				iArr = append(iArr, i)
			}
		}
	}
	if len(iArr) >= 2 {
		sort.Ints(iArr)
		return int(byte(iArr[len(iArr)-2]))
	} else {
		return -1
	}

}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func MinTimeToType(word string) int {

	initialpos := 97
	sum := 0
	for _, c := range word {

		// difference between the first charcter and next in clockwise
		diff1 := int(Abs(int64(c - rune(initialpos))))
		// difference between the first charcter and next in anti-clockwise
		diff2 := int(26 - Abs(int64(c-rune(initialpos))))

		if diff1 > diff2 {
			sum = sum + diff2 + 1
			initialpos = int(c)
		} else {
			sum = sum + diff1 + 1
			initialpos = int(c)
		}
	}

	return sum

}

//https://leetcode.com/problems/maximum-matrix-sum/
func MaxMatrixSum(matrix [][]int) int64 {
	minvalue := 999999
	var sum int64
	sum = 0
	tot_negative := 0
	l := len(matrix[0])

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			absval := Abs(int64(matrix[i][j]))
			sum = sum + absval
			if matrix[i][j] < 0 {
				tot_negative++
			}
			if int(absval) < minvalue {
				minvalue = int(absval)
			}
		}
	}

	if tot_negative%2 == 0 {
		return sum
	}
	return sum - 2*int64(minvalue)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func FindGCDbetweenSmallandLarge(nums []int) int {
	sort.Ints(nums)
	return GCD(nums[0], nums[len(nums)-1])
}

func BinaryStrings(n int) []string {
	if n == 1 {
		bminstring := []string{"0", "1"}
		return bminstring
	}

	remainBinString := BinaryStrings(n - 1)

	retBinString := make([]string, 0, len(remainBinString)*2)

	for _, s := range remainBinString {
		retBinString = append(retBinString, "0"+s)
		retBinString = append(retBinString, "1"+s)
	}
	return retBinString
}

func FindDifferentBinaryString(nums []string) string {
	len := len(nums[0])
	fulllist := BinaryStrings(len)
	for _, l := range fulllist {
		bfound := false
		for _, n := range nums {
			if n == l {
				bfound = true
				break
			}

		}
		if !bfound {
			return l
		}
	}
	return ""
}

//https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
func FindDisappearedNumbers(nums []int) []int {
	newlist := make([]int, 0, len(nums))
	sort.Ints(nums)
	l := len(nums)
	lastpos := 0
	for i := 1; i <= l; i++ {
		j := lastpos
		bfound := false
		for j < l {
			if nums[j] == i {
				bfound = true
				lastpos = j
				break
			} else {
				if nums[j] > i {
					break
				} else {
					j = j + 1
				}
			}
		}
		if !bfound {
			newlist = append(newlist, i)
		}
	}

	return newlist
}

//https://leetcode.com/problems/first-unique-character-in-a-string/
func FirstUniqChar(s string) int {
	// all characters are in small letters -  97 to 122
	alphabetArr := make([]int, 26)
	for _, c := range s {
		alphabetArr[c-97]++
	}

	for i, c := range s {
		if alphabetArr[c-97] == 1 {
			return i
		}
	}
	return -1
}

//https://leetcode.com/problems/sort-characters-by-frequency/
func FrequencySort(s string) string {
	freqMap := make(map[byte]int)
	for _, c := range s {
		_, ok := freqMap[byte(c)]
		if ok {
			freqMap[byte(c)]++
		} else {
			freqMap[byte(c)] = 1
		}
	}
	revfreqMap := reverseMap(freqMap)
	sortKeys := make([]int, 0, len(revfreqMap))
	for k := range revfreqMap {
		sortKeys = append(sortKeys, k)
	}
	sort.Ints(sortKeys)

	freqsortedStr := ""
	for i := len(sortKeys) - 1; i >= 0; i-- {
		alphas := revfreqMap[sortKeys[i]]
		for _, a := range alphas {
			j := 0
			for j < sortKeys[i] {
				freqsortedStr = freqsortedStr + string(a)
				j++
			}
		}
	}
	return freqsortedStr
}

func reverseMap(m map[byte]int) map[int][]byte {
	revMap := make(map[int][]byte)
	for k, v := range m {
		_, ok := revMap[v]
		if ok {
			revMap[v] = append(revMap[v], k)
		} else {
			arr := []byte{k}
			revMap[v] = arr
		}
	}
	return revMap
}

//https://leetcode.com/problems/top-k-frequent-elements/
func TopKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, c := range nums {
		_, ok := freqMap[c]
		if ok {
			freqMap[c]++
		} else {
			freqMap[c] = 1
		}
	}

	revfreqMap := reverseIntMap(freqMap)
	sortVals := make([]int, 0, len(freqMap))
	for _, v := range freqMap {
		sortVals = append(sortVals, v)
	}
	sort.Ints(sortVals)
	sortVals = removeDuplicateValues(sortVals)

	topFreqlist := make([]int, 0, len(nums))

	for i := len(sortVals) - 1; i >= 0; i-- {
		topFreqlist = append(topFreqlist, revfreqMap[sortVals[i]]...)
		if len(topFreqlist) >= k {
			break
		}
	}

	return topFreqlist
}

func reverseIntMap(m map[int]int) map[int][]int {
	revMap := make(map[int][]int)
	for k, v := range m {
		_, ok := revMap[v]
		if ok {
			revMap[v] = append(revMap[v], k)
		} else {
			arr := []int{k}
			revMap[v] = arr
		}
	}
	return revMap
}

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//https://leetcode.com/problems/sort-array-by-increasing-frequency/
func FrequencyIntSort(nums []int) []int {
	freqMap := make(map[int]int)
	for _, c := range nums {
		_, ok := freqMap[c]
		if ok {
			freqMap[c]++
		} else {
			freqMap[c] = 1
		}
	}
	revfreqMap := reverseIntMap(freqMap)
	sortKeys := make([]int, 0, len(revfreqMap))
	for k := range revfreqMap {
		sortKeys = append(sortKeys, k)
	}
	sort.Ints(sortKeys)

	freqsorted := make([]int, 0, len(nums))
	// in increasing order
	for i := 0; i < len(sortKeys); i++ {
		alphas := revfreqMap[sortKeys[i]]
		// if multiple values have same frequency, sort by decreasing order as per the problem.
		sort.Ints(alphas)
		for k := len(alphas) - 1; k >= 0; k-- {
			j := 0
			for j < sortKeys[i] {
				freqsorted = append(freqsorted, alphas[k])
				j++
			}
		}
	}
	return freqsorted

}

func ThirdMaxDistinct(nums []int) int {
	freqMap := make(map[int]int)
	for _, c := range nums {
		_, ok := freqMap[c]
		if ok {
			freqMap[c]++
		} else {
			freqMap[c] = 1
		}
	}
	//revfreqMap := reverseIntMap(freqMap)
	sortKeys := make([]int, 0, len(freqMap))
	for k := range freqMap {
		sortKeys = append(sortKeys, k)
	}
	sort.Ints(sortKeys)
	len := len(sortKeys)

	if len >= 3 {
		return sortKeys[len-3]
	} else {
		return sortKeys[len-1]
	}
}

//https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {

	l := len(nums)
	if l < 1 {
		return 0
	}
	cnt := 1
	updatepos := 1
	j := 1

	for i := 0; i < l; {
		firstnum := nums[i]
		for j = i + 1; j < l; j++ {
			if nums[j] == firstnum {
				continue
			} else {
				cnt++
				firstnum = nums[j]
				nums[updatepos] = nums[j]
				i = j
				updatepos++
			}

		}
		if j >= l {
			break
		}
	}
	return cnt
}

//https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/submissions/
func MinimumDifference(nums []int, k int) int {
	sort.Ints(nums)

	l := 0
	r := k - 1
	leng := len(nums)
	minval := math.MaxInt32
	if leng <= 1 {
		return nums[0]
	}

	for r < leng {
		newmin := nums[r] - nums[l]
		if newmin < int(minval) {
			minval = newmin
		}
		l++
		r++
	}
	return minval
}

//https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/submissions/
func NumOfStrings(patterns []string, word string) int {

	cnt := 0
	for _, s := range patterns {
		if strings.Contains(word, s) {
			cnt++
		}
	}
	return cnt
}

//https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
func IsSumEqual(firstWord string, secondWord string, targetWord string) bool {
	firstWordNum := 0
	firstWordNumStr := make([]byte, len(firstWord))

	secondWordNum := 0
	secondWordNumStr := make([]byte, len(secondWord))

	targetWordNum := 0
	targetWordNumStr := make([]byte, len(targetWord))

	for i, w := range firstWord {
		firstWordNumStr[i] = byte(w) - 97
	}
	firstWordNum, _ = strconv.Atoi(string(firstWordNumStr))

	for i, w := range secondWord {
		secondWordNumStr[i] = byte(w) - 97
	}
	secondWordNum, _ = strconv.Atoi(string(firstWordNumStr))

	for i, w := range targetWord {
		targetWordNumStr[i] = byte(w) - 97
	}
	str := string(targetWordNumStr)
	targetWordNum, _ = strconv.Atoi(str)

	return (firstWordNum + secondWordNum) == targetWordNum
}

func TwoSum(nums []int, target int) []int {

	targArr := [2]int{-1, -1}
	l := len(nums)
	targetMet := false
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				targArr[0] = i
				targArr[1] = j
				targetMet = true
				break
			}
		}
		// exit the outer loop as well if target is met.
		if targetMet {
			break
		}
	}

	return targArr[:]

}

//https://leetcode.com/problems/merge-sorted-array/
func MergeWithin(nums1 []int, m int, nums2 []int, n int) {

	k := m + n - 1
	i := m - 1
	j := n - 1

	for k >= 0 {
		if j == -1 {
			break
		}
		if i == -1 {
			for k >= 0 {
				nums1[k] = nums2[j]
				k--
				j--
			}
			break
		}
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}

	}
}

//https://leetcode.com/problems/intersection-of-two-arrays-ii/
func Intersect(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	lnums1 := len(nums1)
	lnums2 := len(nums2)

	intNums := make([]int, 0, lnums1+lnums2)

	j := 0
	for i := 0; i < lnums1; {
		if nums1[i] == nums2[j] {
			intNums = append(intNums, nums1[i])
			i++
			j++
		} else {
			if nums1[i] < nums2[j] {
				i++
			} else {
				j++
			}
		}
		if j >= lnums2 {
			break
		}

	}
	return intNums
}

//https://leetcode.com/problems/intersection-of-two-arrays/
// only return unique values.
func IntersectionUnique(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	lnums1 := len(nums1)
	lnums2 := len(nums2)

	intNums := make([]int, 0, lnums1+lnums2)

	j := 0
	for i := 0; i < lnums1; {
		if nums1[i] == nums2[j] {
			val := nums1[i]
			intNums = append(intNums, val)
			for i < lnums1 && nums1[i] == val {
				i++
			}
			for j < lnums2 && nums2[j] == val {
				j++
			}
		} else {
			if nums1[i] < nums2[j] {
				i++
			} else {
				j++
			}
		}
		if j >= lnums2 {
			break
		}

	}
	return intNums

}

//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func MaxProfit(prices []int) int {

	max := 0.0
	l := len(prices)

	// o(n2) - so, this might be considered slow. need to find alternatives.
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if prices[i] < prices[j] {
				max = math.Max(max, float64(prices[j]-prices[i]))
			}
		}
	}
	return int(max)
}

// same question as above. but, using kadane's algorithm.
func MaxProfit2(prices []int) int {
	curmax := 0.0
	maxsofar := 0.0
	l := len(prices)

	for i := 1; i < l; i++ {
		diff := prices[i] - prices[i-1]
		curmax += float64(diff)
		curmax = math.Max(0.0, curmax)
		maxsofar = math.Max(curmax, maxsofar)

	}
	return int(maxsofar)
}

//https://leetcode.com/problems/reshape-the-matrix/submissions/
func MatrixReshape(mat [][]int, r int, c int) [][]int {

	inColSiz := len(mat[0])
	inRowSiz := len(mat)

	// does not match, return same matrix
	if inColSiz*inRowSiz != r*c {
		return mat
	}

	// same size ; return same matrix.
	if r == inRowSiz && c == inColSiz {
		return mat
	}

	outmat := make([][]int, r)
	currRow := 0
	currCol := 0

	for i := 0; i < inRowSiz; i++ {
		for j := 0; j < inColSiz; j++ {
			if currCol == 0 {
				outmat[currRow] = make([]int, c)
				outmat[currRow][currCol] = mat[i][j]
			} else {
				if currCol < c {
					outmat[currRow][currCol] = mat[i][j]
				} else {
					currRow++
					currCol = 0
					outmat[currRow] = make([]int, c)
					outmat[currRow][currCol] = mat[i][j]

				}

			}
			currCol++
		}
	}
	return outmat
}

func IsValidSudoku(board [][]byte) bool {

	// check for rows.
	for _, r := range board {
		valCntMap := make(map[byte]int, 9)
		for _, v := range r {
			if v == '.' {
				continue
			}
			if _, ok := valCntMap[v]; ok {
				return false
			} else {
				valCntMap[v] = 1
			}
		}
	}

	// check for columns
	for c := 0; c < 9; c++ {
		valCntMap := make(map[byte]int, 9)
		for r := 0; r < 9; r++ {
			v := board[r][c]
			if v == '.' {
				continue
			}
			if _, ok := valCntMap[v]; ok {
				return false
			} else {
				valCntMap[v] = 1
			}
		}
	}

	// for a block of 3 x 3
	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			valCntMap := make(map[byte]int, 9)
			for i := r; i < r+3; i++ {
				for j := c; j < c+3; j++ {
					v := board[i][j]
					if v == '.' {
						continue
					}
					if _, ok := valCntMap[v]; ok {
						return false
					} else {
						valCntMap[v] = 1
					}
				}
			}
		}
	}

	return true
}

//https://leetcode.com/problems/ransom-note/
func CanConstruct(ransomNote string, magazine string) bool {

	magmap := make(map[rune]int)
	for _, c := range magazine {
		_, ok := magmap[c]
		if ok {
			magmap[c] = magmap[c] + 1
		} else {
			magmap[c] = 1
		}
	}

	ranmap := make(map[rune]int)
	for _, c := range ransomNote {
		_, ok := ranmap[c]
		if ok {
			ranmap[c] = ranmap[c] + 1
		} else {
			ranmap[c] = 1
		}
	}

	for k, v := range ranmap {
		if v2, ok := magmap[k]; ok {
			if v > v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func Binarysearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var mid int

	for l <= r {
		mid = l + ((r - l) / 2)
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// external function to be call. for now, it is just a function that randomly returns true or false
func isBadVersion(version int) bool {
	val := rand.Int()
	if val%2 == 0 {
		return true
	} else {
		return false
	}
}

//https://leetcode.com/problems/first-bad-version/ - used the description section of leetcode to understand and code.
func firstBadVersion(n int) int {
	start := 0
	end := n
	var mid int

	for start < end {
		mid = start + ((end - start) / 2)
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

//https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var mid int

	for l <= r {
		mid = l + ((r - l) / 2)
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func SortedSquares(nums []int) []int {
	var l int

	bFoundPos := false
	// find the position of the index where the array becomes zero or greater than zero.
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			l = i
			bFoundPos = true
			break
		}
	}

	if !bFoundPos {
		l = len(nums)
	}

	j := l - 1
	i := l
	sortedNums := make([]int, len(nums))
	k := 0
	for i < len(nums) && j >= 0 {
		posSqrs := nums[i] * nums[i]
		negSqrs := nums[j] * nums[j]
		if posSqrs < negSqrs {
			sortedNums[k] = posSqrs
			i++
			k++
		} else {
			if posSqrs > negSqrs {
				sortedNums[k] = negSqrs
				j--
				k++
			} else {
				sortedNums[k] = posSqrs
				k++
				sortedNums[k] = negSqrs
				k++
				i++
				j--
			}
		}

	}
	for i < len(nums) {
		posSqrs := nums[i] * nums[i]
		sortedNums[k] = posSqrs
		k++
		i++
	}
	for j >= 0 {
		negSqrs := nums[j] * nums[j]
		sortedNums[k] = negSqrs
		k++
		j--
	}
	return sortedNums
}

//rotate a slice
func Rotate(nums []int, k int) {
	l := len(nums)
	var rotPos int
	if k < l {
		rotPos = l - k
	} else {
		rotPos = l - (k % l)
	}

	rotatedNums := make([]int, len(nums))
	j := 0
	// take the last part first
	for i := rotPos; i < l; i++ {
		rotatedNums[j] = nums[i]
		j++
	}
	for i := 0; i < rotPos; i++ {
		rotatedNums[j] = nums[i]
		j++
	}

	for i := 0; i < len(rotatedNums); i++ {
		nums[i] = rotatedNums[i]
	}
}

func MoveZeroes(nums []int) {

	var zercnt int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zercnt++
		}
	}

	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}

	for k < len(nums) {
		nums[k] = 0
		k++
	}

}

//https://leetcode.com/problems/remove-element/
// this same as earlier problem of moving zeros. Except we replace zero with a value.
func removeElement(nums []int, val int) int {

	var valcnt int

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			valcnt++
		}
	}

	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

//https://leetcode.com/problems/maximum-product-of-three-numbers/
func maximumProduct(nums []int) int {

	sort.Ints(nums)
	l := len(nums)
	out1 := math.MinInt64

	if nums[0] < 0 && nums[1] < 0 {
		out1 = nums[0] * nums[1] * nums[l-1]
	}
	out2 := nums[l-1] * nums[l-2] * nums[l-3]

	return int(math.Max(float64(out1), float64(out2)))

}

//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func TwoSumIncreasing(numbers []int, target int) []int {

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			sum := numbers[i] + numbers[j]
			if sum > target {
				break
			}
			if sum == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	// as per the definition, there will always be a solution.
	return []int{-1, -1}
}
