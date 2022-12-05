package main

type stack[T any] struct {
	s []T
}

func (s *stack[T]) push(i ...T) {
	s.s = append(s.s, i...)
}

func (s *stack[T]) pop() (top T) {
	top = s.top()
	s.s = s.s[:len(s.s)-1]
	return
}

func (s *stack[T]) popN(count int) (top []T) {
	top = s.s[len(s.s)-count:]
	s.s = s.s[:len(s.s)-count]
	return
}

func (s *stack[T]) top() (top T) {
	top = s.s[len(s.s)-1]
	return
}

func (s *stack[T]) size() int {
	return len(s.s)
}

func (s *stack[T]) pushBottom(i T) {
	s.s = append([]T{i}, s.s...)
}
