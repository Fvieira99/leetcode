package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	count := 0
	leftp := 0
	rightp := len(nums) - 1

	for leftp <= rightp {
		if nums[leftp] != val {
			leftp++
			count++
			continue
		}

		if nums[rightp] != val {
			nums[leftp] = nums[rightp]
			leftp++
			rightp--
			count++
			continue
		}

		rightp--
	}

	fmt.Println(nums)
	return count
}
