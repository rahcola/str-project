package main

import (
	"container/list"
)

func MakeLinkedFail(root *ACNode) {
	queue := list.New()
	
	for _, child := range *root.children {
		queue.PushBack(child)
		child.(*ACNode).fail = root
	}

	for queue.Len() > 0 {
		r := queue.Remove(queue.Front()).(*ACNode)
		for _, child := range *r.children {
			child := child.(*ACNode)
			queue.PushBack(child)

			state := r.fail
			_, found := state.LookupChild(child.symbol)
			for !found && state != root {
				state = state.fail
				_, found = state.LookupChild(child.symbol)
			}

			child.fail, _ = state.LookupChild(child.symbol)
			if child.fail == nil {
				child.fail = root
			}
			child.output.AppendVector(child.fail.output)
		}
	}
}