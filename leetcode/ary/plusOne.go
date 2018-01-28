package ary

func plusOne(digits []int) []int {
	newDigits := make([]int, len(digits)+1)
	carry := 1 + digits[len(digits)-1]
	for i := len(digits) - 2; i >= 0; i-- {
		newDigits[i+2] = carry % 10
		carry = digits[i] + (carry / 10)
	}
	newDigits[1] = carry % 10
	carry = (carry / 10)
	if carry == 0 {
		return newDigits[1:len(newDigits)]
	}
	newDigits[0] = carry
	return newDigits
}
