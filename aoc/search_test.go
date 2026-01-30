package aoc

import (
	"image"
	"testing"
)

// =============================================================================
// BFS Tests
// =============================================================================

func TestBFSSimplePath(t *testing.T) {
	// Simple linear graph: 1 -> 2 -> 3 -> 4 -> 5
	neighbors := func(n int) []int {
		if n < 5 {
			return []int{n + 1}
		}
		return nil
	}

	path := BFS(1, 5, neighbors)

	// Path should be [5, 4, 3, 2, 1] (reversed)
	if len(path) != 5 {
		t.Errorf("BFS simple: expected path length 5, got %d", len(path))
	}
	if path[0] != 5 {
		t.Errorf("BFS simple: path should start with goal 5, got %d", path[0])
	}
	if path[len(path)-1] != 1 {
		t.Errorf("BFS simple: path should end with start 1, got %d", path[len(path)-1])
	}
}

func TestBFSGrid(t *testing.T) {
	// 3x3 grid, find path from (0,0) to (2,2)
	neighbors := func(p image.Point) []image.Point {
		var result []image.Point
		for _, d := range []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			np := p.Add(d)
			if np.X >= 0 && np.X < 3 && np.Y >= 0 && np.Y < 3 {
				result = append(result, np)
			}
		}
		return result
	}

	start := image.Pt(0, 0)
	goal := image.Pt(2, 2)
	path := BFS(start, goal, neighbors)

	// Shortest path is 5 steps: (0,0) -> (1,0) -> (2,0) -> (2,1) -> (2,2)
	// Or any other path of length 5
	if len(path) != 5 {
		t.Errorf("BFS grid: expected path length 5, got %d", len(path))
	}
	if path[0] != goal {
		t.Errorf("BFS grid: path should start with goal, got %v", path[0])
	}
	if path[len(path)-1] != start {
		t.Errorf("BFS grid: path should end with start, got %v", path[len(path)-1])
	}
}

func TestBFSStartEqualsGoal(t *testing.T) {
	neighbors := func(n int) []int { return []int{n + 1} }

	path := BFS(5, 5, neighbors)

	// When start == goal, path should just be [goal]
	if len(path) != 1 {
		t.Errorf("BFS start=goal: expected path length 1, got %d", len(path))
	}
	if path[0] != 5 {
		t.Errorf("BFS start=goal: expected [5], got %v", path)
	}
}

func TestBFSNoPath(t *testing.T) {
	// Disconnected graph: 1-2-3 and 4-5 (no connection)
	neighbors := func(n int) []int {
		switch n {
		case 1:
			return []int{2}
		case 2:
			return []int{1, 3}
		case 3:
			return []int{2}
		case 4:
			return []int{5}
		case 5:
			return []int{4}
		}
		return nil
	}

	path := BFS(1, 5, neighbors)

	// No path exists - current implementation returns [goal] even if unreachable
	// This documents the current (incorrect) behavior
	if path[0] != 5 {
		t.Errorf("BFS no path: expected [5] (current behavior), got %v", path)
	}

	// Note: A proper implementation should return empty path or error
	t.Log("Note: BFS returns [goal] even when no path exists - should return empty")
}

func TestBFSWithCycles(t *testing.T) {
	// Graph with cycle: 1 -> 2 -> 3 -> 1, with 3 -> 4
	neighbors := func(n int) []int {
		switch n {
		case 1:
			return []int{2}
		case 2:
			return []int{3}
		case 3:
			return []int{1, 4}
		case 4:
			return nil
		}
		return nil
	}

	path := BFS(1, 4, neighbors)

	// Should find path 1 -> 2 -> 3 -> 4
	if len(path) != 4 {
		t.Errorf("BFS with cycle: expected path length 4, got %d", len(path))
	}
}

// =============================================================================
// DFS Tests
// =============================================================================

func TestDFSSimplePath(t *testing.T) {
	// Simple linear graph: 1 -> 2 -> 3 -> 4 -> 5
	neighbors := func(n int) []int {
		if n < 5 {
			return []int{n + 1}
		}
		return nil
	}

	path := DFS(1, 5, neighbors)

	if len(path) != 5 {
		t.Errorf("DFS simple: expected path length 5, got %d", len(path))
	}
	if path[0] != 5 {
		t.Errorf("DFS simple: path should start with goal 5, got %d", path[0])
	}
}

