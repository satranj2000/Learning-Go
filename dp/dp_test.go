package dp

import "testing"

func TestMaxRob(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
		{1, 20, 0, 2, 3, 5, 19},
		{10},
	}

	outputs := []int{
		4, 12, 42, 10,
	}

	for i, input := range inputs {
		res := rob(input)
		if res != outputs[i] {
			t.Errorf("Expected %v, got %v", outputs[i], res)
		}
	}

}
