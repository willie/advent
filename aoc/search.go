package aoc

// BFS performs breadth-first search from start to goal.
// Returns the shortest path (reversed, goal first) or nil if no path exists.
// For weighted graphs, use Dijkstra instead.
func BFS[T comparable](start T, goal T, neighbors func(current T) []T) []T {
	if start == goal {
		return []T{goal}
	}

	Q := NewQueue(start)
	visited := map[T]*T{start: nil}

	for !Q.Empty() {
		current := Q.Pop()
		if current == goal {
			break
		}

		for _, n := range neighbors(current) {
			if _, ok := visited[n]; !ok {
				visited[n] = &current
				Q.Enqueue(n)
			}
		}
	}

	// Check if goal was reached
	if _, ok := visited[goal]; !ok {
		return nil
	}

	ret := []T{goal}
	for n := visited[goal]; n != nil; n = visited[*n] {
		ret = append(ret, *n)
	}

	return ret
}

// DFS performs depth-first search from start to goal.
// Returns the path (reversed, goal first) or nil if no path exists.
// Note: For weighted graphs, use Dijkstra instead.
func DFS[T comparable](start T, goal T, neighbors func(current T) []T) []T {
	if start == goal {
		return []T{goal}
	}

	S := NewStack(start)
	visited := map[T]*T{start: nil}

	for !S.Empty() {
		current := S.Pop()
		if current == goal {
			break
		}

		for _, n := range neighbors(current) {
			if _, ok := visited[n]; !ok {
				visited[n] = &current
				S.Push(n)
			}
		}
	}

	// Check if goal was reached
	if _, ok := visited[goal]; !ok {
		return nil
	}

	ret := []T{goal}
	for n := visited[goal]; n != nil; n = visited[*n] {
		ret = append(ret, *n)
	}

	return ret
}
