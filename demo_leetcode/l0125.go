package demo_leetcode

func isPalindrome(s string) bool {
	result := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] < 48+10 {
			result = append(result, s[i])
		}
		if s[i] >= 65 && s[i] < 65+26 {
			result = append(result, s[i]+32)
		}
		if s[i] >= 97 && s[i] < 97+26 {
			result = append(result, s[i])
		}
	}
	return isPalindromeAux(string(result))
}

func isPalindromeAux(src string) bool {
	if src == "" {
		return true
	}
	l := len(src)
	for i := 0; i < l/2; i++ {
		if src[i] != src[l-i-1] {
			return false
		}
	}
	return true
}
