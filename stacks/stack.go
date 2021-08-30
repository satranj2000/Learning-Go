package stacks

type RuneStack []rune

func (s *RuneStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *RuneStack) StackLength() int {
	return len(*s)
}

func (s *RuneStack) Push(r rune) {
	*s = append(*s, r)
}

func (s *RuneStack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return rune(0), false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *RuneStack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return rune(0), false
	}
	index := len(*s) - 1
	element := (*s)[index]
	//*s = (*s)[:index]
	return element, true
}
