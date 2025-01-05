package main

func main() {
	removeDuplicates([]int{1, 1, 2})
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lastUniqueIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lastUniqueIndex] {
			nums[lastUniqueIndex+1] = nums[i]
			lastUniqueIndex++
		}
	}
	return lastUniqueIndex + 1
}
