package main

import (
	"fmt"
	"maps"
	"slices"
)

func areAnagrams(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == "" && t == "" {
		return true
	}

	characterCountMap := make(map[rune]int)

	for _, r := range s {
		characterCountMap[r]++
	}

	for _, r := range t {
		if _, found := characterCountMap[r]; !found {
			return false
		}

		characterCountMap[r]--
	}

	return slices.Max(slices.Collect(maps.Values(characterCountMap))) == 0
}

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	anagramMap[strs[0]] = make([]string, 0)

	for _, str := range strs[1:] {
		foundAnagramPair := false
		for anagramKey, anagramGroup := range anagramMap {
			if areAnagrams(str, anagramKey) {
				anagramMap[anagramKey] = append(anagramGroup, str)
				foundAnagramPair = true
				break
			}
		}

		if !foundAnagramPair {
			anagramMap[str] = make([]string, 0)
		}
	}

	anagramGroups := make([][]string, 0)

	for anagramKey, anagramValueGroup := range anagramMap {
		subGroup := append(anagramValueGroup, anagramKey)
		anagramGroups = append(anagramGroups, subGroup)
	}

	return anagramGroups
}

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Printf("%q => %q\n", strs, groupAnagrams(strs))
	fmt.Printf("%q => %q\n", []string{"x"}, groupAnagrams([]string{"x"}))
	fmt.Printf("%q => %q\n", []string{""}, groupAnagrams([]string{""}))
}
