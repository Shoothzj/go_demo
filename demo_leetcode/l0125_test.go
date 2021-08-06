package demo_leetcode

import (
	"fmt"
	"testing"
)

func Test_isPalindromeCase1(t *testing.T) {
	palindrome := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(palindrome)
}

func Test_isPalindromeCase2(t *testing.T) {
	palindrome := isPalindrome("0P")
	fmt.Println(palindrome)
}

func Test_isPalindromeCase3(t *testing.T) {
	palindrome := isPalindrome("Marge, let's \"[went].\" I await {news} telegram.")
	fmt.Println(palindrome)
}
