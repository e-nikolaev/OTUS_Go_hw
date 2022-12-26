package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordCountStruct struct {
	word  string
	count uint
}

func Top10(inputString string) []string {
	inputStringSlice := strings.Fields(inputString)

	// Calculate word counts
	wordCountMap := make(map[string]int, len(inputStringSlice))

	for _, rawWord := range inputStringSlice {
		uniqWord := strings.Trim(strings.ToLower(rawWord), "'.,:;?!-")

		if uniqWord != "" {
			_, found := wordCountMap[uniqWord]
			if found {
				wordCountMap[uniqWord]++
			} else {
				wordCountMap[uniqWord] = 1
			}
		}
	}

	// Create structure from map
	wordCountSlice := []WordCountStruct{}

	for key, val := range wordCountMap {
		wordCountSlice = append(wordCountSlice, WordCountStruct{key, uint(val)})
	}

	// Sort struct by counts (1) and the by words (2)
	sort.Slice(wordCountSlice, func(i, j int) bool {
		if wordCountSlice[i].count != wordCountSlice[j].count {
			return wordCountSlice[i].count > wordCountSlice[j].count
		}
		return wordCountSlice[i].word < wordCountSlice[j].word
	})

	// Calculate result length
	resultLen := 10
	if len(wordCountSlice) < 10 {
		resultLen = len(wordCountSlice)
	}

	// Create slice with result
	resultSlice := make([]string, resultLen)
	for idx := range resultSlice {
		resultSlice[idx] = wordCountSlice[idx].word
	}

	return resultSlice
}
