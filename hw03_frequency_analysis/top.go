package frequency

import (
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func Top10(inputStr string) []string {
	if len(inputStr) == 0 {
		return []string{}
	}

	wordsSequence := strings.Fields(inputStr)

	wordFrequencies := make(map[string]int)
	for i := range wordsSequence {
		wordFrequencies[wordsSequence[i]]++
	}
	
	wordCount := PairList{}
	for key, value := range wordFrequencies {
		wordCount = append(wordCount, Pair{key, value})
	}

	sort.Slice(wordCount, func(i, j int) bool { return wordCount[i].Value > wordCount[j].Value })

	var s []string
	for i := 0; i < 10; i++ {
		s = append(s, wordCount[i].Key)
	}

	return s
}
