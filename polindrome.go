package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	var nl int
	s := strconv.Itoa(x)
	l := len(s)
	if l%2 == 0 {
		nl = l / 2
	} else {
		nl = (l - 1) / 2
	}

	for i, r := range s {
		if i >= nl {
			break
		}
		if string(s[l-1-i]) != string(r) {
			return false
		}
	}

	return true
}

func main() {
	x := 1221
	fmt.Println(isPalindrome(x))
}
