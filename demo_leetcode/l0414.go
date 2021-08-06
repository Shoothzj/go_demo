package demo_leetcode

import "math"

func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if first == num {
			continue
		}
		if first < num {
			first, num = num, first
		}
		if second == num {
			continue
		}
		if second < num {
			second, num = num, second
		}
		if third == num {
			continue
		}
		if third < num {
			third, num = num, third
		}
	}

	if third == math.MinInt64 {
		return first
	}
	return third
}
