package structs

type Node[T comparable] struct {
	content T
	next    *Node[T]
}

func NewNode[T comparable](content T) *Node[T] {
	node := new(Node[T])
	node.content = content
	return node
}
