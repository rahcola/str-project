package editdistance

import (

)

func MakeLinkedGoto(keywords []string) (*ACNode) {
	tmpRoot := NewACNode(0)

	for i, keyword := range keywords {
		state := tmpRoot
		for _, symbol := range keyword {
			symbol := uint8(symbol)
			nextState, found := state.LookupChild(symbol)
			if found {
				state = nextState
			} else {
				newstate := NewACNode(symbol)
				state.AddChild(newstate)
				state = newstate
			}
		}
		state.rep = state.rep.Set(i)
	}

	root := NewRootACNode()
	root.children = tmpRoot.children
	return root
}
