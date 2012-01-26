package editdistance

import (
	"testing"
)

var data []string = readLines("dna.50MB")

/* Long words */

func BenchmarkBasicFewShortPatternsLongWords(b *testing.B) {
	b.StopTimer()
	rules, costs := PatternRuleToPaired(ShortPatterns("ACGT", 20, 30))
	genedit := MakeBasicGenEdit(rules, costs)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N && i < len(data); i = i + 2 {
		genedit(data[i][:1000], data[i+1][:1000])
	}
}

func BenchmarkACFewShortPatternsLongWords(b *testing.B) {
	b.StopTimer()
	rules, costs := PatternRuleToListed(ShortPatterns("ACGT", 20, 30))
	genedit := MakeACGenEdit(rules, costs)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N && i < len(data); i = i + 2 {
		genedit(data[i][:1000], data[i+1][:1000])
	}
}

func BenchmarkBasicManyShortPatternsLongWords(b *testing.B) {
	b.StopTimer()
	rules, costs := PatternRuleToListed(ShortPatterns("ACGT", 2000, 3000))
	genedit := MakeBasicGenEdit(rules, costs)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N && i < len(data); i = i + 2 {
		genedit(data[i][:1000], data[i+1][:1000])
	}
}

func BenchmarkACManyShortPatternsLongWords(b *testing.B) {
	b.StopTimer()
	rules, costs := PatternRuleToListed(ShortPatterns("ACGT", 2000, 3000))
	genedit := MakeACGenEdit(rules, costs)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N && i < len(data); i = i + 2 {
		genedit(data[i][:1000], data[i+1][:1000])
	}
}

func BenchmarkBasicFewLongPatternsLongWords(b *testing.B) {
	b.StopTimer()
	rules, costs := PatternRuleToListed(LongPatterns("ACGT", 20, 30))
	genedit := MakeBasicGenEdit(rules, costs)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N && i < len(data); i = i + 2 {
		genedit(data[i][:1000], data[i+1][:1000])
	}
}

func BenchmarkACFewLongPatternsLongWords(b *testing.B) {
	b.StopTimer()
	rules, costs := PatternRuleToListed(LongPatterns("ACGT", 20, 30))
	genedit := MakeACGenEdit(rules, costs)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N && i < len(data); i = i + 2 {
		genedit(data[i][:1000], data[i+1][:1000])
	}
}
/*
func BenchmarkManyLongPatternsLongWords(*testing.B) {

}

func BenchmarkFewLongAndShortPatternsLongWords(*testing.B) {

}

func BenchmarkManyLongAndShortPatternsLongWords(*testing.B) {

}
*/
/* English words */
/*
func BenchmarkFewShortPatternsEnglishWords(*testing.B) {

}

func BenchmarkManyShortPatternsEnglishWords(*testing.B) {

}

func BenchmarkFewLongPatternsEnglishWords(*testing.B) {

}

func BenchmarkManyLongPatternsEnglishWords(*testing.B) {

}

func BenchmarkFewLongAndShortPatternsEnglishWords(*testing.B) {

}

func BenchmarkManyLongAndShortPatternsEnglishWords(*testing.B) {

}
*/