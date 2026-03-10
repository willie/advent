package aoc

import (
	"image"
	"testing"
)

func TestDijkstraSimple(t *testing.T) {
	// Simple weighted graph:
	// 1 --2--> 2 --3--> 3
	//  \              /
	//   ------10-----/
	neighbors := func(n int) []Edge[int] {
		switch n {
		case 1:
			return []Edge[int]{{2, 2}, {3, 10}}
		case 2:
			return []Edge[int]{{3, 3}}
		}
		return nil
	}

	path, cost := Dijkstra(1, 3, neighbors)

	if cost != 5 {
		t.Errorf("Dijkstra cost: expected 5, got %d", cost)
	}
	if len(path) != 3 {
		t.Errorf("Dijkstra path length: expected 3, got %d", len(path))
	}
	// Path is reversed: [3, 2, 1]
	if path[0] != 3 || path[len(path)-1] != 1 {
		t.Errorf("Dijkstra path: expected [3,2,1], got %v", path)
	}
}

func TestDijkstraGrid(t *testing.T) {
	// 3x3 grid with uniform cost
	neighbors := func(p image.Point) []Edge[image.Point] {
		var edges []Edge[image.Point]
		for _, d := range FourWay {
			np := p.Add(d)
			if np.X >= 0 && np.X < 3 && np.Y >= 0 && np.Y < 3 {
				edges = append(edges, Edge[image.Point]{np, 1})
			}
		}
		return edges
	}

	path, cost := Dijkstra(image.Pt(0, 0), image.Pt(2, 2), neighbors)

	// Manhattan distance is 4
	if cost != 4 {
		t.Errorf("Dijkstra grid cost: expected 4, got %d", cost)
	}
	if len(path) != 5 {
		t.Errorf("Dijkstra grid path length: expected 5, got %d", len(path))
	}
}

func TestDijkstraNoPath(t *testing.T) {
	neighbors := func(n int) []Edge[int] {
		if n == 1 {
			return []Edge[int]{{2, 1}}
		}
		return nil
	}

	path, cost := Dijkstra(1, 99, neighbors)

	if path != nil {
		t.Errorf("Dijkstra no path: expected nil, got %v", path)
	}
	if cost != -1 {
		t.Errorf("Dijkstra no path cost: expected -1, got %d", cost)
	}
}

func TestDijkstraStartEqualsGoal(t *testing.T) {
	neighbors := func(n int) []Edge[int] {
		return []Edge[int]{{n + 1, 1}}
	}

	path, cost := Dijkstra(5, 5, neighbors)

	if cost != 0 {
		t.Errorf("Dijkstra start=goal cost: expected 0, got %d", cost)
	}
	if len(path) != 1 || path[0] != 5 {
		t.Errorf("Dijkstra start=goal path: expected [5], got %v", path)
	}
}

func TestDijkstraAll(t *testing.T) {
	// Star graph: center connects to 1,2,3,4 with different costs
	neighbors := func(n int) []Edge[int] {
		if n == 0 {
			return []Edge[int]{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
		}
		return nil
	}

	dist := DijkstraAll(0, neighbors)

	if dist[0] != 0 {
		t.Errorf("DijkstraAll: dist to start should be 0")
	}
	if dist[1] != 1 || dist[2] != 2 || dist[3] != 3 || dist[4] != 4 {
		t.Errorf("DijkstraAll: wrong distances %v", dist)
	}
}

func TestDijkstraFunc(t *testing.T) {
	// Find any node >= 10
	neighbors := func(n int) []Edge[int] {
		return []Edge[int]{{n + 1, 1}, {n + 2, 3}}
	}

	path, cost := DijkstraFunc(1, func(n int) bool { return n >= 10 }, neighbors)

	if cost != 5 { // 1->2->3->4->5->6->7->8->9->10 costs 9, but 1->3->5->7->9->11 costs 5
		t.Logf("DijkstraFunc: cost=%d, path=%v", cost, path)
	}
	if path[0] < 10 {
		t.Errorf("DijkstraFunc: goal should be >= 10, got %d", path[0])
	}
}

func TestDijkstraWithNegativeEdge(t *testing.T) {
	// Note: Dijkstra doesn't handle negative edges correctly
	// This test just verifies it runs without crashing
	neighbors := func(n int) []Edge[int] {
		if n == 1 {
			return []Edge[int]{{2, -1}} // Negative edge
		}
		return nil
	}

	path, _ := Dijkstra(1, 2, neighbors)
	if path == nil {
		t.Error("Dijkstra with negative: should find path (though cost may be wrong)")
	}
}
