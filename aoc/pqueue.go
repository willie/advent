package aoc

import "container/heap"

// PriorityQueue is a min-heap based priority queue.
// Items with lower priority values are dequeued first.
type PriorityQueue[T any] struct {
	items *pqItems[T]
}

// PQItem represents an item in the priority queue.
type PQItem[T any] struct {
	Value    T
	Priority int
	index    int
}

// Internal heap implementation
type pqItems[T any] []*PQItem[T]

func (pq pqItems[T]) Len() int { return len(pq) }

func (pq pqItems[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq pqItems[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pqItems[T]) Push(x any) {
	n := len(*pq)
	item := x.(*PQItem[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *pqItems[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// NewPriorityQueue creates a new empty priority queue.
func NewPriorityQueue[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		items: &pqItems[T]{},
	}
	heap.Init(pq.items)
	return pq
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue[T]) Push(value T, priority int) {
	heap.Push(pq.items, &PQItem[T]{
		Value:    value,
		Priority: priority,
	})
}

// Pop removes and returns the item with the lowest priority.
func (pq *PriorityQueue[T]) Pop() (T, int) {
	item := heap.Pop(pq.items).(*PQItem[T])
	return item.Value, item.Priority
}

// Peek returns the item with the lowest priority without removing it.
func (pq *PriorityQueue[T]) Peek() (T, int) {
	if pq.Len() == 0 {
		var zero T
		return zero, 0
	}
	item := (*pq.items)[0]
	return item.Value, item.Priority
}

// Len returns the number of items in the queue.
func (pq *PriorityQueue[T]) Len() int {
	return pq.items.Len()
}

// Empty returns true if the queue is empty.
func (pq *PriorityQueue[T]) Empty() bool {
	return pq.Len() == 0
}
