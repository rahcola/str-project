package editdistance

import (
	"testing"
)

func TestOne(t *testing.T) {
	A := "abcabc"
	B := "cdcd"
	g1 := [][]string{{"abc", "cd"}, {"ab", "c"}, {"ca", "d"}, {"bc", "cd"}}
	g2 := [][]string{{"abc", "ab", "ca", "bc"}, {"cd", "c", "d", "cd"}}
	c := []float64{2, 1, 1, 1}

	basic := MakeBasicGenEdit(g1, c)
	ac := MakeACGenEdit(g2, c)

	if basic(A, B) != 3 {
		t.Error("basic(A, B) returned:", basic(A, B))
	}
	if ac(A, B) != 3 {
		t.Error("ac(A, B) returned:", ac(A, B))
	}
}

func TestTwo(t *testing.T) {
	A := "helmi"
	B := "kuppi"
	g1 := [][]string{{"h", "k"}, {"e", "u"}, {"l", "p"}, {"m", "p"}}
	g2 := [][]string{{"h", "e", "l", "m"}, {"k", "u", "p", "p"}}
	c := []float64{1, 1, 1, 1}

	basic := MakeBasicGenEdit(g1, c)
	ac := MakeACGenEdit(g2, c)

	if basic(A, B) != 4 {
		t.Error("basic(A, B) returned:", basic(A, B))
	}
	if ac(A, B) != 4 {
		t.Error("ac(A, B) returned:", ac(A, B))
	}
}

func TestThree(t *testing.T) {
	A := "helmi"
	B := "kuppi"
	g1 := [][]string{{"h", "k"}, {"e", "u"}, {"l", "p"}, {"m", "p"}, {"i", "i"}}
	g2 := [][]string{{"h", "e", "l", "m", "i"}, {"k", "u", "p", "p", "i"}}
	c := []float64{1, 1, 1, 1, -1}

	basic := MakeBasicGenEdit(g1, c)
	ac := MakeACGenEdit(g2, c)

	if basic(A, B) != 3 {
		t.Error("basic(A, B) returned:", basic(A, B))
	}
	if ac(A, B) != 3 {
		t.Error("ac(A, B) returned:", ac(A, B))
	}
}