package leet_test

import (
	"sort"
	"testing"

	"sathish.com/gomod/leet"
)

func ConvertBinaryStringListtoIntList(t *testing.T) {

	testStrList := []string{"0", "1"}

	testIntList, ok := leet.ConvertBinaryStringListtoIntList(testStrList)
	if !ok {
		t.Errorf("Input 2; but the return is false")
	}
	if len(testIntList) != 2 {
		t.Errorf("Input 2; but the size is wrong")
	}
	validIntList := []int{0, 1}
	for i, v := range testIntList {
		if v != validIntList[i] {
			t.Errorf("Input 2; individual values are wrong")
		}
	}

	testStrList2 := []string{"00", "1001", "1010"}
	testIntList2, ok := leet.ConvertBinaryStringListtoIntList(testStrList2)
	if !ok {
		t.Errorf("Input 3; but the return is false")
	}
	if len(testIntList2) != 3 {
		t.Errorf("Input 3; but the size is wrong")
	}
	validIntList2 := []int{0, 9, 10}
	for i, v := range testIntList2 {
		if v != validIntList2[i] {
			t.Errorf("Input 3; individual values are wrong")
		}
	}
}

func TestMaxMatrixSum(t *testing.T) {
	matrix1 := [][]int{{1, -1}, {-1, 1}}
	matrix2 := [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}
	matrix3 := [][]int{{1, -2, 3}, {1, 4, 5}, {-1, 3, 6}}
	matrix4 := [][]int{{9, -2, 3}, {1, 4, 5}, {1, 0, 6}}

	if leet.MaxMatrixSum(matrix1) != 4 {
		t.Errorf("Wrong output for the matrix 1 ")
	}

	if leet.MaxMatrixSum(matrix2) != 16 {
		t.Errorf("Wrong output for the matrix 2 ")
	}

	if leet.MaxMatrixSum(matrix3) != 26 {
		t.Errorf("Wrong output for the matrix 3 ")
	}

	if leet.MaxMatrixSum(matrix4) != 31 {
		t.Errorf("Wrong output for the matrix 4 ")
	}
}

func TestFindGCDbetweenSmallandLarge(t *testing.T) {
	arr1 := []int{2, 5, 6, 9, 10}
	arr2 := []int{7, 5, 6, 8, 3}
	arr3 := []int{3, 1023}

	if leet.FindGCDbetweenSmallandLarge(arr1) != 2 {
		t.Errorf("Wrong output for the array 1 ")
	}

	if leet.FindGCDbetweenSmallandLarge(arr2) != 1 {
		t.Errorf("Wrong output for the array 1 ")
	}

	if leet.FindGCDbetweenSmallandLarge(arr3) != 3 {
		t.Errorf("Wrong output for the array 3 ")
	}
}

func TestFirstUniqChar(t *testing.T) {
	inputstr := []string{"abc", "leetcode", "iloveleetciode", "lll"}
	outputarr := []int{0, 0, 3, -1}
	for i, s := range inputstr {
		res := leet.FirstUniqChar(s)
		if res != outputarr[i] {
			t.Errorf("Wrong output. Expected %v, got %v", outputarr[i], res)
		}
	}
}

func TestFrequencySort(t *testing.T) {
	inputstr := []string{"tree", "cccaaa", "Aabb"}
	outputstr := []string{"eetr", "cccaaa", "bbAa"}

	for i, s := range inputstr {
		res := leet.FrequencySort(s)
		if res != outputstr[i] {
			t.Errorf("Wrong output. Expected %v, got %v", outputstr[i], res)
		}
	}

}

func TestTopKFrequent(t *testing.T) {
	inputvals := [][]int{{1, 1, 1, 2, 2, 3}, {1}, {1, 2, 3}, {3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}}
	inputks := []int{2, 1, 3, 10}
	outvals := [][]int{{1, 2}, {1}, {1, 2, 3}, {1, 2, 3, 4, 5, 6, 7, 8, 10, 11}}

	for i, val := range inputvals {
		res := leet.TopKFrequent(val, inputks[i])
		sort.Ints(res)
		for j := 0; j < len(res); j++ {
			if res[j] != outvals[i][j] {
				t.Errorf("Wrong output. Expected %v, got %v", outvals[i], res)
				break
			}
		}
	}
}

