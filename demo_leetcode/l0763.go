package demo_leetcode

// partitionLabels
func partitionLabels(s string) []int {
	var lastArray [26]int
	for i, char := range s {
		lastArray[char-97] = i
	}
	var j = 0
	var anchor = 0
	var result []int
	for i, char := range s {
		if lastArray[char-97] > j {
			j = lastArray[char-97]
		}
		if i == j {
			result = append(result, i-anchor+1)
			anchor = i + 1
		}
	}
	return result
}
