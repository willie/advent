package aoc

import "testing"

// =============================================================================
// Node Creation Tests
// =============================================================================

func TestNewNode(t *testing.T) {
	node := NewNode(42)

	if node.Data != 42 {
		t.Errorf("NewNode Data: expected 42, got %d", node.Data)
	}
	if len(node.Children) != 0 {
		t.Errorf("NewNode Children: expected 0, got %d", len(node.Children))
	}
	if node.Parent != nil {
		t.Error("NewNode Parent: should be nil")
	}
}

func TestNewNodeWithString(t *testing.T) {
	node := NewNode("root")

	if node.Data != "root" {
		t.Errorf("NewNode string: expected 'root', got '%s'", node.Data)
	}
}

// =============================================================================
// AddNode Tests
// =============================================================================

func TestAddNode(t *testing.T) {
	root := NewNode("root")
	child := root.AddNode("child")

	if len(root.Children) != 1 {
		t.Errorf("AddNode: root should have 1 child, got %d", len(root.Children))
	}
	if child.Data != "child" {
		t.Errorf("AddNode: child data should be 'child', got '%s'", child.Data)
	}
	if child.Parent != root {
		t.Error("AddNode: child's parent should be root")
	}
}

func TestAddMultipleChildren(t *testing.T) {
	root := NewNode("root")
	child1 := root.AddNode("child1")
	child2 := root.AddNode("child2")
	child3 := root.AddNode("child3")

	if len(root.Children) != 3 {
		t.Errorf("AddNode multiple: root should have 3 children, got %d", len(root.Children))
	}
	if root.Children[0] != child1 || root.Children[1] != child2 || root.Children[2] != child3 {
		t.Error("AddNode multiple: children order wrong")
	}
}

func TestAddNestedChildren(t *testing.T) {
	root := NewNode("root")
	child := root.AddNode("child")
	grandchild := child.AddNode("grandchild")

	if grandchild.Parent != child {
		t.Error("Nested: grandchild's parent should be child")
	}
	if child.Parent != root {
		t.Error("Nested: child's parent should be root")
	}
	if len(child.Children) != 1 {
		t.Errorf("Nested: child should have 1 child, got %d", len(child.Children))
	}
}

// =============================================================================
// PreOrder Traversal Tests
// =============================================================================

func TestTraversePreOrder(t *testing.T) {
	//     1
	//    / \
	//   2   3
	//  /
	// 4
	root := NewNode(1)
	child2 := root.AddNode(2)
	root.AddNode(3)
	child2.AddNode(4)

	var visited []int
	root.TraversePreOrder(func(n *Node[int]) {
		visited = append(visited, n.Data)
	})

	// PreOrder: root, then children left-to-right
	expected := []int{1, 2, 4, 3}
	if len(visited) != len(expected) {
		t.Errorf("PreOrder: expected %d nodes, got %d", len(expected), len(visited))
	}
	for i, v := range visited {
		if v != expected[i] {
			t.Errorf("PreOrder: expected %v, got %v", expected, visited)
			break
		}
	}
}

func TestTraversePreOrderSingle(t *testing.T) {
	root := NewNode(42)

	var visited []int
	root.TraversePreOrder(func(n *Node[int]) {
		visited = append(visited, n.Data)
	})

	if len(visited) != 1 || visited[0] != 42 {
		t.Errorf("PreOrder single: expected [42], got %v", visited)
	}
}

func TestDataTraversePreOrder(t *testing.T) {
	root := NewNode("a")
	root.AddNode("b")
	root.AddNode("c")

	var visited []string
	root.DataTraversePreOrder(func(data string) {
		visited = append(visited, data)
	})

	expected := []string{"a", "b", "c"}
	for i, v := range visited {
		if v != expected[i] {
			t.Errorf("DataTraversePreOrder: expected %v, got %v", expected, visited)
			break
		}
	}
}

// =============================================================================
// PostOrder Traversal Tests
// =============================================================================

func TestTraversePostOrder(t *testing.T) {
	//     1
	//    / \
	//   2   3
	//  /
	// 4
	root := NewNode(1)
	child2 := root.AddNode(2)
	root.AddNode(3)
	child2.AddNode(4)

	var visited []int
	root.TraversePostOrder(func(n *Node[int]) {
		visited = append(visited, n.Data)
	})

	// PostOrder: children first, then root
	expected := []int{4, 2, 3, 1}
	if len(visited) != len(expected) {
		t.Errorf("PostOrder: expected %d nodes, got %d", len(expected), len(visited))
	}
	for i, v := range visited {
		if v != expected[i] {
			t.Errorf("PostOrder: expected %v, got %v", expected, visited)
			break
		}
	}
}