func TestFrequencyIntSort(t *testing.T) {
	input := [][]int{{1, 1, 2, 2, 2, 3}, {2, 3, 1, 3, 2}, {-1, 1, -6, 4, 5, -6, 1, 4, 1}}
	output := [][]int{{3, 1, 1, 2, 2, 2}, {1, 3, 3, 2, 2}, {5, -1, 4, 4, -6, -6, 1, 1, 1}}
	// if multiple values have same frequency, sort by decreasing order as per the problem.

	for i, s := range input {
		res := leet.FrequencyIntSort(s)
		for j := 0; j < len(res); j++ {
			if res[j] != output[i][j] {
				t.Errorf("Wrong output. Expected %v, got %v", output[i], res)
				break
			}
		}
	}

}

func TestThirdMaxDistinct(t *testing.T) {
	inputs := [][]int{{3, 2, 1}, {1, 2}, {2, 2, 3, 1}}
	output := []int{1, 2, 1}

	for i, val := range inputs {
		res := leet.ThirdMaxDistinct(val)
		if res != output[i] {
			t.Errorf("Wrong output. Expected %v, got %v", output[i], res)
		}
	}

}

func TestRemoveDuplicates(t *testing.T) {
	inputs := [][]int{{1, 1, 2, 2, 3, 3}, {1, 1, 2}, {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}
	outlen := []int{3, 2, 5}

	for i, val := range inputs {
		res := leet.RemoveDuplicates(val)
		if res != outlen[i] {
			t.Errorf("Wrong output. Expected %v but got %v", outlen[i], res)
		}
	}
}

func TestMinimumDifference(t *testing.T) {
	inparam1 := [][]int{{90}, {9, 4, 1, 7}, {2, 3, 10}}
	inparam2 := []int{0, 2, 2}
	outval := []int{90, 2, 1}

	for i, v := range inparam1 {
		res := leet.MinimumDifference(v, inparam2[i])
		if res != outval[i] {
			t.Errorf("Expected %v, got %v", outval[i], res)
		}
	}
}

func TestIsSumEqual(t *testing.T) {

	res := leet.IsSumEqual("acb", "cba", "cdb")
	if !res {
		t.Errorf("Expected sum to be true but got False")
	}

	res = leet.IsSumEqual("aaa", "a", "aab")
	if res {
		t.Errorf("Expected sum to be false but got True")
	}

	res = leet.IsSumEqual("aaa", "a", "aaaa")
	if !res {
		t.Errorf("Expected sum to be true but got False")
	}
}

func TestTwoSum(t *testing.T) {
	//assumption is that there is only one valid answer.
	inputs := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
		{30, 2, 5, 3, 9},
		{0, 40, 20, 0}, // important test case for 0 value
		{-1, -2, -3, -4},
	}
	inputTargets := []int{9, 6, 6, 12, 0, -5}

	outputs := [][]int{
		{0, 1},
		{1, 2},
		{0, 1},
		{3, 4},
		{0, 3},
		{0, 3},
	}

	for i, v := range inputs {
		res := leet.TwoSum(v, inputTargets[i])

		if res[0] != outputs[i][0] || res[1] != outputs[i][1] {
			t.Errorf("Expected %v, got %v for the test #%v", outputs[i], res, i)
		}
	}
}

func TestMergeWithin(t *testing.T) {
	inputsMArr := [][]int{
		{1, 2, 3, 0, 0, 0},
		{1},
		{0},
		{-8, -4, -3, 0, 0, 0}, // all negatives test case
		{1, 1, 0, 0},          // all similar numbers
		{0, 0, 0, 0},          // all zeros test case.
	}
	inputsM := []int{3, 1, 0, 3, 2, 2}

	inputsNArr := [][]int{
		{2, 5, 6},
		{},
		{1},
		{-7, -6, -5},
		{1, 1},
		{3, 4},
	}
	inputsN := []int{3, 0, 1, 3, 2, 2}

	outputs := [][]int{
		{1, 2, 2, 3, 5, 6},
		{1},
		{1},
		{-8, -7, -6, -5, -4, -3},
		{1, 1, 1, 1},
		{0, 0, 3, 4},
	}

	for i := range inputsMArr {
		leet.MergeWithin(inputsMArr[i], inputsM[i], inputsNArr[i], inputsN[i])

		if !CompareArrayValues(inputsMArr[i], outputs[i]) {
			t.Errorf("Wrong result for the input %v", i)
		}

	}
}

