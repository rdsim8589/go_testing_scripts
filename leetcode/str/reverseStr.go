package str

func ReverseString(s string) string {
	newStr := ""
	for _, char := range s {
		newStr = string(char) + newStr
	}
	return newStr
}
