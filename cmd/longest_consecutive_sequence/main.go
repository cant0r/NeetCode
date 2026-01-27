package main

import (
	"fmt"
	"slices"
)

func getChainDepth(key int, sequenceMap map[int]int) int {
	element, found := sequenceMap[key]

	if key == element && found {
		return 1
	}

	return 1 + getChainDepth(element, sequenceMap)
}

func longestConsecutive(nums []int) int {
	uniqueNumbers := make(map[int]struct{})

	for _, num := range nums {
		uniqueNumbers[num] = struct{}{}
	}

	longestSequence := 0

	for num := range uniqueNumbers {
		currentSequence := 1

		for {
			if _, foundNext := uniqueNumbers[num+currentSequence]; foundNext {
				currentSequence++
			} else {
				break
			}
		}

		if currentSequence > longestSequence {
			longestSequence = currentSequence
		}
	}

	return longestSequence
}

func longestConsecutiveLame(nums []int) int {
	sequenceMap := make(map[int]int) // key: number, value: parent (number-1)

	if len(nums) == 0 {
		return 0
	}

	for _, num := range nums {
		// Number tracked
		if _, found := sequenceMap[num]; found {
			continue
		}

		assignedParent := false

		// Check for parent
		if _, parentFound := sequenceMap[num-1]; parentFound {
			sequenceMap[num] = num - 1
			assignedParent = true
		}

		// Check for reverse parent
		if _, reverseParentFound := sequenceMap[num+1]; reverseParentFound {
			sequenceMap[num+1] = num
			if !assignedParent {
				sequenceMap[num] = num
			}
		} else if !assignedParent {
			sequenceMap[num] = num
		}
	}

	depths := make([]int, 0)

	for key, _ := range sequenceMap {
		depths = append(depths, getChainDepth(key, sequenceMap))
	}

	return slices.Max(depths)
}

func main() {
	nums := []int{2, 20, 4, 10, 3, 4, 5}
	fmt.Printf("nums=%v longestConsecutive=%d\n", nums, longestConsecutive(nums))
	fmt.Printf("nums=%v longestConsecutive=%d\n", nums, longestConsecutiveLame(nums))
	nums = []int{0, 3, 2, 5, 4, 6, 1, 1}
	fmt.Printf("nums=%v longestConsecutive=%d\n", nums, longestConsecutive(nums))
	fmt.Printf("nums=%v longestConsecutive=%d\n", nums, longestConsecutiveLame(nums))
}