func TestTraversePostOrderSingle(t *testing.T) {
	root := NewNode(42)

	var visited []int
	root.TraversePostOrder(func(n *Node[int]) {
		visited = append(visited, n.Data)
	})

	if len(visited) != 1 || visited[0] != 42 {
		t.Errorf("PostOrder single: expected [42], got %v", visited)
	}
}

func TestDataTraversePostOrder(t *testing.T) {
	root := NewNode("a")
	root.AddNode("b")
	root.AddNode("c")

	var visited []string
	root.DataTraversePostOrder(func(data string) {
		visited = append(visited, data)
	})

	expected := []string{"b", "c", "a"}
	for i, v := range visited {
		if v != expected[i] {
			t.Errorf("DataTraversePostOrder: expected %v, got %v", expected, visited)
			break
		}
	}
}

// =============================================================================
// Tree Structure Tests
// =============================================================================

func TestTreeDepth(t *testing.T) {
	root := NewNode(1)
	child := root.AddNode(2)
	grandchild := child.AddNode(3)

	// Count depth by following parent pointers
	depth := 0
	for n := grandchild; n != nil; n = n.Parent {
		depth++
	}

	if depth != 3 {
		t.Errorf("Tree depth: expected 3, got %d", depth)
	}
}

func TestFindAncestor(t *testing.T) {
	root := NewNode("root")
	child := root.AddNode("child")
	grandchild := child.AddNode("grandchild")

	// Walk up to find root
	current := grandchild
	for current.Parent != nil {
		current = current.Parent
	}

	if current != root {
		t.Error("FindAncestor: should find root")
	}
}

func TestCountNodes(t *testing.T) {
	root := NewNode(1)
	root.AddNode(2)
	child3 := root.AddNode(3)
	child3.AddNode(4)
	child3.AddNode(5)

	count := 0
	root.TraversePreOrder(func(n *Node[int]) {
		count++
	})

	if count != 5 {
		t.Errorf("CountNodes: expected 5, got %d", count)
	}
}

// =============================================================================
// AoC Specific Patterns
// =============================================================================

func TestTreeForDirectoryStructure(t *testing.T) {
	// Common AoC pattern: file system tree (e.g., 2022 Day 7)
	type DirEntry struct {
		Name  string
		Size  int
		IsDir bool
	}

	root := NewNode(DirEntry{Name: "/", IsDir: true})
	dir_a := root.AddNode(DirEntry{Name: "a", IsDir: true})
	root.AddNode(DirEntry{Name: "b.txt", Size: 100})
	dir_a.AddNode(DirEntry{Name: "c.txt", Size: 50})
	dir_a.AddNode(DirEntry{Name: "d.txt", Size: 75})

	// Calculate total size
	totalSize := 0
	root.TraversePostOrder(func(n *Node[DirEntry]) {
		if !n.Data.IsDir {
			totalSize += n.Data.Size
		}
	})

	if totalSize != 225 {
		t.Errorf("Directory tree: expected total 225, got %d", totalSize)
	}
}

func TestTreeForExpressionParsing(t *testing.T) {
	// Expression tree: (2 + 3) * 4
	//       *
	//      / \
	//     +   4
	//    / \
	//   2   3

	type Expr struct {
		Op    string
		Value int
	}

	mult := NewNode(Expr{Op: "*"})
	add := mult.AddNode(Expr{Op: "+"})
	mult.AddNode(Expr{Value: 4})
	add.AddNode(Expr{Value: 2})
	add.AddNode(Expr{Value: 3})

	// Evaluate using post-order
	var stack []int
	mult.TraversePostOrder(func(n *Node[Expr]) {
		if n.Data.Op == "" {
			stack = append(stack, n.Data.Value)
		} else {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch n.Data.Op {
			case "+":
				stack = append(stack, a+b)
			case "*":
				stack = append(stack, a*b)
			}
		}
	})

	if len(stack) != 1 || stack[0] != 20 {
		t.Errorf("Expression tree: expected 20, got %v", stack)
	}
}
