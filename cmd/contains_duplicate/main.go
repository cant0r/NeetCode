package main

import "fmt"

func hasDuplicate(nums []int) bool {
	intSet := make(map[int]struct{})

	for _, num := range nums {
		if _, found := intSet[num]; found {
			return true
		}

		intSet[num] = struct{}{}
	}

	return false
}

func main() {
	nums_one := []int{1, 2, 3, 3}
	nums_two := []int{1, 2, 3, 4}

	fmt.Printf("%v => %v\n", nums_one, hasDuplicate(nums_one))
	fmt.Printf("%v => %v\n", nums_two, hasDuplicate(nums_two))
}
