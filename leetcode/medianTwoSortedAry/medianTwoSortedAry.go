package medianTwoSortedAry

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var medianValOne, medianValTwo int
	var median float64
	totalLen := len(nums1) + len(nums2)
	j, k := 0, 0
	for i := 0; i <= (totalLen/2 + 1); i++ {
		if i == totalLen/2-1 {
			if nums1[j] > nums2[k] {
				medianValOne = nums2[k]
			} else {
				medianValOne = nums1[j]
			}
		} else if i == (totalLen / 2) {
			if nums1[j] > nums2[k] {
				medianValTwo = nums2[k]
			} else {
				medianValTwo = nums1[j]
			}
			break
		}
		fmt.Printf("total idx counter: %d, median idx: %d, median idx: %d\n", i, totalLen/2-1, totalLen/2)
		fmt.Printf("num1 idx: %d, num1 len: %d, num2 idx: %d, num2 len: %d, totalLen: %d \n", j, len(nums1), k, len(nums2), totalLen)
		if j < len(nums1)-1 {
			if nums1[j] <= nums2[k] {
				j++
				continue
			}
		}
		if k < len(nums2)-1 {
			if nums1[j] > nums2[k] {
				k++
				continue
			}
		}

	}

	if totalLen%2 == 1 {
		fmt.Println("odd")
		median = float64(medianValTwo)
	} else {
		fmt.Println("even")
		median = float64(medianValOne+medianValTwo) / 2.0
	}
	fmt.Println(totalLen, median, medianValOne, medianValTwo)
	return median
}
