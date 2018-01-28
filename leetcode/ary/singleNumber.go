package ary

//using the bitwise method
func SingleNumber(nums []int) int {
	valHolder := 0

	for _, val := range nums {
		valHolder ^= val
	}
	return valHolder
}
