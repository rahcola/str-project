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
	symbol uint8
	output BitArray
	fail *ACNode
	children Children
}

func NewACNode(symbol uint8, outputSize int) (*ACNode) {
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

func (node *ACNode) isRoot() (bool) {
	return node.root
}

/*Standard binary search. If symbol is not found, returns false and
the index of the next biggest symbol in the vector. Take note, this
index might be outside of the vector!*/

func binarySearch(arr Children, left int, right int, symbol uint8) (int, bool) {
	if left > right {
		return left, false
	}

	mid := (right - left) / 2 + left
	val := arr[mid].symbol

	if val > symbol {
		return binarySearch(arr, left, mid - 1, symbol)
	}
	if val < symbol {
		return binarySearch(arr, mid + 1, right, symbol)
	}
	return mid, true
}

func (node *ACNode) AddChild(child *ACNode) {
	node.children = append(node.children, child)
	sort.Sort(node.children)
}

func (node *ACNode) LookupChild(symbol uint8) (*ACNode, bool) {
	i := sort.Search(len(node.children), func(i int) bool {
		return node.children[i].symbol >= symbol
	})

	if i == len(node.children) || node.children[i].symbol != symbol {
		if node.isRoot() {
			return node, true
		}
		return nil, false
	}
	return node.children[i], true

}

func (node *ACNode) Push(symbol uint8) (*ACNode) {
	_, found := node.LookupChild(symbol)
	for !found {
		node = node.fail
		_, found = node.LookupChild(symbol)
	}
	node, _ = node.LookupChild(symbol)
	return node
}