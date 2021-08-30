package datastructs

import (
	"testing"
)

func TestStack(t *testing.T) {
	myStack := Stack{}

	inputs := []int{10, 20, 30}
	outputs := []int{30, 20, 10}

	for i := 0; i < 3; i++ {
		myStack.Push(inputs[i])
	}

	for i := 0; i < 3; i++ {
		res, _ := myStack.Pop()
		if res != outputs[i] {
			t.Errorf("Wrong value pop-ed. Expected %v, got %v", outputs[i], res)
		}
	}
}

func TestStack2(t *testing.T) {
	myStack := Stack{}

	inputs := []int{10}
	outputs := []int{10}

	myStack.Push(inputs[0])
	v, _ := myStack.Pop()
	if v != outputs[0] {
		t.Errorf("Wrong pop. Expected %v ", outputs[0])
	}

	_, err := myStack.Pop()
	if err == nil {
		t.Errorf("Excepted an error. But got nil ")
	}
}
