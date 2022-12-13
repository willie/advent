package aoc

func BFS[T comparable](start T, goal T, neighbors func(current T) []T) []T {
	Q := NewQueue(start)
	visited := map[T]*T{start: nil}
	// visited[start] = nil

	for !Q.Empty() {
		current := Q.Pop()
		if current == goal {
			break
		}

		for _, n := range neighbors(current) {
			if _, ok := visited[n]; !ok {
				visited[n] = &current
				Q.PushBottom(n)
			}
		}
	}

	ret := []T{goal}
	for n := visited[goal]; n != nil; n = visited[*n] {
		ret = append(ret, *n)
	}

	return ret
}

/*

func BFS[T comparable](start T, goal T, neighbors func(current T) []T) []T {
	Q := queue.New[T]()
	Q.Enqueue(start)

	visited := make(map[T]*T)
	visited[start] = nil

	for !Q.Empty() {
		current := Q.Dequeue()
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

	ret := []T{goal}
	for n := visited[goal]; n != nil; n = visited[*n] {
		ret = append(ret, *n)
	}

	return ret
}

*/
