package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashSet := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		expectPair := target - num
		if idx, ok := hashSet[expectPair]; ok {
			fmt.Printf("Numbers %d, %d in array mun %d", num, expectPair, target)
			foundIdx := []int{i, idx}
			return foundIdx
		}
		hashSet[num] = i
	}
	return nil
}

func main() {
	nums := []int{3, 5, 9, 6, -1}
	target := 5
	fmt.Println(twoSum(nums, target))
}
