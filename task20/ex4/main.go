package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TextAnalyzer struct {
	Text string
}

func (t TextAnalyzer) CountWords() int {
	words := strings.Fields(t.Text)
	return len(words)
}

func (t TextAnalyzer) CountSentences() int {
	count := 0
	for _, ch := range t.Text {
		if ch == '.' || ch == '!' || ch == '?' {
			count++
		}
	}
	return count
}

func (t TextAnalyzer) CountVowels() int {
	vowels := "aeiouyAEIOUY"
	count := 0

	for _, ch := range t.Text {
		for _, v := range vowels {
			if ch == v {
				count++
			}
		}
	}
	return count
}

type Analyzer interface {
	CountWords() int
	CountSentences() int
	CountVowels() int
}

func main() {
	text := TextAnalyzer{
		Text: "Go is fast. It is simple! Do you like Go?",
	}

	var a Analyzer = text

	fmt.Printf("Слов: %d, предложений: %d, гласных: %d\n",
		a.CountWords(),
		a.CountSentences(),
		a.CountVowels())
}