package editdistance

import (
	"math"
	"utf8"
)

func MakeBasicGenEdit(G [][]string, c []float64) (func(string, string) float64) {
	cost := func(A string, B string, pi int, d [][]float64) float64 {
		a := len(A) - len(G[pi][0])
		b := len(B) - len(G[pi][1])
		if a >= 0 && b >= 0 && A[a:] == G[pi][0] && B[b:] == G[pi][1] {
			return d[b][a]+c[pi]
		}
		return math.Inf(1)
	}

	return func(Ap string, Bp string) float64 {
		A := utf8.NewString(Ap)
		B := utf8.NewString(Bp)
		d := makeMatrix(B.RuneCount()+1, A.RuneCount()+1)

		for x := 0; x <= A.RuneCount(); x++ {
			for y := 0; y <= B.RuneCount(); y++ {
				if x == 0 && y == 0 {
					d[y][x] = 0
				} else {
					min := math.Inf(1)
					if x > 0 && y > 0 && A.At(x-1) == B.At(y-1) {
						min = d[y-1][x-1]
					}
					for pi, _ := range G {
						min = math.Fmin(min,
							cost(A.Slice(0, x),
							B.Slice(0, y),
							pi,
							d))
					}

					d[y][x] = min
				}
			}
		}

		return d[B.RuneCount()][A.RuneCount()]
	}
}