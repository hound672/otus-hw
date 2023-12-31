package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const (
	maxResultLen = 10
)

type wordsCount struct {
	Word string
	Freq int
}

func Top10(s string) []string {
	words := strings.Fields(s)

	cache := make(map[string]int)
	for _, s := range words {
		cache[s]++
	}

	slice := make([]wordsCount, 0, len(cache))
	for k, v := range cache {
		slice = append(slice, wordsCount{k, v})
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].Freq == slice[j].Freq {
			return slice[i].Word < slice[j].Word
		}
		return slice[i].Freq > slice[j].Freq
	})

	resultLen := len(slice)
	if resultLen > maxResultLen {
		resultLen = maxResultLen
	}

	result := make([]string, resultLen)
	for i := 0; i < resultLen; i++ {
		result[i] = slice[i].Word
	}

	return result
}
