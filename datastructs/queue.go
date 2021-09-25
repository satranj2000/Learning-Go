package datastructs

import "log"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(d int) {
	q.items = append(q.items, d)
}

func (q *Queue) Dequeue() (val int, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Dequeue failed with error -  ", r)
			err = r.(error)
		}
	}()
	val = q.items[0]
	q.items = q.items[1:len(q.items)]
	return val, err
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) <= 0
}