func CompareArrayValues(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestIntersect(t *testing.T) {
	nums1 := [][]int{
		{1, 2, 2, 1},
		{4, 9, 5},
		{10, 20, 30},
		{1},
		{1, 2, 3, 4, 5, 5},
		{5, 4, 5, 4},
		{3, 3, 3, 4, 4, 4, 5},
		{1},
		{4, 5, 3},
		{5},
	}

	nums2 := [][]int{
		{2, 2},
		{9, 4, 9, 8, 4},
		{40, 50},
		{9, 8, 7, 6, 5, 1},
		{1, 2, 3, 5, 4},
		{4, 5, 4, 5},
		{3, 4, 4, 4, 10},
		{1},
		{3},
		{3, 5, 4},
	}

	outputs := [][]int{
		{2, 2},
		{4, 9},
		{},
		{1},
		{1, 2, 3, 4, 5},
		{4, 4, 5, 5},
		{3, 4, 4, 4},
		{1},
		{3},
		{5},
	}

	for i := range nums1 {
		res := leet.Intersect(nums1[i], nums2[i])
		if !CompareArrayValues(res, outputs[i]) {
			t.Errorf("Expected %v, got %v", res, outputs[i])
		}
	}
}

func TestIntersectUnique(t *testing.T) {
	nums1 := [][]int{
		{1, 2, 2, 1},
		{4, 9, 5},
		{10, 20, 30},
		{1},
		{1, 2, 3, 4, 5, 5},
		{5, 4, 5, 4},
		{3, 3, 3, 4, 4, 4, 5},
		{1},
		{4, 5, 3},
		{5},
	}

	nums2 := [][]int{
		{2, 2},
		{9, 4, 9, 8, 4},
		{40, 50},
		{9, 8, 7, 6, 5, 1},
		{1, 2, 3, 5, 4},
		{4, 5, 4, 5},
		{3, 4, 4, 4, 10},
		{1},
		{3},
		{3, 5, 4},
	}

	outputs := [][]int{
		{2},
		{4, 9},
		{},
		{1},
		{1, 2, 3, 4, 5},
		{4, 5},
		{3, 4},
		{1},
		{3},
		{5},
	}

	for i := range nums1 {
		res := leet.IntersectionUnique(nums1[i], nums2[i])
		if !CompareArrayValues(res, outputs[i]) {
			t.Errorf("Expected %v, got %v", res, outputs[i])
		}
	}
}

func TestMatrixReshape(t *testing.T) {
	input := [][]int{
		{1, 2},
		{3, 4},
	}

	output := []int{
		1, 2, 3, 4,
	}

	res := leet.MatrixReshape(input, 1, 4)

	if !CompareArrayValues(res[0], output) {
		t.Errorf("Did not match the output")
	}
}

func TestMatrixReshape2(t *testing.T) {
	input := [][]int{
		{1, 2},
		{3, 4},
	}

	output := [][]int{
		{1, 2},
		{3, 4},
	}

	res := leet.MatrixReshape(input, 2, 2)

	if !Compare2dArrayValues(res, output) {
		t.Errorf("Did not match the output")
	}
}

func Compare2dArrayValues(x [][]int, y [][]int) bool {
	rx := len(x)
	cx := len(x[0])

	ry := len(y)
	cy := len(y[0])

	if rx != ry || cx != cy {
		return false
	}

	for i := 0; i < rx; i++ {
		for j := 0; j < cx; j++ {
			if x[i][j] != y[i][j] {
				return false
			}
		}
	}
	return true
}

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	res := leet.IsValidSudoku(board)

	if !res {
		t.Error("Expected the board to be a valid Sudoko. Returned it as invalid")
	}
}

