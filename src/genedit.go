package main

import (
	"fmt"
	"math"
)

func makeD(rows int, cols int) [][]float64 {
	d := make([][]float64, rows)
	for i := 0; i < len(d); i++ {
		d[i] = make([]float64, cols)
	}
	return d
}

func genEDBasic(A string, B string, G [][2]string, c []int) float64 {
	d := makeD(len(B)+1, len(A)+1)

	cost := func(A string, B string, pi int) float64 {
		a := len(A) - len(G[pi][0])
		b := len(B) - len(G[pi][1])
		if a >= 0 && b >= 0 && A[a:] == G[pi][0] && B[b:] == G[pi][1] {
			return d[b][a]+float64(c[pi])
		}
		return math.Inf(1)
	}

	minCost := func(A string, B string) float64 {
		min := math.Inf(1)
		for pi, _ := range G {
			min = math.Fmin(min, cost(A, B, pi))
		}
		return min
	}

	for x := 0; x <= len(A); x++ {
		for y := 0; y <= len(B); y++ {
			if x == 0 && y == 0 {
				d[y][x] = 0
			} else {
				d[y][x] = minCost(A[:x], B[:y])
			}
		}
	}

	return d[len(B)][len(A)]
}

func genEDAC(A string, B string, G [2][]string, c []int) float64 {
	d := makeD(len(B)+1, len(A)+1)

	Aroot := MakeLinkedGoto(G[0])
	MakeLinkedFail(Aroot)
	Broot := MakeLinkedGoto(G[1])
	MakeLinkedFail(Broot)

	Astate := Aroot
	Bstate := Broot

	minCost := func(A string, B string) float64 {
		min := math.Inf(1)
		p := Astate.output.Intersection(Bstate.output)
		p.ForEach(func(i int) {
			a := len(A) - len(G[0][i])
			b := len(B) - len(G[1][i])
			min = math.Fmin(min, d[b][a]+float64(c[i]))
		})
		return min
	}

	for x := 0; x <= len(A); x++ {
		Bstate = Broot
		for y := 0; y <= len(B); y++ {
			if x == 0 && y == 0 {
				d[y][x] = 0
			} else {
				d[y][x] = minCost(A[:x], B[:y])
			}

			if y < len(B) {
				Bstate = Bstate.Push(B[y])
			}
		}
		if x < len(A) {
			Astate = Astate.Push(A[x])
		}
	}

	return d[len(B)][len(A)]
}

func main() {
	A := "abcabc"
	B := "cdcd"
	g1 := [][2]string{{"abc", "cd"}, {"ab", "c"}, {"ca", "d"}, {"bc", "cd"}}
	g2 := [2][]string{{"abc", "ab", "ca", "bc"}, {"cd", "c", "d", "cd"}}
	c := []int{2, 1, 1, 1}

	fmt.Println(genEDBasic(A, B, g1, c))
	fmt.Println(genEDAC(A, B, g2, c))
}
