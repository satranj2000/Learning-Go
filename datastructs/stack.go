package datastructs

import "log"

type Stack struct {
	items []int
}

func (s *Stack) Push(d int) {
	s.items = append(s.items, d)
}

func (s *Stack) Pop() (val int, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Pop failed with error -  ", r)
			err = r.(error)
		}
	}()
	val = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, err
}
