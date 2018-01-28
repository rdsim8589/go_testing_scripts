package ary

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
			}
		}
	}
}
