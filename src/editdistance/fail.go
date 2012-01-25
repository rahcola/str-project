package editdistance

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
		for _, s := range *r.children {
			s := s.(*ACNode)
			queue.PushBack(s)

			state := r.fail
			_, found := state.LookupChild(s.symbol)
			for !found {
				state = state.fail
				_, found = state.LookupChild(s.symbol)
			}

			s.fail, _ = state.LookupChild(s.symbol)
			s.output = s.output.Union(s.fail.output)
		}
	}
}