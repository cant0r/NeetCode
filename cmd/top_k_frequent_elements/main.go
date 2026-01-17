package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	freqGroups := make(map[int][]int)

	if k > len(nums) {
		return nil
	}

	for _, num := range nums {
		freqMap[num]++
	}

	for number, frequency := range freqMap {
		freqGroups[frequency] = append(freqGroups[frequency], number)
	}

	topKGroups := make([]int, 0)

	for i := len(nums); i > 0; i-- {
		if numbers := freqGroups[i]; len(numbers) > 0 {
			topKGroups = append(topKGroups, numbers...)
		}

		if len(topKGroups) == k {
			break
		}
	}
	return topKGroups

}

func main() {
	fmt.Printf("nums=%v k=%v => %v\n", []int{1, 2, 2, 3, 3, 3}, 2, topKFrequent([]int{1, 2, 2, 3, 3, 3}, 2))
	fmt.Printf("nums=%v k=%v => %v\n", []int{7}, 1, topKFrequent([]int{7}, 1))
	fmt.Printf("nums=%v k=%v => %v\n", []int{7, 7, 8, 8}, 2, topKFrequent([]int{7, 7, 8, 8}, 2))
	fmt.Printf("nums=%v k=%v => %v\n", []int{7, 7, 8, 8}, 3, topKFrequent([]int{7, 7, 8, 8}, 3))
	fmt.Printf("nums=%v k=%v => %v\n", []int{7}, 1, topKFrequent([]int{7}, 1))
	fmt.Printf("nums=%v k=%v => %v\n", []int{}, 0, topKFrequent([]int{}, 0))
}
