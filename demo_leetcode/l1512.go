package demo_leetcode

func numIdenticalPairs(nums []int) int {
	auxMap := make(map[int]int)
	for _, r := range nums {
		if auxMap[r] == 0 {
			auxMap[r] = 1
		} else {
			auxMap[r] = auxMap[r] + 1
		}
	}
	result := 0
	for _, val := range auxMap {
		if val != 1 {
			result += (val * (val - 1)) / 2
		}
	}
	return result
}
