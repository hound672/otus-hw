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
	cache := make(map[string]int)
	var words = strings.Fields(s)
	for _, s := range words {
		cache[s]++
	}

	var slice = make([]wordsCount, 0, len(cache))
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

	var result = make([]string, resultLen)
	for i := 0; i < resultLen; i++ {
		result[i] = slice[i].Word
	}

	return result
}
