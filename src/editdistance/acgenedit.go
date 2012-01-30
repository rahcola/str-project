package editdistance

import (
	"math"
	"utf8"
)

func MakeACGenEdit(G [][]string, c []float64) func(string, string) float64 {
	Aroot := MakeLinkedGoto(G[0])
	MakeLinkedFail(Aroot)
	Broot := MakeLinkedGoto(G[1])
	MakeLinkedFail(Broot)

	minCost := func(Ap string,
	Astate *ACNode,
	Bp string,
	Bstate *ACNode,
	d [][]float64) float64 {
		A := utf8.NewString(Ap)
		B := utf8.NewString(Bp)
		min := math.Inf(1)
		if A.RuneCount() > 0 && B.RuneCount() > 0 && Astate.symbol == Bstate.symbol {
			min = d[B.RuneCount()-1][A.RuneCount()-1]
		}
		p := Astate.output.Intersection(Bstate.output)
		p.ForEach(func(i int) {
			a := A.RuneCount() - utf8.NewString(G[0][i]).RuneCount()
			b := B.RuneCount() - utf8.NewString(G[1][i]).RuneCount()
			min = math.Fmin(min, d[b][a]+c[i])
		})
		return min
	}

	return func(Ap string, Bp string) float64 {
		A := utf8.NewString(Ap)
		B := utf8.NewString(Bp)
		d := makeMatrix(B.RuneCount()+1, A.RuneCount()+1)

		Astate := Aroot
		Bstate := Broot

		for x := 0; x <= A.RuneCount(); x++ {
			Bstate = Broot
			for y := 0; y <= B.RuneCount(); y++ {
				if x == 0 && y == 0 {
					d[y][x] = 0
				} else {
					d[y][x] = minCost(A.Slice(0, x),
						Astate,
						B.Slice(0, y),
						Bstate,
						d)
				}

				if y < B.RuneCount() {
					Bstate = Bstate.Push(B.At(y))
				}
			}
			if x < A.RuneCount() {
				Astate = Astate.Push(A.At(x))
			}
		}

		return d[B.RuneCount()][A.RuneCount()]
	}
}
