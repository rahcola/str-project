package editdistance

import (
	"utf8"
	"strings"
	"rand"
	"os"
	"bufio"
)

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
	buf := make([]byte, 0, 1024)
	for err == nil {
		buf = append(buf, line...)
		if !isPrefix {
			lines = append(lines, string(buf))
			buf = make([]byte, 0, 1024)
		}
		line, isPrefix, err = reader.ReadLine()
	}
	return lines
}

func Substrings(str string) []string {
	s := utf8.NewString(str)
	n := s.RuneCount()
	substrings := make([]string, 0, (n * (n+1)) / 2)
	for start := 0; start < n; start++ {
		for end := start + 1; end <= n; end++ {
			substrings = append(substrings, s.Slice(start, end))
		}
	}
	return substrings
}

func Reverse(str string) string {
	n := len(str)
	ret := make([]int, n)
	for _, rune := range str {
		n--
		ret[n] = rune
	}
	return string(ret[n:])
}

func DNAInversion(str string) string {
	return strings.Map(func (rune int) int {
		x := -1
		switch rune {
		case 'A': x = 'T'
		case 'T': x = 'A'
		case 'G': x = 'C'
		case 'C': x = 'G'
		}
		return x
	},
		Reverse(str))
}

type PatternRule struct {
	leftSide string
	rightSide string
	cost float64
}

func RandomString(alpha string, min int, max int) string {
	alphabet := utf8.NewString(alpha)
	runes := make([]int, rand.Intn(max-(min-1)) + min)
	for i, _ := range runes {
		p := rand.Intn(alphabet.RuneCount())
		runes[i] = alphabet.At(p)
	}
	return string(runes)
}

func DNAInversionRules(pattern string) []PatternRule {
	substrings := Substrings(pattern)
	inversions := make([]string, len(substrings))
	for i, s := range substrings {
		inversions[i] = DNAInversion(s)
	}

	rules := make([]PatternRule, len(substrings))
	for i, _ := range rules {
		rules[i] = PatternRule{substrings[i], inversions[i], 0}
	}
	return rules
}

func GenPatternRule(alpha string, min int, max int) *PatternRule {
	leftSide := RandomString(alpha, min, max)
	rightSide := RandomString(alpha, min, max)
	cost := rand.Float64() * 10

	return &PatternRule{leftSide, rightSide, cost}
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