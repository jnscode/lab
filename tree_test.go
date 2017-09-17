package tree

import (
	"testing"
)

func CreateTree() *Node {
	n0 := NewNode(1)
	n1 := NewNode(2)
	n2 := NewNode(3)
	n3 := NewNode(4)
	n4 := NewNode(5)

	n0.left = n1
	n0.right = n2
	n1.left = n3
	n1.right = n4
	return n0
}

func TestCreateNode(t *testing.T) {
	n := NewNode(1)
	if n == nil {
		t.Error("create node failed")
		return
	}

	if n.value != 1 {
		t.Error("init node failed")
		return
	}
}

func TestCreateTree(t *testing.T) {
	r := NewNode(1)
	r.left = NewNode(2)
	r.right = NewNode(3)
}

func TestDeepSearch(t *testing.T) {
	r := CreateTree()
	n := r.DeepSearch(3)
	println(n)
}

func TestBreadthSearch(t *testing.T) {
	r := CreateTree()
	n := r.BreadthSearch(5)
	println(n)
}
