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
					min := math.Inf(1)
					if x > 0 && y > 0 && Astate.symbol == Bstate.symbol {
						min = d[y-1][x-1]
					}
					if len(Astate.output) > 0 && len(Bstate.output) > 0 {
						p := Astate.output.Intersection(Bstate.output)
						index := 0
						for i := 0; i < len(p); i++ {
							word := p[i]
							for k := 0; k < 32; k++ {
								if word & 1 != 0 {
									a := x - utf8.NewString(G[0][index]).RuneCount()
									b := y - utf8.NewString(G[1][index]).RuneCount()
									min = math.Fmin(min, d[b][a]+c[index])
								}
								word = word >> 1
								index++
							}
						}
					}
					d[y][x] = min
				}
				Bstate = Bstate.Push(B.At(y))
			}
			Astate = Astate.Push(A.At(x))
		}

		return d[BLen][ALen]
	}
}
