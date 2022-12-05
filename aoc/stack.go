package aoc

type Stack[T any] struct {
	s []T
}

func (s *Stack[T]) Push(i ...T) {
	s.s = append(s.s, i...)
}

func (s *Stack[T]) Pop() (top T) {
	top = s.Top()
	s.s = s.s[:len(s.s)-1]
	return
}

func (s *Stack[T]) PopN(count int) (top []T) {
	top = s.s[len(s.s)-count:]
	s.s = s.s[:len(s.s)-count]
	return
}

func (s *Stack[T]) Top() (top T) {
	top = s.s[len(s.s)-1]
	return
}

func (s *Stack[T]) Size() int {
	return len(s.s)
}

func (s *Stack[T]) PushBottom(i T) {
	s.s = append([]T{i}, s.s...)
}
