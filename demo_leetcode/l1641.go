package demo_leetcode

func countVowelStrings(n int) int {
	if n == 0 {
		return 0
	}
	resultArray := []int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		newArray := make([]int, 5)
		newArray[0] = resultArray[0]
		newArray[1] = resultArray[0] + resultArray[1]
		newArray[2] = resultArray[0] + resultArray[1] + resultArray[2]
		newArray[3] = resultArray[0] + resultArray[1] + resultArray[2] + resultArray[3]
		newArray[4] = resultArray[0] + resultArray[1] + resultArray[2] + resultArray[3] + resultArray[4]
		resultArray = newArray
	}
	result := 0
	for _, index := range resultArray {
		result += index
	}
	return result
}