func TestIsValidSudoku2(t *testing.T) {
	// first board has duplicates. So, it should return false (2 9s)
	board := [][]byte{
		{'9', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	res := leet.IsValidSudoku(board)

	if res {
		t.Error("Expected the board to be a invalid Sudoko. Returned it as Valid")
	}
}

func TestIsValidSudoku3(t *testing.T) {
	// second board has duplicates. So, it should return false (2 5s)
	board := [][]byte{
		{'1', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '5', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	res := leet.IsValidSudoku(board)

	if res {
		t.Error("Expected the board to be a invalid Sudoko. Returned it as Valid")
	}
}

func TestCanConstruct(t *testing.T) {
	ransomnotes := []string{
		"a",
		"aa",
		"aa",
		"xyzabc",
		"xyzzabc",
	}

	magazines := []string{
		"b",
		"ab",
		"aab",
		"aaaabbbccccdddxxyyz",
		"aaaabbbccccdddxxyyz",
	}

	results := []bool{
		false,
		false,
		true,
		true,
		false,
	}

	for i, _ := range ransomnotes {
		res := leet.CanConstruct(ransomnotes[i], magazines[i])
		if res != results[i] {
			t.Errorf("Did not return the expected value")
		}
	}

}

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}

	res := leet.Binarysearch(nums, 9)
	if res != 4 {
		t.Errorf("Expected the position to be 4. But got %v", res)
	}
}

func TestBinarySearch2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}

	res := leet.Binarysearch(nums, 2)
	if res != -1 {
		t.Errorf("Expected the position to be -1. But got %v", res)
	}
}

func TestBinarySearch3(t *testing.T) {
	nums := []int{2, 5}

	res := leet.Binarysearch(nums, 2)
	if res != 0 {
		t.Errorf("Expected the position to be 0. But got %v", res)
	}
}

func TestBinarySearch4(t *testing.T) {
	nums := []int{2, 5}

	res := leet.Binarysearch(nums, 5)
	if res != 1 {
		t.Errorf("Expected the position to be 5. But got %v", res)
	}
}

func TestSortedSquares(t *testing.T) {
	inputs := [][]int{
		{-4, -1, 0, 3, 10},
		{-7, -3, 2, 3, 11},
		{10},
		{-10},
		{-3, -2, -1},
		{1, 2, 3},
	}
	outputs := [][]int{
		{0, 1, 9, 16, 100},
		{4, 9, 9, 49, 121},
		{100},
		{100},
		{1, 4, 9},
		{1, 4, 9},
	}

	for i, input := range inputs {
		res := leet.SortedSquares(input)
		if !CompareArrayValues(res, outputs[i]) {
			t.Errorf("Expected %v, got %v", outputs[i], input)
		}
	}
}

func TestRotateSlice(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	kPos := []int{1, 2, 3, 4, 5}

	outputs := [][]int{
		{3, 1, 2},
		{2, 3, 1},
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}

	for i, input := range inputs {
		leet.Rotate(input, kPos[i])
		if !CompareArrayValues(input, outputs[i]) {
			t.Errorf("Expected %v, got %v for position %v", outputs[i], input, kPos[i])
		}
	}
}

func TestMoveZeros(t *testing.T) {
	inputs := [][]int{
		{0, 1, 0, 5, 0, 6},
		{0, 0, 0, 5, 6, 0},
		{0, 0, 5, 6, 7, 8},
	}
	outputs := [][]int{
		{1, 5, 6, 0, 0, 0},
		{5, 6, 0, 0, 0, 0},
		{5, 6, 7, 8, 0, 0},
	}

	for i, input := range inputs {
		leet.MoveZeroes(input)
		if !CompareArrayValues(input, outputs[i]) {
			t.Errorf("Expected %v, got %v ", outputs[i], input)
		}
	}
}

func TestTwoSumIncreasing(t *testing.T) {
	inputs := [][]int{
		{2, 3, 4},
		{-1, 0},
		{-1, 0, 1},
	}
	targets := []int{
		6,
		-1,
		0,
	}
	outputs := [][]int{
		{1, 3},
		{1, 2},
		{1, 3},
	}

	for i, input := range inputs {
		res := leet.TwoSumIncreasing(input, targets[i])
		if !CompareArrayValues(res, outputs[i]) {
			t.Errorf("Expected %v, got %v for the target %v", outputs[i], res, targets[i])
		}
	}

}

func TestLengthOfLongestSubstring(t *testing.T) {
	inputs := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"abccfghi",
		"abccfghih",
		"abc",
	}
	outputs := []int{3, 1, 3, 5, 5, 3}

	for i, input := range inputs {
		res := leet.LengthOfLongestSubstring(input)
		if res != outputs[i] {
			t.Errorf("Expected %v, got %v", outputs[i], res)
		}
	}

}

