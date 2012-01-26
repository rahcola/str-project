package editdistance

import (
	"math"
)

func MakeACGenEdit(G [][]string, c []float64) func(string, string) float64 {
	Aroot := MakeLinkedGoto(G[0])
	MakeLinkedFail(Aroot)
	Broot := MakeLinkedGoto(G[1])
	MakeLinkedFail(Broot)

	minCost := func(A string,
	Astate *ACNode,
	B string,
	Bstate *ACNode,
	d [][]float64) float64 {
		min := math.Inf(1)
		if len(A) > 0 && len(B) > 0 && Astate.symbol == Bstate.symbol {
			min = d[len(B)-1][len(A)-1]
		}
		p := Astate.Output().Intersection(Bstate.Output())
		p.ForEach(func(i int) {
			a := len(A) - len(G[0][i])
			b := len(B) - len(G[1][i])
			min = math.Fmin(min, d[b][a]+c[i])
		})
		return min
	}

	return func(A string, B string) float64 {
		d := makeMatrix(len(B)+1, len(A)+1)

		Astate := Aroot
		Bstate := Broot

		for x := 0; x <= len(A); x++ {
			Bstate = Broot
			for y := 0; y <= len(B); y++ {
				if x == 0 && y == 0 {
					d[y][x] = 0
				} else {
					d[y][x] = minCost(A[:x],
						Astate,
						B[:y],
						Bstate,
						d)
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
}
