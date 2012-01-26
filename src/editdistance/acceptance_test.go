package editdistance

import (
	"testing"
)

func TestZero(t *testing.T) {
	root := MakeLinkedGoto([]string{"she", "he", "his", "hers"})
	MakeLinkedFail(root)

	o := make([]bool, 4)
	ac := root.Push('h').Push('e')
	ac.output.ForEach(func(i int) {
		o[i] = true
	})
	if !o[1] {
		t.Error("Output for 'he' wrong: ", o)
	}

	o = make([]bool, 4)
	ac = root.Push('h').Push('e').Push('r').Push('s')
	ac.output.ForEach(func(i int) {
		o[i] = true
	})
	if !o[3] {
		t.Error("Output for 'hers' wrong: ", o)
	}

	o = make([]bool, 4)
	ac = root.Push('h').Push('i').Push('s')
	ac.output.ForEach(func(i int) {
		o[i] = true
	})
	if !o[2] {
		t.Error("Output for 'his' wrong: ", o)
	}

	o = make([]bool, 4)
	ac = root.Push('s').Push('h').Push('e')
	ac.output.ForEach(func(i int) {
		o[i] = true
	})
	if !(o[0] && o[1]) {
		t.Error("Output for 'she' wrong: ", o)
	}
}

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