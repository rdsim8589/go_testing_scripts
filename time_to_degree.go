package main

import (
	"fmt"
	"math"
)

func time_to_degree(hour float64, min float64) {
//	var hour_ratio, min_ratio angle_diff float64
	hour_ratio := float64(360/12)
	min_ratio := float64(360/60)
	angle_diff := math.Abs((hour_ratio * hour) - (min_ratio * min))
	fmt.Printf("hour: %d, min: %f, angle diff %f\n", hour, min, angle_diff)
}
func main() {
	time_to_degree(6, 30)
	time_to_degree(3, 45)
	time_to_degree(0, 0)
	time_to_degree(2, 27)
}
