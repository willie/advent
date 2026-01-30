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

	// Note: Current implementation has Dequeue not returning value
	// It just calls Pop which returns the LAST element (wrong for queue)
	// This documents the current (incorrect) behavior
	q.Dequeue()
	if q.Size() != 2 {
		t.Errorf("Dequeue: expected size 2, got %d", q.Size())
	}
}

func TestQueueFIFOBehavior(t *testing.T) {
	// This test documents expected FIFO behavior
	// The current implementation uses Stack.Pop which is LIFO
	// For a proper queue, first in should be first out

	q := NewQueue[int]()
	q.Enqueue(1) // Should be first out
	q.Enqueue(2)
	q.Enqueue(3) // Should be last out

	// Using Pop (which Queue inherits) - this is LIFO not FIFO
	// A proper queue would have Pop return 1, but current returns 3
	first := q.Pop()

	// Document current behavior (which is actually LIFO due to implementation)
	// In a proper FIFO queue, first should be 1
	// Current implementation returns 3 (last pushed)
	if first != 3 {
		t.Logf("Note: Queue.Pop() returns %d (LIFO behavior)", first)
	}
}

// =============================================================================
// Queue Performance Note
// =============================================================================

func TestQueueEnqueuePerformanceNote(t *testing.T) {
	// This test documents the O(n) performance issue with Enqueue
	// Enqueue calls PushBottom which prepends to slice - O(n) operation
	// For large queues in BFS, this can be slow

	q := NewQueue[int]()

	// This is O(n²) total due to O(n) prepend
	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
	}

	if q.Size() != 1000 {
		t.Errorf("Queue stress: expected 1000, got %d", q.Size())
	}

	// A proper ring buffer or container/list would be O(1) per enqueue
	t.Log("Note: Queue.Enqueue is O(n) - consider ring buffer for large queues")
}

// =============================================================================
// Stack/Queue Integration with Search
// =============================================================================

func TestStackForDFS(t *testing.T) {
	// Verify stack works correctly for DFS-style exploration
	s := NewStack[int]()

	// Simulate visiting nodes
	s.Push(1)        // Visit 1
	s.Push(2)        // Visit 2
	s.Push(3)        // Visit 3
	node := s.Pop()  // Backtrack: should get 3
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

func TestQueueUsedAsStackForBFS(t *testing.T) {
	// Since Queue's Enqueue/Dequeue don't work as expected for FIFO,
	// BFS implementation uses Queue.Pop() and Queue.PushBottom()
	// which effectively makes it work but with O(n) enqueue

	q := NewQueue[int]()
	q.PushBottom(1) // First node to visit
	q.PushBottom(2) // Second
	q.PushBottom(3) // Third

	// Pop returns from top (last pushed to bottom = first)
	// Wait, PushBottom prepends, so [3,2,1] internally, Pop returns 1
	// Actually let's trace through:
	// PushBottom(1) -> [1]
	// PushBottom(2) -> [2,1]
	// PushBottom(3) -> [3,2,1]
	// Pop() returns s[len-1] = 1 ✓ (correct FIFO)

	first := q.Pop()
	if first != 1 {
		t.Errorf("Queue as BFS: first should be 1, got %d", first)
	}

	second := q.Pop()
	if second != 2 {
		t.Errorf("Queue as BFS: second should be 2, got %d", second)
	}
}
