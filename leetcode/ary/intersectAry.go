package ary

func convertAryToMap(nums []int) map[int]int {
	numByOccurance := make(map[int]int)

	for _, num := range nums {
		if val, ok := numByOccurance[num]; !ok {
			numByOccurance[num] = 1
		} else {
			numByOccurance[num] = 1 + val
		}
	}
	return numByOccurance
}

func returnMinNum(num1 int, num2 int) int {
	if num1 >= num2 {
		return num2
	}
	return num1
}

func Intersect(nums1 []int, nums2 []int) []int {
	var num1ByOccurance, num2ByOccurance map[int]int
	var numsIntersect []int
	num1ByOccurance = convertAryToMap(nums1)
	num2ByOccurance = convertAryToMap(nums2)

	for key, num1Val := range num1ByOccurance {
		if num2Val, ok := num2ByOccurance[key]; ok {
			minNum := returnMinNum(num1Val, num2Val)
			for i := 0; i < minNum; i++ {
				numsIntersect = append(numsIntersect, key)
			}
		}
	}
	return numsIntersect
}
