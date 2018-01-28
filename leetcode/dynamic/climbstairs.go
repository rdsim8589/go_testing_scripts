package dynamic

type dyno struct {
	stepsByPossible map[int]int
}

func (d *dyno) recurseStairs(n int) int {
	var valOne, valTwo int
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if val, ok := d.stepsByPossible[n-1]; !ok {
		valOne = d.recurseStairs(n - 1)
		d.stepsByPossible[n-1] = valOne
	} else {
		valOne = val
	}
	if val, ok := d.stepsByPossible[n-2]; !ok {
		valTwo = d.recurseStairs(n - 2)
		d.stepsByPossible[n-2] = valTwo
	} else {
		valTwo = val
	}
	return valOne + valTwo
}

func ClimbStairs(n int) int {
	d := dyno{stepsByPossible: make(map[int]int)}
	return d.recurseStairs(n)
}
