package frequency

import (
	"sort"
	"strings"
)

func Top10(inputStr string) []string {
	if len(inputStr) == 0 {
		return nil
	}

	wordsSequence := strings.Fields(inputStr)

	// Make array with key-value, where key is a word, key is it's frequency
	wordFrequencies := make(map[string]int)
	for i := range wordsSequence {
		wordFrequencies[wordsSequence[i]]++
	}

	// Init string array "keys" with 0-size and capacity of wordFrequencies
	keys := make([]string, 0, len(wordFrequencies))
	// Fill array with keys from wordFrequencies
	for key := range wordFrequencies {
		keys = append(keys, key)
	}

	// Sort words in array "keys" by desc values in Frequencies
	sort.Slice(keys, func(i, j int) bool { return wordFrequencies[keys[i]] > wordFrequencies[keys[j]] })

	// In case there's less than 10 elements
	result := make([]string, 0, 10)
	for i, v := range keys {
		if i > 9 {
			break
		}
		result = append(result, v)
	}

	return result
}
