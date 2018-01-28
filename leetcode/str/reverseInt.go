package str

import (
	"strconv"
)

func Reverse(x int) int {
	var revIntStr string
	for x > 0 {
		strInt := strconv.Itoa(x % 10)
		revIntStr += strInt
		x /= 10
	}
	revInt, err := strconv.Atoi(revIntStr)
	if err != nil {
		return 0
	}
	return revInt
}
