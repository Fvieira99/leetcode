package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(-121))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	revert := ""

	for i := len(str) - 1; i >= 0; i-- {
		revert += string(str[i])
	}

	return revert == str
}

// BEST SOLUTION 0MS
// func isPalindrome(x int) bool {
// 	s := strconv.Itoa(x)

// 	li := 0
// 	ri := len(s) - 1

// 	for li < len(s)/2 {
// 		if s[li] != s[ri] {
// 			return false
// 		}

// 		li++
// 		ri--
// 	}

// 	return true
// }
