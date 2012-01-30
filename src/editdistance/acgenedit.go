package editdistance

import (
	"math"
)

func MakeACGenEdit(G [][]string, c []float64) func(string, string) float64 {
	Aroot := MakeLinkedGoto(G[0])
	MakeLinkedFail(Aroot)
	Broot := MakeLinkedGoto(G[1])
	MakeLinkedFail(Broot)

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
					min := math.Inf(1)
					if x > 0 && y > 0 && Astate.symbol == Bstate.symbol {
						min = d[y-1][x-1]
					}
					p := Astate.output.Intersection(Bstate.output)
					p.ForEach(func(i int) {
						a := x - len(G[0][i])
						b := y - len(G[1][i])
						min = math.Fmin(min, d[b][a]+c[i])
					})
					d[y][x] = min
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
