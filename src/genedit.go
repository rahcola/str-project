package main

import (
	. "utf8"
	"fmt"
)

func main() {
	keywords := []*String{NewString("he"),
		NewString("she"),
		NewString("hers"),
		NewString("his")}

	root := MakeLinkedGoto(keywords)
	MakeLinkedFail(root)

	state := root
	for i, c := range "ojfeiuhewureghreoijn" {
		_, found := state.LookupChild(c)
		for !found && state != root {
			state = state.fail
			_, found = state.LookupChild(c)
		}

		state, _ = state.LookupChild(c)

		if state.output.Len() > 0 {
			fmt.Println(i, " ", state.output)
		}
	}
}
