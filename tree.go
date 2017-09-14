package tree

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree *Node

func NewNode(v int) *Node {
	return &Node{value: v}
}
