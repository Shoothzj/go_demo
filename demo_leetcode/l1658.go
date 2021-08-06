package demo_leetcode

import "math"

func minOperations(nums []int, x int) int {
	current := 0
	for _, val := range nums {
		current += val
	}
	length := len(nums)
	result := math.MaxInt32
	left := 0

	for right := 0; right < length; right++ {
		current -= nums[right]
		for current < x && left <= right {
			current += nums[left]
			left += 1
		}
		if current == x {
			aux := (length - 1 - right) + left
			if result > aux {
				result = aux
			}
		}
	}
	if result == math.MaxInt32 {
		return -1
	}
	return result
}
