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
