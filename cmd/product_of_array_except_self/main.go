package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for prefixProduct, i := 1, 0; i < len(nums); i++ {
		result[i] = prefixProduct
		prefixProduct *= nums[i]
	}

	for suffixProduct, i := 1, len(nums)-1; i >= 0; i-- {
		result[i] *= suffixProduct
		suffixProduct *= nums[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 4, 6}
	fmt.Printf("nums=%v => solution=%v\n", nums, productExceptSelf(nums))

	nums = []int{-1, 0, 1, 2, 3}
	fmt.Printf("nums=%v => solution=%v\n", nums, productExceptSelf(nums))
}