func TestCheckForInclusion(t *testing.T) {
	s1inputs := []string{
		"ab",
		"ab",
		"adc",
		"hello",
	}
	s2inputs := []string{
		"eidbaooo",
		"eidboaoo",
		"dcda",
		"ooolleoooleh",
	}
	outputs := []bool{
		true, false, true, false,
	}

	for i := range s1inputs {
		res := leet.CheckInclusion(s1inputs[i], s2inputs[i])
		if res != outputs[i] {
			t.Errorf("Expected %v, got %v", outputs[i], res)
		}
	}
}

func TestFloodFill(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	output := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}

	leet.FloodFill(image, 1, 1, 2)

	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[0]); j++ {
			if output[i][j] != image[i][j] {
				t.Errorf("Expected %v, got %v", output, image)
			}
		}
	}
}

func TestFloodFill2(t *testing.T) {
	image := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}

	output := [][]int{
		{2, 2, 2},
		{2, 2, 2},
	}

	leet.FloodFill(image, 0, 0, 2)

	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[0]); j++ {
			if output[i][j] != image[i][j] {
				t.Errorf("Expected %v, got %v", output, image)
			}
		}
	}
}

func TestFloodFill3(t *testing.T) {
	image := [][]int{
		{0, 0, 0},
		{0, 1, 1},
	}

	output := [][]int{
		{0, 0, 0},
		{0, 1, 1},
	}

	leet.FloodFill(image, 1, 1, 1)

	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[0]); j++ {
			if output[i][j] != image[i][j] {
				t.Errorf("Expected %v, got %v", output, image)
			}
		}
	}
}

func TestMaxAreaOfIsland(t *testing.T) {

	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	res := leet.MaxAreaOfIsland(grid)

	if res != 6 {
		t.Errorf("Expected 6 but got %v", res)
	}

}

func TestMaxAreaOfIsland2(t *testing.T) {

	grid := [][]int{
		{0, 1},
		{1, 1},
	}

	res := leet.MaxAreaOfIsland(grid)

	if res != 3 {
		t.Errorf("Expected 6 but got %v", res)
	}

}

func TestIslandPerimeter(t *testing.T) {
	input := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}

	res := leet.IslandPerimeter(input)
	if res != 16 {
		t.Errorf("Expected 16, got %v", res)
	}
}

func TestUpdateMatrix(t *testing.T) {
	input := [][]int{
		{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
	}

	output := [][]int{
		{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
	}

	res := leet.UpdateMatrix(input)

	if !Compare2dArrayValues(res, output) {
		t.Errorf("Expected %v, got %v", output, res)
	}
}

func TestUpdateMatrix2(t *testing.T) {
	input := [][]int{
		{0, 0, 0}, {0, 1, 0}, {1, 1, 1},
	}

	output := [][]int{
		{0, 0, 0}, {0, 1, 0}, {1, 2, 1},
	}

	res := leet.UpdateMatrix(input)

	if !Compare2dArrayValues(res, output) {
		t.Errorf("Expected %v, got %v", output, res)
	}
}

func TestUpdateMatrix3(t *testing.T) {
	input := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1},
		{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1},
		{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}}

	output := [][]int{
		{0, 0, 0}, {0, 1, 0}, {1, 2, 1},
	}

	res := leet.UpdateMatrix(input)

	if !Compare2dArrayValues(res, output) {
		t.Errorf("Expected %v, got %v", output, res)
	}
}

func TestOrangesRotting(t *testing.T) {
	input := [][]int{
		{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
	}

	res := leet.OrangesRotting(input)
	if res != 4 {
		t.Errorf("Expected 4, got %v", res)
	}

}

func TestOrangesRotting2(t *testing.T) {
	input := [][]int{
		{2, 1, 1}, {0, 1, 1}, {1, 0, 1},
	}

	res := leet.OrangesRotting(input)
	if res != -1 {
		t.Errorf("Expected -1, got %v", res)
	}

}

func TestOrangesRotting3(t *testing.T) {
	input := [][]int{
		{0, 2},
	}

	res := leet.OrangesRotting(input)
	if res != 0 {
		t.Errorf("Expected 0, got %v", res)
	}

}

func TestOrangesRotting4(t *testing.T) {
	input := [][]int{
		{2, 1, 1}, {1, 1, 1}, {0, 1, 2},
	}

	res := leet.OrangesRotting(input)
	if res != 2 {
		t.Errorf("Expected 2, got %v", res)
	}

}