func TestDFSBranching(t *testing.T) {
	// Binary tree structure:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	neighbors := func(n int) []int {
		switch n {
		case 1:
			return []int{2, 3}
		case 2:
			return []int{4, 5}
		}
		return nil
	}

	path := DFS(1, 5, neighbors)

	// Should find path 1 -> 2 -> 5
	if len(path) != 3 {
		t.Errorf("DFS branching: expected path length 3, got %d: %v", len(path), path)
	}
}

func TestDFSStartEqualsGoal(t *testing.T) {
	neighbors := func(n int) []int { return []int{n + 1} }

	path := DFS(5, 5, neighbors)

	if len(path) != 1 {
		t.Errorf("DFS start=goal: expected path length 1, got %d", len(path))
	}
}

func TestDFSNoPath(t *testing.T) {
	// Disconnected nodes
	neighbors := func(n int) []int {
		if n == 1 {
			return []int{2}
		}
		return nil
	}

	path := DFS(1, 99, neighbors)

	// Current implementation returns [goal] even when unreachable
	if path[0] != 99 {
		t.Errorf("DFS no path: expected [99] (current behavior), got %v", path)
	}

	t.Log("Note: DFS returns [goal] even when no path exists - should return empty")
}

// =============================================================================
// BFS vs DFS Comparison
// =============================================================================

func TestBFSFindsShortest(t *testing.T) {
	// Diamond graph:
	//     1
	//    / \
	//   2   3
	//    \ /
	//     4
	// BFS should find shortest path (length 2), DFS might find longer

	neighbors := func(n int) []int {
		switch n {
		case 1:
			return []int{2, 3}
		case 2:
			return []int{4}
		case 3:
			return []int{4}
		}
		return nil
	}

	bfsPath := BFS(1, 4, neighbors)
	dfsPath := DFS(1, 4, neighbors)

	// Both should find a path
	if len(bfsPath) == 0 || len(dfsPath) == 0 {
		t.Error("Both BFS and DFS should find a path")
	}

	// BFS should find optimal (length 3: 1->2->4 or 1->3->4)
	if len(bfsPath) != 3 {
		t.Errorf("BFS should find shortest path of length 3, got %d", len(bfsPath))
	}
}

// =============================================================================
// Search with Complex State
// =============================================================================

type State struct {
	Pos   image.Point
	Keys  int // Bitmask of collected keys
}

func TestBFSComplexState(t *testing.T) {
	// Simulate a maze with keys
	// State includes position AND collected keys

	neighbors := func(s State) []State {
		var result []State
		// Can move in 4 directions
		for _, d := range []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			np := s.Pos.Add(d)
			if np.X >= 0 && np.X < 3 && np.Y >= 0 && np.Y < 3 {
				newState := State{Pos: np, Keys: s.Keys}
				// Collect key at (1,1)
				if np.X == 1 && np.Y == 1 {
					newState.Keys |= 1
				}
				result = append(result, newState)
			}
		}
		return result
	}

	start := State{Pos: image.Pt(0, 0), Keys: 0}
	goal := State{Pos: image.Pt(2, 2), Keys: 1} // Need to collect key first

	path := BFS(start, goal, neighbors)

	// Path should go through (1,1) to collect key, then to goal
	if len(path) == 0 {
		t.Error("BFS complex: should find path")
	}
	if path[0] != goal {
		t.Errorf("BFS complex: path should end at goal, got %v", path[0])
	}
}

// =============================================================================
// Performance Characteristics
// =============================================================================

func TestSearchLargeGraph(t *testing.T) {
	// Test with larger graph to verify it completes
	// 100 nodes in a line
	neighbors := func(n int) []int {
		if n < 100 {
			return []int{n + 1}
		}
		return nil
	}

	bfsPath := BFS(1, 100, neighbors)
	if len(bfsPath) != 100 {
		t.Errorf("BFS large: expected path length 100, got %d", len(bfsPath))
	}

	dfsPath := DFS(1, 100, neighbors)
	if len(dfsPath) != 100 {
		t.Errorf("DFS large: expected path length 100, got %d", len(dfsPath))
	}
}
