package aoc

type Node[T any] struct {
	Data     T
	Children []*Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		Data:     data,
		Children: []*Node[T]{},
	}
}

func (n *Node[T]) AddNode(data T) (child *Node[T]) {
	child = NewNode(data)
	n.Children = append(n.Children, child)
	return
}

func (n *Node[T]) TraversePreOrder(f func(*Node[T])) {
	f(n)
	for _, i := range n.Children {
		i.TraversePreOrder(f)
	}
}

func (n *Node[T]) TraversePostOrder(f func(*Node[T])) {
	for _, i := range n.Children {
		i.TraversePostOrder(f)
	}
	f(n)
}

func (n *Node[T]) DataTraversePreOrder(f func(T)) {
	n.TraversePreOrder(func(x *Node[T]) {
		f(x.Data)
	})
}

func (n *Node[T]) DataTraversePostOrder(f func(T)) {
	n.TraversePostOrder(func(x *Node[T]) {
		f(x.Data)
	})
}
