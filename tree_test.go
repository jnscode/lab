package tree

import (
	"testing"
)

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
