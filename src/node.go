package main

import (
	"container/vector"
)

type ACNode struct {
	root bool
	symbol int
	output *vector.Vector
	fail *ACNode
	children *vector.Vector
}

func NewACNode(symbol int) (*ACNode) {
	return &ACNode{false, symbol, new(vector.Vector), nil, new(vector.Vector)}
}

func NewRootACNode() (*ACNode) {
	return &ACNode{true, -1, new(vector.Vector), nil, new(vector.Vector)}
}

func (node *ACNode) isRoot() (bool) {
	return node.root
}

/*Standard binary search. If symbol is not found, returns false and
the index of the next biggest symbol in the vector. Take note, this
index might be outside of the vector!*/

func binarySearch(vec *vector.Vector, left int, right int, symbol int) (int, bool) {
	if left > right {
		return left, false
	}

	mid := (right - left) / 2 + left
	val := vec.At(mid).(*ACNode).symbol

	if val > symbol {
		return binarySearch(vec, left, mid - 1, symbol)
	}
	if val < symbol {
		return binarySearch(vec, mid + 1, right, symbol)
	}
	return mid, true
}

func (node *ACNode) AddChild(child *ACNode) {
	i, _ := binarySearch(node.children,
		0,
		node.children.Len() - 1,
		child.symbol)

	if i >= node.children.Len() {
		node.children.Push(child)
	} else {
		node.children.Insert(i, child)
	}
}

func (node *ACNode) LookupChild(symbol int) (*ACNode, bool) {
	i, found := binarySearch(node.children,
		0,
		node.children.Len() - 1,
		symbol)

	if !found {
		if node.isRoot() {
			return node, true
		}
		return nil, false
	}
	return node.children.At(i).(*ACNode), true

}