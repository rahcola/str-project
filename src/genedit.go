package main

import (
	"fmt"
)

func main() {
	A := "abcabc"
	B := "cdcd"
	g1 := [][2]string{{"abc", "cd"}, {"ab", "c"}, {"ca", "d"}, {"bc", "cd"}}
	g2 := [2][]string{{"abc", "ab", "ca", "bc"}, {"cd", "c", "d", "cd"}}
	c := []float64{2, 1, 1, 1}

	fmt.Println(MakeBasicGenEdit(g1, c)(A, B))
	fmt.Println(MakeACGenEdit(g2, c)(A, B))
}
