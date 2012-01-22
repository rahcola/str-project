package main

import (
	"fmt"
	"math"
)

func genEDBasic(A string, B string, G [][]string, c []int) (float64) {
	d := make([][]float64, len(A) + 1)
	for i := 0; i < len(A) + 1; i++ {
		d[i] = make([]float64, len(B) + 1)
	}

	minCost := func (A string, B string) (float64) {
		min := math.Inf(1)
		for pi, p := range G {
			for a := 0; a <= len(A); a++ {
				for b := 0; b <= len(B); b++ {
					if (A[a:] == p[0] && B[b:] == p[1]) {
						min = math.Fmin(d[a][b] + float64(c[pi]), min)
					}
				}
			}
		}
		return min
	}

	for row := 0; row <= len(A); row++ {
		for col := 0; col <= len(B); col++ {
			if row == 0 && col == 0 {
				continue
			}
			d[row][col] = minCost(A[:row], B[:col])
		}
	}
	return d[len(A)][len(B)]
}

func genEDAhoCorasick(A string, B string, G [][]string, c []int) (float64) {
	
}

func main() {

	A := "abcabc"
	B := "cdcd"
	g := [][]string{{"abc", "cd"}, {"ab", "c"}, {"ca", "d"}, {"bc", "cd"}}
	c := []int{2, 1, 1, 1}

	fmt.Println(genEDBasic(A, B, g, c))

	/*
	keywords := []string{"he", "äß", "she", "hers", "his"}

	root := MakeLinkedGoto(keywords)
	MakeLinkedFail(root)

	state := root
	for i, sym := range "foiewiufäßnhiuadfherseoij" {
		_, found := state.LookupChild(sym)
		for !found {
			state = state.fail
			_, found = state.LookupChild(sym)
		}

		state, _ = state.LookupChild(sym)
		if state.output.Len() > 0 {
			fmt.Print(i)
			fmt.Println(state.output)
		}
	}
	 */
}
