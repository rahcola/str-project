package editdistance

import (
	"sort"
)

type Children []*ACNode

func (c Children) Len() int {
	return len(c)
}

func (c Children) Less(i int, j int) bool {
	return c[i].symbol < c[j].symbol
}

func (c Children) Swap(i int, j int) {
	c[i], c[j] = c[j], c[i]
}

type ACNode struct {
	root bool
	symbol int
	output BitArray
	fail *ACNode
	children Children
}

func NewACNode(symbol int, outputSize int) (*ACNode) {
	return &ACNode{false,
		symbol,
		NewBitArray(outputSize),
		nil,
		make([]*ACNode, 0, 1)}
}

func NewRootACNode(outputSize int) (*ACNode) {
	return &ACNode{true,
		0,
		NewBitArray(outputSize),
		nil,
		make([]*ACNode, 0, 1)}
}

func (node ACNode) isRoot() (bool) {
	return node.root
}

func BinarySearch(arr Children, symbol int) (*ACNode, bool) {
	if len(arr) == 0 {
		return nil, false
	}

	mid := len(arr) / 2
	val := arr[mid].symbol
	if val > symbol {
		return BinarySearch(arr[:mid], symbol)
	}
	if val < symbol {
		return BinarySearch(arr[mid+1:], symbol)
	}
	return arr[mid], true
}

func (node *ACNode) AddChild(child *ACNode) {
	node.children = append(node.children, child)
	sort.Sort(node.children)
}

func (node *ACNode) LookupChild(symbol int) (*ACNode, bool) {
	child, found := BinarySearch(node.children, symbol)

	if !found {
		if node.isRoot() {
			return node, true
		}
		return nil, false
	}
	return child, true

}

func (node *ACNode) Push(symbol int) (*ACNode) {
	_, found := node.LookupChild(symbol)
	for !found {
		node = node.fail
		_, found = node.LookupChild(symbol)
	}
	node, _ = node.LookupChild(symbol)
	return node
}