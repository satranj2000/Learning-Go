package datastructs

import (
	"errors"
	"log"
)

type Queue struct {
	items       []int
	bInitialize bool
	maxsz       int
	currmax     int
	currmin     int
}

func (q *Queue) Enqueue(d int) error {
	if !q.bInitialize {
		q.bInitialize = true
		q.maxsz = 4096
		q.currmax = 0
		q.currmin = 0
		q.items = make([]int, q.maxsz)
	}
	if q.currmax > q.maxsz {
		log.Println("Cannot allocate more than 1024 values in the queue")
		return errors.New("cannot allocate more than 1024 values in the queue")
	}
	q.items[q.currmax] = d
	q.currmax++
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("empty queue")
	}
	val := q.items[q.currmin]
	q.currmin++
	return val, nil
}

func (q *Queue) IsEmpty() bool {
	return (q.currmax <= q.currmin)
}
