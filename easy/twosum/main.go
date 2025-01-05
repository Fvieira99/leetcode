package main

import "fmt"

// import (
// 	"fmt"
// )

func main() {
	twosum([]int{2, 11, 7, 15}, 9)
}

func twosum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, v := range nums {
		diff := target - v

		idx, found := hm[diff]
		if found {
			fmt.Println([]int{i, idx})
			return []int{i, idx}
		}
		hm[v] = i
	}
	return nil
}
