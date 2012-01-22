package main

import (

)

func MakeLinkedGoto(keywords []string) (*ACNode) {
	tmpRoot := NewACNode(0)

	for _, keyword := range keywords {
		state := tmpRoot
		for _, symbol := range keyword {
			nextState, found := state.LookupChild(symbol)
			if found {
				state = nextState
			} else {
				newstate := NewACNode(symbol)
				state.AddChild(newstate)
				state = newstate
			}
		}
		state.output.Push(keyword)
	}

	root := NewRootACNode()
	root.children = tmpRoot.children
	return root
}
