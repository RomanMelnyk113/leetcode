package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	p := 1221
	isP := isPalindrome(p)
	if isP != true {
		t.Errorf("%d is not palindrome", p)
	}
}

func TestIsPalindromeLinked(t *testing.T) {
	p4 := ListNode{1, nil}
	p3 := ListNode{2, &p4}
	p2 := ListNode{2, &p3}
	p1 := ListNode{1, &p2}
	isP := isPalindromeLinked(&p1)
	if isP != true {
		t.Errorf("is not palindrome")
	}
}
