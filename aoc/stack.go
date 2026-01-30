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

// Queue is a FIFO queue with O(1) amortized enqueue and dequeue.
// Uses two stacks internally for efficient operations.
type Queue[T any] struct {
	inbox  []T // elements are pushed here
	outbox []T // elements are popped from here
}

func NewQueue[T any](in ...T) (q *Queue[T]) {
	q = &Queue[T]{}
	for _, v := range in {
		q.Enqueue(v)
	}
	return
}

// Enqueue adds an element to the back of the queue. O(1)
func (q *Queue[T]) Enqueue(i T) {
	q.inbox = append(q.inbox, i)
}

// Push is an alias for Enqueue for compatibility.
func (q *Queue[T]) Push(i ...T) {
	for _, v := range i {
		q.Enqueue(v)
	}
}

// Dequeue removes and returns the front element. O(1) amortized.
func (q *Queue[T]) Dequeue() T {
	q.refill()
	val := q.outbox[len(q.outbox)-1]
	q.outbox = q.outbox[:len(q.outbox)-1]
	return val
}

// Pop is an alias for Dequeue for compatibility.
func (q *Queue[T]) Pop() T {
	return q.Dequeue()
}

// Top returns the front element without removing it. O(1) amortized.
func (q *Queue[T]) Top() T {
	q.refill()
	return q.outbox[len(q.outbox)-1]
}

// refill moves elements from inbox to outbox when outbox is empty.
func (q *Queue[T]) refill() {
	if len(q.outbox) == 0 {
		for len(q.inbox) > 0 {
			n := len(q.inbox) - 1
			q.outbox = append(q.outbox, q.inbox[n])
			q.inbox = q.inbox[:n]
		}
	}
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int { return len(q.inbox) + len(q.outbox) }

// Empty returns true if the queue has no elements.
func (q *Queue[T]) Empty() bool { return q.Size() == 0 }
