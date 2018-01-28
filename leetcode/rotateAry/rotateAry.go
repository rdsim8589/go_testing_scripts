package rotateAry

import (
	"fmt"
)

func Rotate(nums []int, k int) {
	if k%len(nums) == 0 {
		return
	}
	checkIdx := 0
	idx := 0
	currentVal := nums[0]
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums)
		nextIdx := (idx + k) % len(nums)
		nextVal := nums[nextIdx]
		fmt.Printf("nextIdx: %d, nextVal: %d, idx: %d, currentVal: %d\n\n", nextIdx, nextVal, idx, currentVal)
		nums[nextIdx] = currentVal
		idx = nextIdx
		currentVal = nextVal
		if idx == checkIdx {
			idx++
			checkIdx++
			currentVal = nums[nextIdx+1]
		}

	}
	fmt.Println(nums)
}

func RotateTwo(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for startIdx := 0; count < len(nums); startIdx++ {
		prevVal := nums[startIdx]
		nextIdx := (startIdx + k) % len(nums)
		currentVal := nums[nextIdx]
		for nextIdx != startIdx {
			nums[nextIdx] = prevVal
			fmt.Println(nums)
			fmt.Printf("prevVal: %d, nextIdx: %d, currentVal: %d\n\n", prevVal, nextIdx, currentVal)

			nextIdx = (nextIdx + k) % len(nums)
			fmt.Println("nextIdx", nextIdx)
			prevVal = currentVal
			currentVal = nums[nextIdx]
			count++
		}
		fmt.Println("next idx", nextIdx)
		nums[nextIdx] = prevVal
		fmt.Println(nums)
		count++
	}

}
