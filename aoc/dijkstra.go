package aoc

// Edge represents a weighted edge to a neighbor node.
type Edge[T any] struct {
	To   T
	Cost int
}

// Dijkstra finds the shortest path from start to goal in a weighted graph.
// Returns the path (reversed, goal first) and total cost.
// If no path exists, returns nil and -1.
func Dijkstra[T comparable](start, goal T, neighbors func(T) []Edge[T]) ([]T, int) {
	pq := NewPriorityQueue[T]()
	pq.Push(start, 0)

	dist := map[T]int{start: 0}
	prev := map[T]*T{start: nil}

	for !pq.Empty() {
		current, currentDist := pq.Pop()

		// Skip if we've found a better path already
		if d, ok := dist[current]; ok && currentDist > d {
			continue
		}

		if current == goal {
			// Reconstruct path
			path := []T{goal}
			for p := prev[goal]; p != nil; p = prev[*p] {
				path = append(path, *p)
			}
			return path, currentDist
		}

		for _, edge := range neighbors(current) {
			newDist := currentDist + edge.Cost
			if oldDist, ok := dist[edge.To]; !ok || newDist < oldDist {
				dist[edge.To] = newDist
				prev[edge.To] = &current
				pq.Push(edge.To, newDist)
			}
		}
	}

	return nil, -1
}

// DijkstraAll finds shortest distances from start to all reachable nodes.
// Returns a map of node to minimum distance.
func DijkstraAll[T comparable](start T, neighbors func(T) []Edge[T]) map[T]int {
	pq := NewPriorityQueue[T]()
	pq.Push(start, 0)

	dist := map[T]int{start: 0}

	for !pq.Empty() {
		current, currentDist := pq.Pop()

		if d, ok := dist[current]; ok && currentDist > d {
			continue
		}

		for _, edge := range neighbors(current) {
			newDist := currentDist + edge.Cost
			if oldDist, ok := dist[edge.To]; !ok || newDist < oldDist {
				dist[edge.To] = newDist
				pq.Push(edge.To, newDist)
			}
		}
	}

	return dist
}

// DijkstraFunc finds shortest path using a goal function instead of a specific goal.
// Useful when there are multiple valid goals.
func DijkstraFunc[T comparable](start T, isGoal func(T) bool, neighbors func(T) []Edge[T]) ([]T, int) {
	pq := NewPriorityQueue[T]()
	pq.Push(start, 0)

	dist := map[T]int{start: 0}
	prev := map[T]*T{start: nil}

	for !pq.Empty() {
		current, currentDist := pq.Pop()

		if d, ok := dist[current]; ok && currentDist > d {
			continue
		}

		if isGoal(current) {
			path := []T{current}
			for p := prev[current]; p != nil; p = prev[*p] {
				path = append(path, *p)
			}
			return path, currentDist
		}

		for _, edge := range neighbors(current) {
			newDist := currentDist + edge.Cost
			if oldDist, ok := dist[edge.To]; !ok || newDist < oldDist {
				dist[edge.To] = newDist
				prev[edge.To] = &current
				pq.Push(edge.To, newDist)
			}
		}
	}

	return nil, -1
}
