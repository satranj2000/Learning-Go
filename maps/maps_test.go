package maps

import "testing"

func TestContainsDuplicate(t *testing.T) {

	inputsList := [][]int{
		{1, 2, 3},
		{1, 2, 1, 3, 4},
		{1},
		{1, 2, 3, 4, 5, 6, 7, 8, 2},
	}

	outputList := []bool{false, true, false, true}

	for i, input := range inputsList {
		if ContainsDuplicate(input) != outputList[i] {
			t.Errorf("The Contains Duplicate function failed for input %v. Expected %v, Got %v", i, outputList[i], !outputList[i])
		}
	}
}

func TestMaxSubArray(t *testing.T) {

	inputsList := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
	}

	outputList := []int{6, 1, 23}

	for i, v := range inputsList {
		res := maxSubArray(v)
		if res != outputList[i] {
			t.Errorf("Expected %v, got %v", outputList[i], res)
		}
	}
}
