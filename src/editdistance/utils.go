package editdistance

import (
	"utf8"
	"strings"
	"rand"
	"os"
	"bufio"
)

func IntMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func makeMatrix(rows int, cols int) [][]float64 {
	d := make([][]float64, rows)
	for i := 0; i < len(d); i++ {
		d[i] = make([]float64, cols)
	}
	return d
}

func readLines(path string) []string {
	lines := make([]string, 0, 100)

	file, err := os.Open(path)
	if err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	line, isPrefix, err := reader.ReadLine()
	buf := make([]byte, 1024)
	for err == nil {
		buf = append(buf, line...)
		if !isPrefix {
			lines = append(lines, string(buf))
			buf = make([]byte, 1024)
		}
		line, isPrefix, err = reader.ReadLine()
	}
	return lines
}

type PatternRule struct {
	leftSide string
	rightSide string
	cost float64
}

func GenPatternRule(alpha string, min int, max int) *PatternRule {
	alphabet := utf8.NewString(alpha)

	leftSide := make([]string, rand.Intn(max-(min-1))+min)
	for i, _ := range leftSide {
		p := rand.Intn(alphabet.RuneCount())
		leftSide[i] = alphabet.Slice(p, p+1)
	}

	rightSide := make([]string, rand.Intn(max-(min-1))+min)
	for i, _ := range rightSide {
		p := rand.Intn(alphabet.RuneCount())
		rightSide[i] = alphabet.Slice(p, p+1)
	}

	cost := rand.Float64() * 10

	return &PatternRule{strings.Join(leftSide, ""),
		strings.Join(rightSide, ""),
		cost}
}

func ShortPatterns(alphabet string, min int, max int) []PatternRule {
	rules := make([]PatternRule, rand.Intn(max-(min-1))+min)
	for i, _ := range rules {
		rules[i] = *GenPatternRule(alphabet, 1, 3)
	}
	return rules
}

func LongPatterns(alphabet string, min int, max int) []PatternRule {
	rules := make([]PatternRule, rand.Intn(max-(min-1))+min)
	for i, _ := range rules {
		rules[i] = *GenPatternRule(alphabet, 100, 300)
	}
	return rules
}

func PatternRuleToPaired(rules []PatternRule) ([][]string, []float64) {
	pairs := make([][]string, len(rules))
	costs := make([]float64, len(rules))
	for i, r := range rules {
		pair := make([]string, 2)
		pair[0] = r.leftSide
		pair[1] = r.rightSide

		pairs[i] = pair
		costs[i] = r.cost
	}
	return pairs, costs
}

func PatternRuleToListed(rules []PatternRule) ([][]string, []float64) {
	leftSides := make([]string, len(rules))
	rightSides := make([]string, len(rules))
	costs := make([]float64, len(rules))
	for i, r := range rules {
		leftSides[i] = r.leftSide
		rightSides[i] = r.rightSide
		costs[i] = r.cost
	}

	ret := make([][]string, 2)
	ret[0] = leftSides
	ret[1] = rightSides
	return ret, costs
}