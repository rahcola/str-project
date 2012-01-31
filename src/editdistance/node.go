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

func BinarySearch(arr Children, symbol int) (*ACNode, bool) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		val := arr[mid].symbol
		if val > symbol {
			high = mid - 1
		} else	if val < symbol {
			low = mid + 1
		} else {
			return arr[mid], true
		}
	}
	return nil, false
}

func (node *ACNode) Output() BitArray {
	r := NewBitArray(0).Union(node.output)
	v := node.fail
	for v != nil {
		r = r.Union(v.output)
		v = v.fail
	}
	return r
}

func (node *ACNode) AddChild(child *ACNode) {
	node.children = append(node.children, child)
	sort.Sort(node.children)
}

func (node *ACNode) LookupChild(symbol int) (*ACNode, bool) {
	child, found := BinarySearch(node.children, symbol)

	if !found {
		if node.root {
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