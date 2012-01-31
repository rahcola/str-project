package editdistance

import (
	"math"
	"utf8"
)

func MakeACGenEdit(Gp [][]string, c []float64) func(string, string) float64 {
	Aroot := MakeLinkedGoto(Gp[0])
	MakeLinkedFail(Aroot)
	Broot := MakeLinkedGoto(Gp[1])
	MakeLinkedFail(Broot)
	G := make([][]*utf8.String, len(Gp))
	for i, _ := range G {
		G[i] = make([]*utf8.String, len(Gp[i]))
		for j, _ := range G[i] {
			G[i][j] = utf8.NewString(Gp[i][j])
		}
	}

	return func(Ap string, Bp string) float64 {
		A := utf8.NewString(Ap+" ")
		ALen := A.RuneCount()-1
		B := utf8.NewString(Bp+" ")
		BLen := B.RuneCount()-1
		d := makeMatrix(BLen+1, ALen+1)

		Astate := Aroot
		Bstate := Broot

		for x := 0; x <= ALen; x++ {
			Bstate = Broot
			for y := 0; y <= BLen; y++ {
				if x == 0 && y == 0 {
					d[y][x] = 0
				} else {
					d[y][x] = math.Inf(1)
					if x > 0 && y > 0 && Astate.symbol == Bstate.symbol {
						d[y][x] = d[y-1][x-1]
					}
					if len(Astate.output) > 0 && len(Bstate.output) > 0 {
						p := Astate.output.Intersection(Bstate.output)
						p.ForEach(func (i int) {
							a := x - G[0][i].RuneCount()
							b := y - G[1][i].RuneCount()
							d[y][x] = math.Fmin(d[y][x], d[b][a]+c[i])
						})
					}
				}
				Bstate = Bstate.Push(B.At(y))
			}
			Astate = Astate.Push(A.At(x))
		}

		return d[BLen][ALen]
	}
}
