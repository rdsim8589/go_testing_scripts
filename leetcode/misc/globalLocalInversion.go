package misc

func IsIdealPermutationBrute(A []int) bool {
	globalCount, localCount := 0, 0
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			localCount++
		}
		for j := i; j < len(A); j++ {
			if A[i] > A[j] {
				globalCount++
			}
			if globalCount > localCount {
				return false
			}
		}
	}
	return true
}

func IsIdealPermutation(A []int) bool {
	//loop through A
	maxVal := A[0]
	for i := 0; i < len(A)-2; i++ {
		if maxVal < A[i] {
			maxVal = A[i]
		}
		if maxVal > A[i+2] {
			return false
		}
	}
	return true

}
