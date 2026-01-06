package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)

	for index_one, num := range nums {
		if pair_index, found := sumMap[target-num]; found {
			if index_one < pair_index {
				return []int{index_one, pair_index}
			} else {
				return []int{pair_index, index_one}
			}
		} else {
			sumMap[num] = index_one
		}
	}

	return nil
}

func main() {
	nums := []int{3, 4, 5, 6}
	target := 7

	fmt.Printf("%v target=%d => %v\n", nums, target, twoSum(nums, target))
}
