package _017_printNumbers

import "math"

func printNumbers(n int) []int {
	s := math.Pow(10, float64(n))
	result := make([]int, int(s)-1)
	for i := 0; i < int(s)-1; i++ {
		result[i] = i + 1
	}
	return result
}
