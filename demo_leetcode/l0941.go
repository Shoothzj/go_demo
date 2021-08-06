package demo_leetcode

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	leftIndex := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			break
		} else if arr[i] == arr[i+1] {
			return false
		} else {
			leftIndex++
		}
	}
	if leftIndex == 0 {
		return false
	}
	rightIndex := len(arr) - 1
	for i := rightIndex; i > 0; i-- {
		if arr[i] > arr[i-1] {
			break
		} else if arr[i] == arr[i-1] {
			return false
		} else {
			rightIndex--
		}
	}
	if rightIndex == len(arr)-1 {
		return false
	}
	return leftIndex == rightIndex
}
