package main

import (
"utf8"
)

/*
func updateGoto(g map[State](map[int]State),
state State,
symbol int,
newstate State) {
	_, ok := g[state]
	if !ok {
		g[state] = make(map[int]State)
	}
	g[state][symbol] = newstate
}
*/
func MakeLinkedGoto(keywords []*utf8.String) (*ACNode) {
	root := NewACNode(-1)
	//root.fail = root

	for i, keyword := range keywords {
		kwLen := keyword.RuneCount()
		if kwLen < 1 {
			continue
		}

		state := root
		j := 0
		nextState, found := state.LookupChild(keyword.At(j));
		for found {
			state = nextState
			j++
			nextState, found = state.LookupChild(keyword.At(j));
		}

		for p := j; p < kwLen; p++ {
			newstate := NewACNode(keyword.At(p))
			state.AddChild(newstate)
			state = newstate
		}

		state.output.Push(i)
	}

	return root
}
/*
func MakeGoto(keywords []*utf8.String) (func(State, int) (State, bool),
	                                func(State) ([]*utf8.String, bool)) {
	g := make(map[State]map[int]State)
	output := make(map[State][]*utf8.String)
	var newstate State = 0

	for _, keyword := range keywords {
		if keyword.RuneCount() == 0 {
			continue
		}
		var state State = 0
		j := 0

		next_state, ok := g[state][keyword.At(j)]
		for ok {
			state = next_state
			j++
			next_state, ok = g[state][keyword.At(j)]
		}

		for p := j; p < keyword.RuneCount(); p++ {
			newstate++
			updateGoto(g, state, keyword.At(p), newstate)
			state = newstate
		}

		output[state] = append(output[state], keyword)
	}

	return func(state State, symbol int) (State, bool) {
		newstate, ok := g[state][symbol]
		if !ok {
			if state == 0 {
				return 0, true
			} else {
				return -1, false
			}
		}
		return newstate, true
	}, func(state State) ([]*utf8.String, bool) {
		s, ok := output[state]
		if !ok {
			return nil, false
		}
		return s, true
	}
}
*/