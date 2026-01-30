package aoc

import "testing"

// =============================================================================
// Stack Tests
// =============================================================================

func TestNewStack(t *testing.T) {
	s := NewStack(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("NewStack: expected size 3, got %d", s.Size())
	}
}

func TestNewStackEmpty(t *testing.T) {
	s := NewStack[int]()
	if s.Size() != 0 {
		t.Errorf("NewStack empty: expected size 0, got %d", s.Size())
	}
	if !s.Empty() {
		t.Error("NewStack empty: should be empty")
	}
}

func TestStackPush(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2, 3)

	if s.Size() != 3 {
		t.Errorf("Push: expected size 3, got %d", s.Size())
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack(1, 2, 3)

	top := s.Pop()
	if top != 3 {
		t.Errorf("Pop: expected 3, got %d", top)
	}
	if s.Size() != 2 {
		t.Errorf("Pop: expected size 2, got %d", s.Size())
	}
}

func TestStackPopN(t *testing.T) {
	s := NewStack(1, 2, 3, 4, 5)

	popped := s.PopN(3)
	if len(popped) != 3 {
		t.Errorf("PopN: expected 3 elements, got %d", len(popped))
	}
	// Should get [3, 4, 5] (top 3)
	expected := []int{3, 4, 5}
	for i, v := range popped {
		if v != expected[i] {
			t.Errorf("PopN: expected %v, got %v", expected, popped)
			break
		}
	}
	if s.Size() != 2 {
		t.Errorf("PopN: expected size 2, got %d", s.Size())
	}
}

func TestStackTop(t *testing.T) {
	s := NewStack(1, 2, 3)

	top := s.Top()
	if top != 3 {
		t.Errorf("Top: expected 3, got %d", top)
	}
	// Top should not remove element
	if s.Size() != 3 {
		t.Errorf("Top: size should remain 3, got %d", s.Size())
	}
}

func TestStackPushBottom(t *testing.T) {
	s := NewStack(2, 3)
	s.PushBottom(1)

	// Stack should now be [1, 2, 3] with 3 on top
	if s.Size() != 3 {
		t.Errorf("PushBottom: expected size 3, got %d", s.Size())
	}
	if s.Top() != 3 {
		t.Errorf("PushBottom: top should still be 3, got %d", s.Top())
	}

	// Pop all to verify order
	if s.Pop() != 3 {
		t.Error("PushBottom: order wrong")
	}
	if s.Pop() != 2 {
		t.Error("PushBottom: order wrong")
	}
	if s.Pop() != 1 {
		t.Error("PushBottom: order wrong")
	}
}

func TestStackLIFO(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Should come out in reverse order (LIFO)
	if s.Pop() != 3 {
		t.Error("LIFO: first pop should be 3")
	}
	if s.Pop() != 2 {
		t.Error("LIFO: second pop should be 2")
	}
	if s.Pop() != 1 {
		t.Error("LIFO: third pop should be 1")
	}
}

func TestStackWithStrings(t *testing.T) {
	s := NewStack("a", "b", "c")
	if s.Pop() != "c" {
		t.Error("Stack with strings: should work")
	}
}

// =============================================================================
// Queue Tests
// =============================================================================

func TestNewQueue(t *testing.T) {
	q := NewQueue(1, 2, 3)
	if q.Size() != 3 {
		t.Errorf("NewQueue: expected size 3, got %d", q.Size())
	}
}

func TestNewQueueEmpty(t *testing.T) {
	q := NewQueue[int]()
	if q.Size() != 0 {
		t.Errorf("NewQueue empty: expected size 0, got %d", q.Size())
	}
	if !q.Empty() {
		t.Error("NewQueue empty: should be empty")
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Enqueue: expected size 3, got %d", q.Size())
	}
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue(1, 2, 3)

	val := q.Dequeue()
	if val != 1 {
		t.Errorf("Dequeue: expected 1, got %d", val)
	}
	if q.Size() != 2 {
		t.Errorf("Dequeue: expected size 2, got %d", q.Size())
	}
}

func TestQueueFIFOBehavior(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1) // Should be first out
	q.Enqueue(2)
	q.Enqueue(3) // Should be last out

	// FIFO: first in, first out
	first := q.Pop()
	if first != 1 {
		t.Errorf("FIFO: first should be 1, got %d", first)
	}

	second := q.Pop()
	if second != 2 {
		t.Errorf("FIFO: second should be 2, got %d", second)
	}

	third := q.Pop()
	if third != 3 {
		t.Errorf("FIFO: third should be 3, got %d", third)
	}
}

func TestQueueTop(t *testing.T) {
	q := NewQueue(1, 2, 3)

	// Top should return front element without removing it
	if q.Top() != 1 {
		t.Errorf("Top: expected 1, got %d", q.Top())
	}
	if q.Size() != 3 {
		t.Errorf("Top: should not change size, got %d", q.Size())
	}
}

func TestQueuePush(t *testing.T) {
	q := NewQueue[int]()
	q.Push(1, 2, 3)

	if q.Size() != 3 {
		t.Errorf("Push: expected size 3, got %d", q.Size())
	}

	// Should still be FIFO
	if q.Pop() != 1 {
		t.Error("Push: first element should be 1")
	}
}

// =============================================================================
// Queue Performance (now O(1) amortized)
// =============================================================================

func TestQueuePerformance(t *testing.T) {
	q := NewQueue[int]()

	// This should now be O(n) total instead of O(n²)
	for i := 0; i < 10000; i++ {
		q.Enqueue(i)
	}

	if q.Size() != 10000 {
		t.Errorf("Queue stress: expected 10000, got %d", q.Size())
	}

	// Verify FIFO order
	for i := 0; i < 10000; i++ {
		val := q.Dequeue()
		if val != i {
			t.Errorf("FIFO order: expected %d, got %d", i, val)
			break
		}
	}
}

// =============================================================================
// Stack/Queue Integration with Search
// =============================================================================

func TestStackForDFS(t *testing.T) {
	// Verify stack works correctly for DFS-style exploration
	s := NewStack[int]()

	// Simulate visiting nodes
	s.Push(1)       // Visit 1
	s.Push(2)       // Visit 2
	s.Push(3)       // Visit 3
	node := s.Pop() // Backtrack: should get 3
	if node != 3 {
		t.Errorf("DFS simulation: expected 3, got %d", node)
	}

	s.Push(4) // From 2, visit 4
	node = s.Pop()
	if node != 4 {
		t.Errorf("DFS simulation: expected 4, got %d", node)
	}

	node = s.Pop() // Backtrack to 2
	if node != 2 {
		t.Errorf("DFS simulation: expected 2, got %d", node)
	}
}

func TestQueueForBFS(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1) // First node to visit
	q.Enqueue(2) // Second
	q.Enqueue(3) // Third

	// FIFO: first enqueued is first dequeued
	first := q.Dequeue()
	if first != 1 {
		t.Errorf("BFS: first should be 1, got %d", first)
	}

	second := q.Dequeue()
	if second != 2 {
		t.Errorf("BFS: second should be 2, got %d", second)
	}

	third := q.Dequeue()
	if third != 3 {
		t.Errorf("BFS: third should be 3, got %d", third)
	}
}
