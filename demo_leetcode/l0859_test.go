package demo_leetcode

import (
	"fmt"
	"testing"
)

func TestBuddyStringsCase1(t *testing.T) {
	result := buddyStrings("ab", "ba")
	fmt.Println(result)
}

func TestBuddyStringsCase2(t *testing.T) {
	result := buddyStrings("abcd", "badc")
	fmt.Println(result)
}

func TestBuddyStringsCase3(t *testing.T) {
	result := buddyStrings("abc", "abc")
	fmt.Println(result)
}
