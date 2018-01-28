package misc

func strToSet(s string) map[string]bool {
	set := make(map[string]bool)
	for _, char := range s {
		set[string(char)] = true
	}
	return set
}
func NumJewelsInStones(J string, S string) int {
	var set map[string]bool

	countJewel := 0
	set = strToSet(J)
	for _, scaned := range S {
		if set[string(scaned)] == true {
			countJewel++
		}
	}
	return countJewel
}
