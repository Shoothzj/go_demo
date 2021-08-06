package demo_leetcode

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		m := make(map[int32]int)
		for _, char := range s {
			if m[char] == 1 {
				return true
			} else {
				m[char] = 1
			}
		}
		return false
	}
	auxStr, auxGoal := -1, -1
	aux2Str, aux2Goal := -1, -1
	for i := 0; i < len(s); i++ {
		sChar := s[i]
		gChar := goal[i]
		if sChar != gChar {
			if auxStr == -1 {
				auxStr = int(sChar)
				auxGoal = int(gChar)
			} else if aux2Str == -1 {
				aux2Str = int(sChar)
				aux2Goal = int(gChar)
			} else {
				return false
			}
		}
	}
	if auxStr == aux2Goal && auxGoal == aux2Str {
		return true
	}
	return false
}
