package main

import (
	"fmt"

	"github.com/rdsim8589/go_testing_scripts/leetcode/ary"
	"github.com/rdsim8589/go_testing_scripts/leetcode/bestTimeBuySellStock"
	"github.com/rdsim8589/go_testing_scripts/leetcode/dynamic"
	"github.com/rdsim8589/go_testing_scripts/leetcode/medianTwoSortedAry"
	"github.com/rdsim8589/go_testing_scripts/leetcode/rotateAry"
	"github.com/rdsim8589/go_testing_scripts/leetcode/str"
)

func mainIntersectAry() {
	var nums1, nums2 []int

	nums1 = []int{1, 2, 2, 3, 4, 1}
	nums2 = []int{5, 6}
	fmt.Printf("Expected: [2,2], Got: %v\n", ary.Intersect(nums1, nums2))
}

func mainReverseInt() {

	fmt.Printf("Expected: 321, Got: %d\n", str.Reverse(123))

}
func mainSingleNumber() {
	var nums []int

	nums = []int{1, 4, 5, 2, 4, 5, 1}
	fmt.Printf("Expected: 2, Got: %d\n", ary.SingleNumber(nums))

}

func mainClimbStairs() {
	fmt.Printf("Expected: 5, Got: %d\n", dynamic.ClimbStairs(4))
	fmt.Printf("Expected: ?, Got: %d\n", dynamic.ClimbStairs(5))
	fmt.Printf("Expected: 1, Got: %d\n", dynamic.ClimbStairs(1))
	fmt.Printf("Expected: 0, Got: %d\n", dynamic.ClimbStairs(0))

}
func mainReverseString() {
	s := "hello"
	fmt.Printf("Expected: 'olleh', Got: %s\n", str.ReverseString(s))

}

func mainContainsDuplicate() {
	var num []int
	num = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Expected: False , Got: %t\n", ary.ContainsDuplicate(num))

	num = []int{1, 2, 3, 4, 5, 1, 7}
	fmt.Printf("Expected: True , Got: %t\n", ary.ContainsDuplicate(num))

}
func mainRotateAry() {
	var ary []int

	// ary = []int{1, 2, 3, 4, 5, 6}
	// rotateArray.Rotate(ary, 5)
	// fmt.Printf("Expected: [2,3,4,5,6,1], Got: %v\n", ary)

	// ary = []int{1, 2, 3, 4, 5, 6}
	// rotateArray.Rotate(ary, 2)
	// fmt.Printf("Expected: [5,6,1,2,3,4], Got: %v\n", ary)

	ary = []int{1, 2, 3, 4, 5, 6}
	rotateAry.RotateTwo(ary, 2)
	fmt.Printf("Expected: [5,6,1,2,3,4], Got: %v\n", ary)

	ary = []int{1, 2}
	rotateAry.Rotate(ary, 2)
	fmt.Printf("Expected: [5,6,1,2,3,4], Got: %v\n", ary)
}

func mainbestSellBuyStocks() {
	var stocks []int

	stocks = []int{1, 2, 3, 4, 5}
	fmt.Printf("Expected: 4, Got: %d\n", bestSellBuyStock.MaxProfit(stocks))

	stocks = []int{5, 4, 3, 2, 1}
	fmt.Printf("Expected: 0, Got: %d\n", bestSellBuyStock.MaxProfit(stocks))

	stocks = []int{1, 2, 6, 4, 5}
	fmt.Printf("Expected: 6, Got: %d\n", bestSellBuyStock.MaxProfit(stocks))

	stocks = []int{1, 1, 1, 1, 1}
	fmt.Printf("Expected: 0, Got: %d\n", bestSellBuyStock.MaxProfit(stocks))

	stocks = []int{1}
	fmt.Printf("Expected: 0, Got: %d\n", bestSellBuyStock.MaxProfit(stocks))
}

func mainMedTwoSortedAry() {
	var num1, num2 []int

	num1 = []int{1, 4, 6}
	num2 = []int{2, 3}
	fmt.Printf("Expected:%f, Got:%f\n", 3.0, medianTwoSortedAry.FindMedianSortedArrays(num1, num2))

	num1 = []int{1}
	num2 = []int{1}
	fmt.Printf("Expected:%f, Got:%f\n", 1.0, medianTwoSortedAry.FindMedianSortedArrays(num1, num2))

	num1 = []int{1}
	num2 = []int{}
	fmt.Printf("Expected:%f, Got:%f\n", 1.0, medianTwoSortedAry.FindMedianSortedArrays(num1, num2))

	num1 = []int{1, 4, 6}
	num2 = []int{2, 3}
	fmt.Printf("Expected:%f, Got:%f\n", 3.0, medianTwoSortedAry.FindMedianSortedArrays(num1, num2))
}

// func mainHeaters() {
// 	var houses, heater []int
//
// 	houses = []int{1, 2, 3, 4, 5, 6, 7}
// 	heater = []int{1}
// 	fmt.Printf("houses: %v, heater:%v, expected: %d, got: %d\n", houses, heater, 6, heaters.FindRadius(houses, heater))
//
// 	houses = []int{1, 2, 3, 4, 5, 6, 7}
// 	heater = []int{3, 6}
// 	fmt.Printf("houses: %v, heater:%v, expected: %d, got: %d\n", houses, heater, 2, heaters.FindRadius(houses, heater))
//
// 	houses = []int{1, 2, 3, 4, 5, 6}
// 	heater = []int{3, 4}
// 	fmt.Printf("houses: %v, heater:%v, expected: %d, got: %d\n", houses, heater, 2, heaters.FindRadius(houses, heater))
//
// 	houses = []int{1, 2, 3, 4, 5, 6}
// 	heater = []int{2, 3, 4, 5}
// 	fmt.Printf("houses: %v, heater:%v, expected: %d, got: %d\n", houses, heater, 1, heaters.FindRadius(houses, heater))
//
// 	houses = []int{1, 2, 3, 5, 15}
// 	heater = []int{2, 30}
// 	fmt.Printf("houses: %v, heater:%v, expected: %d, got: %d\n", houses, heater, 13, heaters.FindRadius(houses, heater))
//
// }

func main() {
	mainIntersectAry()
}
