package editdistance

import (
	"math"
	"utf8"
)

func MakeACGenEdit(Gp [][]string, c []float64) func(string, string) float64 {
	Proot := MakeLinkedGoto(Gp[0])
	MakeLinkedFail(Proot)
	Troot := MakeLinkedGoto(Gp[1])
	MakeLinkedFail(Troot)
	G := make([][]*utf8.String, len(Gp))
	for i, _ := range G {
		G[i] = make([]*utf8.String, len(Gp[i]))
		for j, _ := range G[i] {
			G[i][j] = utf8.NewString(Gp[i][j])
		}
	}

	return func(pattern string, text string) float64 {
		P := utf8.NewString(pattern+" ")
		PLen := P.RuneCount()-1
		T := utf8.NewString(text+" ")
		TLen := T.RuneCount()-1
		d := makeMatrix(TLen+1, PLen+1)
		p := NewBitArray(len(Gp[0]))

		Pstate := Proot
		Tstate := Troot

		for x := 0; x <= PLen; x++ {
			Tstate = Troot
			for y := 0; y <= TLen; y++ {
				if x == 0 && y == 0 {
					d[y][x] = 0
				} else {
					d[y][x] = math.Inf(1)
					if x > 0 && y > 0 && Pstate.symbol == Tstate.symbol {
						d[y][x] = d[y-1][x-1]
					}
					if len(Pstate.output) > 0 && len(Tstate.output) > 0 {
						p = p.Intersection(Pstate.output, Tstate.output)
						p.ForEach(func (i int) {
							a := x - G[0][i].RuneCount()
							b := y - G[1][i].RuneCount()
							d[y][x] = math.Fmin(d[y][x], d[b][a]+c[i])
						})
					}
				}
				Tstate = Tstate.Push(T.At(y))
			}
			Pstate = Pstate.Push(P.At(x))
		}

		return d[TLen][PLen]
	}
}
