package aoc

import "testing"

func TestMemoize(t *testing.T) {
	callCount := 0
	expensive := func(n int) int {
		callCount++
		return n * 2
	}

	memoized := Memoize(expensive)

	// First call should invoke function
	if memoized(5) != 10 {
		t.Error("Memoize: wrong result")
	}
	if callCount != 1 {
		t.Errorf("Memoize: expected 1 call, got %d", callCount)
	}

	// Second call with same arg should use cache
	if memoized(5) != 10 {
		t.Error("Memoize: wrong cached result")
	}
	if callCount != 1 {
		t.Errorf("Memoize: should use cache, got %d calls", callCount)
	}

	// Different arg should invoke function
	if memoized(3) != 6 {
		t.Error("Memoize: wrong result for new arg")
	}
	if callCount != 2 {
		t.Errorf("Memoize: expected 2 calls, got %d", callCount)
	}
}

func TestMemoize2(t *testing.T) {
	callCount := 0
	add := func(a, b int) int {
		callCount++
		return a + b
	}

	memoized := Memoize2(add)

	if memoized(2, 3) != 5 {
		t.Error("Memoize2: wrong result")
	}
	if callCount != 1 {
		t.Errorf("Memoize2: expected 1 call, got %d", callCount)
	}

	// Cached
	memoized(2, 3)
	if callCount != 1 {
		t.Error("Memoize2: should use cache")
	}

	// Different args
	memoized(3, 2)
	if callCount != 2 {
		t.Error("Memoize2: different args should call function")
	}
}

func TestMemoizeRecursive(t *testing.T) {
	// Classic fibonacci with memoization
	fib := MemoizeRecursive(func(recurse func(int) int, n int) int {
		if n <= 1 {
			return n
		}
		return recurse(n-1) + recurse(n-2)
	})

	if fib(10) != 55 {
		t.Errorf("MemoizeRecursive fib(10): expected 55, got %d", fib(10))
	}
	if fib(20) != 6765 {
		t.Errorf("MemoizeRecursive fib(20): expected 6765, got %d", fib(20))
	}
}

func TestMemoizeRecursiveLarge(t *testing.T) {
	// Without memoization, this would be very slow
	fib := MemoizeRecursive(func(recurse func(int) int, n int) int {
		if n <= 1 {
			return n
		}
		return recurse(n-1) + recurse(n-2)
	})

	// fib(40) = 102334155, would take forever without memoization
	result := fib(40)
	if result != 102334155 {
		t.Errorf("MemoizeRecursive fib(40): expected 102334155, got %d", result)
	}
}
