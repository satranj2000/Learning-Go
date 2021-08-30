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
