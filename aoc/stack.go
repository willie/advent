package aoc

type Stack[T any] struct {
	s []T
}

func NewStack[T any](in ...T) (q *Stack[T]) {
	q = &Stack[T]{}
	q.Push(in...)
	return
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

func (s *Stack[T]) Size() int      { return len(s.s) }
func (s *Stack[T]) Empty() bool    { return len(s.s) == 0 }
func (s *Stack[T]) PushBottom(i T) { s.s = append([]T{i}, s.s...) }

type Queue[T any] struct {
	Stack[T]
}

func NewQueue[T any](in ...T) (q *Queue[T]) {
	q = &Queue[T]{}
	q.Push(in...)
	return
}

func (s *Queue[T]) Dequeue()    { s.Pop() }
func (s *Queue[T]) Enqueue(i T) { s.PushBottom(i) }
