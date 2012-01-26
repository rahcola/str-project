package editdistance

import (

)

func MakeLinkedGoto(keywords []string) (*ACNode) {
	tmpRoot := NewACNode(0, len(keywords))

	for i, keyword := range keywords {
		state := tmpRoot
		for _, symbol := range keyword {
			symbol := uint8(symbol)
			nextState, found := state.LookupChild(symbol)
			if found {
				state = nextState
			} else {
				newstate := NewACNode(symbol, len(keywords))
				state.AddChild(newstate)
				state = newstate
			}
		}
		state.output = state.output.Set(i)
	}

	root := NewRootACNode(len(keywords))
	root.children = tmpRoot.children
	return root
}
