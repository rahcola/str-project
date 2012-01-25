package main

func makeMatrix(rows int, cols int) [][]float64 {
	d := make([][]float64, rows)
	for i := 0; i < len(d); i++ {
		d[i] = make([]float64, cols)
	}
	return d
}