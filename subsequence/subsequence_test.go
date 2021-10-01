package subsequence_test

import (
	"testing"

	"sathish.com/gomod/subsequence"
)

func TestFindIncreasingSubsequence(t *testing.T) {
	//case 1
	nulreturnval := subsequence.FindIncreasingSubsequence([]int{})
	if nulreturnval != nil {
		t.Errorf("Expected nil, but got something else")
	}

}
func TestFindIncreasingSubsequence2(t *testing.T) {
	//case 2
	v := 1
	singlevalue := subsequence.FindIncreasingSubsequence([]int{v})
	if singlevalue[0] != v {
		t.Errorf("Expected %d, but got %d", v, singlevalue[0])
	}
	if len(singlevalue) != 1 {
		t.Errorf("The size of return slice of an array of size should be 1")
	}
}

func TestGeneratePascalTriangle(t *testing.T) {
	subsequence.GeneratePascalTriangle(3)
	subsequence.GeneratePascalTriangle(5)
}

func TestGetRowPascalTriangle(t *testing.T) {
	inputstr := []int{3, 5, 1}
	outputstr := [][]int{
		{1, 2, 1}, {1, 4, 6, 4, 1}, {1},
	}

	for i, v := range inputstr {
		out := subsequence.GetRowPascalTriangle(v)
		for k := 0; k < len(out); k++ {
			if out[k] != outputstr[i][k] {
				t.Errorf("Expected %v, but got %v in the test number(%v)", out[k], outputstr[i][k], i)
			}
		}
	}
}

func TestCombine(t *testing.T) {
	outputs := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	res := subsequence.Combine(4, 2)

	for i := 0; i < len(outputs); i++ {
		for j := 0; j < len(outputs[0]); j++ {
			if res[i][j] != outputs[i][j] {
				t.Errorf("Expected %v, got %v", outputs, res)
			}
		}
	}
}

func TestCombine2(t *testing.T) {
	outputs := [][]int{
		{1, 2, 3},
		{1, 2, 4},
		{1, 3, 4},
		{2, 3, 4},
	}
	res := subsequence.Combine(4, 3)

	for i := 0; i < len(outputs); i++ {
		for j := 0; j < len(outputs[0]); j++ {
			if res[i][j] != outputs[i][j] {
				t.Errorf("Expected %v, got %v", outputs, res)
			}
		}
	}
}

func TestCombine3(t *testing.T) {
	outputs := [][]int{
		{1, 2, 3, 4},
	}
	res := subsequence.Combine(4, 4)

	for i := 0; i < len(outputs); i++ {
		for j := 0; j < len(outputs[0]); j++ {
			if res[i][j] != outputs[i][j] {
				t.Errorf("Expected %v, got %v", outputs, res)
			}
		}
	}
}

func TestCombine4(t *testing.T) {
	outputs := [][]int{
		{1},
	}
	res := subsequence.Combine(1, 1)

	for i := 0; i < len(outputs); i++ {
		for j := 0; j < len(outputs[0]); j++ {
			if res[i][j] != outputs[i][j] {
				t.Errorf("Expected %v, got %v", outputs, res)
			}
		}
	}
}

func TestCombine5(t *testing.T) {
	outputs := [][]int{
		{1}, {2},
	}
	res := subsequence.Combine(2, 1)

	if len(outputs) != len(res) {
		t.Errorf("Expected %v, got %v", outputs, res)
	}

	for i := 0; i < len(outputs); i++ {
		for j := 0; j < len(outputs[0]); j++ {
			if res[i][j] != outputs[i][j] {
				t.Errorf("Expected %v, got %v", outputs, res)
			}
		}
	}
}
