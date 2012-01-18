package main

import (
	. "utf8"
	"fmt"
)

func main() {
	keywords := []*String{NewString("he"), NewString("she")}

	_, o := MakeGoto(keywords)
	a, ok := o(0)
	if ok {
		fmt.Println(a)
	}
}
