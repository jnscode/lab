package tree

import (
	"container/list"
)

type Type int // interface{}

type Node struct {
	value Type
	left  *Node
	right *Node
}

func NewNode(v Type) *Node {
	return &Node{value: v}
}

func (t *Node) DeepSearch(v Type) *Node {
	if t == nil {
		println("root is nil")
		return nil
	}

	println("check value", t.value)
	if t.value == v {
		println("find it")
		return t
	}

	println("search left node")
	l := t.left.DeepSearch(v)
	if l != nil {
		return l
	}

	println("search right node")
	r := t.right.DeepSearch(v)
	if r != nil {
		return r
	}

	return nil
}

func (t *Node) BreadthSearch(v Type) *Node {
	if t == nil {
		println("root is nil")
		return nil
	}

	q := list.New()
	q.PushBack(t)

	for i := q.Front(); i != nil; {
		n := i.Value.(*Node)
		println("check value", n.value)

		if n.value == v {
			println("find it")
			return n
		}

		if n.left != nil {
			q.PushBack(n.left)
		}

		if n.right != nil {
			q.PushBack(n.right)
		}

		cur := i
		i = i.Next()
		q.Remove(cur)
	}

	return nil
}
