package demo_leetcode

func firstUniqChar(s string) int {
	auxMap := make(map[int32]int)
	for _, r := range s {
		if auxMap[r] == 0 {
			auxMap[r] = 1
		} else {
			auxMap[r] = auxMap[r] + 1
		}
	}
	for i, r := range s {
		if auxMap[r] == 1 {
			return i
		}
	}
	return -1
}
