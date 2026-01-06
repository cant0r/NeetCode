package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	characterCountMap := make(map[rune]int)

	for _, r := range s {
		characterCountMap[r]++
	}

	for _, r := range t {
		if _, found := characterCountMap[r]; found {
			characterCountMap[r]--
		} else {
			return false
		}
	}

	for _, count := range characterCountMap {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	s, t := "racecar", "carrace"
	fmt.Printf("%s %s => %v\n", s, t, isAnagram(s, t))

	s, t = "jar", "jem"
	fmt.Printf("%s %s => %v\n", s, t, isAnagram(s, t))
}
