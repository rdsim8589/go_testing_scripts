package ary

func ContainsDuplicate(nums []int) bool {
	numByOccurance := make(map[int]int)
	for _, num := range nums {
		if _, ok := numByOccurance[num]; ok {
			return true
		}
		numByOccurance[num] = 1
	}
	return false
}
