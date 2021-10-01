package subsequence

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func Combine(n int, k int) [][]int {
	allcombo := [][]int{}
	currcomb := []int{}

	// when k == 1
	if k == 1 {
		for i := 1; i <= n; i++ {
			currcomb = append(currcomb, i)
			allcombo = append(allcombo, currcomb)
			currcomb = []int{}
		}
		return allcombo
	}

	// all other cases where k > 1
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			currcomb = append(currcomb, i)
			for m := 1; m < k && j+m <= n; m++ {
				currcomb = append(currcomb, j+m)
			}
			if len(currcomb) == k {
				allcombo = append(allcombo, currcomb)
			}
			currcomb = []int{}
		}
	}
	return allcombo
}

func PermuteUnique(nums []int) [][]int {
	perms := FindPermutations(nums)
	var outperms [][]int
	permmap := make(map[string]bool)
	collen := len(perms[0])
	for _, p := range perms {
		var key string
		for i := 0; i < collen; i++ {
			key = fmt.Sprint(key, p[i])
		}
		_, ok := permmap[key]
		if !ok {
			permmap[key] = true
			outperms = append(outperms, p)
		}
	}
	sort.Slice(outperms, func(i, j int) bool { return outperms[i][0] < outperms[j][0] })
	return outperms
}

func FindPermutations(nums []int) [][]int {
	if len(nums) == 0 {
		var v [][]int = nil
		var v1 []int
		v = append(v, v1)
		return v
	}
	firstElem := nums[0]
	restElems := nums[1:]
	restPermutations := FindPermutations(restElems)
	var allPermutations [][]int = nil
	for _, c := range restPermutations {
		for j := 0; j <= len(c); j++ {
			var newc []int
			newc = append(newc, c[0:j]...)
			newc = append(newc, firstElem)
			newc = append(newc, c[j:]...)
			allPermutations = append(allPermutations, newc)
		}
	}

	//sort.Slice(allPermutations, func(i, j int) bool { return allPermutations[i][0] < allPermutations[j][0] })
	return allPermutations
}

func FindAllIncreasingSubsequences(nums []int) [][]int {
	var incsub []int = FindIncreasingSubsequence(nums)
	fmt.Println(incsub)
	var list [][]int = FindAllCombo(incsub)
	return list[0 : len(list)-1]
}

func FindAllIncreasingSubseqAtleasttwo(nums []int) [][]int {
	var incsub []int = FindIncreasingSubsequence(nums)
	fmt.Println(incsub)
	var list [][]int = FindAllCombo(incsub)
	var list2 [][]int
	for _, c := range list {
		if len(c) >= 2 {
			list2 = append(list2, c)
		}
	}
	return list2
}

func FindIncreasingSubsequence(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	startv := nums[0]
	var nums2 []int
	nums2 = append(nums2, startv)
	for i := 1; i < len(nums); i++ {
		if startv <= nums[i] {
			nums2 = append(nums2, nums[i])
			startv = nums[i]
		}

	}
	return nums2
}

func FindAllComboUnique(nums []int) [][]int {
	perms := FindAllCombo(nums)
	var outperms [][]int
	permmap := make(map[string]bool)

	for _, p := range perms {
		var key string
		for i := 0; i < len(p); i++ {
			key = fmt.Sprint(key, p[i])
		}
		_, ok := permmap[key]
		if !ok {
			permmap[key] = true
			outperms = append(outperms, p)
		}
	}
	//sort.Slice(outperms, func(i, j int) bool { return outperms[i][0] < outperms[j][0] })
	return outperms
}

//reference for implementation - https://www.youtube.com/watch?v=NA2Oj9xqaZQ
func FindAllCombo(nums []int) [][]int {
	if len(nums) == 0 {
		var v [][]int = nil
		var v1 []int
		v = append(v, v1)
		return v
	}
	firstElem := nums[0]
	restElem := nums[1:]

	var currCombos [][]int = nil
	// get an slice of slice from the rest of the elements.
	restCombo := FindAllCombo(restElem)

	//combine that with first element and create a new array of array
	for _, c := range restCombo {
		l := len(c)
		//create a new slice
		newarray := make([]int, l, l+1)
		// use copy to perform a deep copy of the original c value - *** IMPORTANT ***
		copy(newarray, c)
		newarray = append(newarray, firstElem)
		//newarray = append(restCombo[i], firstElem)
		currCombos = append(currCombos, newarray)
	}
	var retcombos [][]int = nil
	retcombos = append(retcombos, restCombo...)
	retcombos = append(retcombos, currCombos...)
	// return the combination - array of arrays of the rest of Elements & ( array of array of the RestElems + first element)
	return retcombos
}

func LetterCasePermutation(s string) []string {
	if len(s) == 0 {
		var v []string
		s := ""
		v = append(v, s)
		return v
	}
	firstElem := s[0]
	restElems := s[1:]
	restPermutations := LetterCasePermutation(restElems)
	var allPermutations []string = nil
	for _, s := range restPermutations {
		var combstr string
		if unicode.IsLetter(rune(firstElem)) {
			upperFirstElem := strings.ToUpper(string(firstElem))
			lowerFirstElem := strings.ToLower(string(firstElem))
			combstr = string(upperFirstElem) + s
			allPermutations = append(allPermutations, combstr)
			combstr = string(lowerFirstElem) + s
			allPermutations = append(allPermutations, combstr)
		} else {
			combstr = string(firstElem) + s
			allPermutations = append(allPermutations, combstr)
		}
	}
	return allPermutations
}

func GeneratePascalTriangle(numRows int) [][]int {
	var pascalTrg [][]int
	firstrow := []int{1}
	pascalTrg = append(pascalTrg, firstrow)
	for i := 1; i < numRows; i++ {
		row := make([]int, 0, numRows+1)
		row = append(row, 1)
		for j := 1; j < i; j++ {
			row = append(row, pascalTrg[i-1][j-1]+pascalTrg[i-1][j])
		}
		row = append(row, 1)
		pascalTrg = append(pascalTrg, row)
	}
	return pascalTrg
}

func GetRowPascalTriangle(rowIndex int) []int {
	var pascalTrg [][]int
	firstrow := []int{1}
	pascalTrg = append(pascalTrg, firstrow)
	for i := 1; i < rowIndex; i++ {
		row := make([]int, 0, rowIndex+1)
		row = append(row, 1)
		for j := 1; j < i; j++ {
			row = append(row, pascalTrg[i-1][j-1]+pascalTrg[i-1][j])
		}
		row = append(row, 1)
		pascalTrg = append(pascalTrg, row)
	}
	return pascalTrg[rowIndex-1]
}
