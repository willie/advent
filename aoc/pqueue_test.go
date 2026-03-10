package aoc

import "testing"

func TestPriorityQueueBasic(t *testing.T) {
	pq := NewPriorityQueue[string]()

	pq.Push("low", 1)
	pq.Push("high", 10)
	pq.Push("medium", 5)

	// Should come out in priority order (lowest first)
	val, pri := pq.Pop()
	if val != "low" || pri != 1 {
		t.Errorf("PQ first: expected (low, 1), got (%s, %d)", val, pri)
	}

	val, pri = pq.Pop()
	if val != "medium" || pri != 5 {
		t.Errorf("PQ second: expected (medium, 5), got (%s, %d)", val, pri)
	}

	val, pri = pq.Pop()
	if val != "high" || pri != 10 {
		t.Errorf("PQ third: expected (high, 10), got (%s, %d)", val, pri)
	}
}

func TestPriorityQueueEmpty(t *testing.T) {
	pq := NewPriorityQueue[int]()

	if !pq.Empty() {
		t.Error("New PQ should be empty")
	}
	if pq.Len() != 0 {
		t.Error("New PQ should have length 0")
	}

	pq.Push(1, 1)
	if pq.Empty() {
		t.Error("PQ with item should not be empty")
	}

	pq.Pop()
	if !pq.Empty() {
		t.Error("PQ after pop should be empty")
	}
}

func TestPriorityQueuePeek(t *testing.T) {
	pq := NewPriorityQueue[string]()
	pq.Push("first", 1)
	pq.Push("second", 2)

	val, pri := pq.Peek()
	if val != "first" || pri != 1 {
		t.Errorf("Peek: expected (first, 1), got (%s, %d)", val, pri)
	}

	// Peek should not remove
	if pq.Len() != 2 {
		t.Error("Peek should not remove item")
	}
}

func TestPriorityQueueSamePriority(t *testing.T) {
	pq := NewPriorityQueue[string]()
	pq.Push("a", 1)
	pq.Push("b", 1)
	pq.Push("c", 1)

	// All have same priority, should all come out
	count := 0
	for !pq.Empty() {
		pq.Pop()
		count++
	}
	if count != 3 {
		t.Errorf("Same priority: expected 3 items, got %d", count)
	}
}

func TestPriorityQueueWithInts(t *testing.T) {
	pq := NewPriorityQueue[int]()

	// Push in random order
	values := []int{5, 3, 8, 1, 9, 2}
	for _, v := range values {
		pq.Push(v, v) // Priority = value
	}

	// Should come out sorted
	expected := []int{1, 2, 3, 5, 8, 9}
	for _, exp := range expected {
		val, _ := pq.Pop()
		if val != exp {
			t.Errorf("PQ sorted: expected %d, got %d", exp, val)
		}
	}
}

func TestPriorityQueueLarge(t *testing.T) {
	pq := NewPriorityQueue[int]()

	// Push 1000 items in reverse order
	for i := 1000; i > 0; i-- {
		pq.Push(i, i)
	}

	// Should come out in order 1, 2, 3, ...
	for i := 1; i <= 1000; i++ {
		val, _ := pq.Pop()
		if val != i {
			t.Errorf("Large PQ: expected %d, got %d", i, val)
			break
		}
	}
}

func TestPriorityQueueNegativePriority(t *testing.T) {
	pq := NewPriorityQueue[string]()
	pq.Push("negative", -5)
	pq.Push("zero", 0)
	pq.Push("positive", 5)

	val, pri := pq.Pop()
	if val != "negative" || pri != -5 {
		t.Errorf("Negative priority: expected (negative, -5), got (%s, %d)", val, pri)
	}
}

func TestPriorityQueueForDijkstra(t *testing.T) {
	// Simulate Dijkstra usage pattern
	type State struct {
		node int
		dist int
	}

	pq := NewPriorityQueue[State]()

	// Add initial states
	pq.Push(State{1, 0}, 0)
	pq.Push(State{2, 5}, 5)
	pq.Push(State{3, 3}, 3)

	// Process in order
	s, _ := pq.Pop()
	if s.node != 1 {
		t.Errorf("Dijkstra pattern: expected node 1 first, got %d", s.node)
	}

	// Add more states (discovered neighbors)
	pq.Push(State{4, 2}, 2)

	s, _ = pq.Pop()
	if s.node != 4 {
		t.Errorf("Dijkstra pattern: expected node 4 second, got %d", s.node)
	}
}
