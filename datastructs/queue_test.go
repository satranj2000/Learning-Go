package datastructs

import (
	"testing"
)

func TestQueue(t *testing.T) {
	myQueue := Queue{}

	inputs := []int{10, 20, 30}
	outputs := []int{10, 20, 30}

	for i := 0; i < 3; i++ {
		myQueue.Enqueue(inputs[i])
	}

	for i := 0; i < 3; i++ {
		res, _ := myQueue.Dequeue()
		if res != outputs[i] {
			t.Errorf("Wrong value pop-ed. Expected %v, got %v", outputs[i], res)
		}
	}
}

func TestQueue2(t *testing.T) {
	myQueue := Queue{}

	inputs := []int{10}
	outputs := []int{10}

	myQueue.Enqueue(inputs[0])
	v, err := myQueue.Dequeue()
	if v != outputs[0] {
		t.Errorf("Wrong pop. Expected %v ", outputs[0])
	}

	_, err = myQueue.Dequeue()
	if err == nil {
		t.Errorf("Expected an error. did not recieved it")
	}
}
