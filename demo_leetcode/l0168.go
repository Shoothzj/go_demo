package demo_leetcode

import "strings"

func convertToTitle(columnNumber int) string {
	var sb strings.Builder
	for ; columnNumber != 0; {
		i := columnNumber % 26
		if i > 0 {
			sb.WriteByte(byte(i + 65 - 1))
			columnNumber = columnNumber/26
		} else {
			sb.WriteString("Z")
			columnNumber = columnNumber/26 - 1
		}
	}
	return Reverse(sb.String())
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
