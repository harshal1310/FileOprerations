package counter

import (
	"sort"
)

type wordCount struct {
	WordFreq int
	Word     string
}

func MergerWordCount(wordFreqForAllFiles, wordFreqFile map[string]int) {
	for word, count := range wordFreqFile {
		_, exists := wordFreqForAllFiles[word]
		if exists {
			wordFreqForAllFiles[word] += count
		} else {
			wordFreqForAllFiles[word] = count
		}
	}
}
func PrintTopWords(wordFreqForAllFiles map[string]int, topN int) []wordCount {
	totalUniqueWords := len(wordFreqForAllFiles)
	countWords := make([]wordCount, totalUniqueWords)
	index := 0
	for word, count := range wordFreqForAllFiles {
		countWords[index] = wordCount{Word: word, WordFreq: count}
		index++
	}
	sort.Slice(countWords, func(i, j int) bool {
		return countWords[i].WordFreq > countWords[j].WordFreq
	})
	if len(countWords) < topN {
		return countWords
	}
	return countWords[:topN]
}
