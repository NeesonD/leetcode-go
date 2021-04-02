package _567_checkInclusion

import "testing"

func TestT(t *testing.T) {
	a := []int{1, 2}
	b := make([]int, len(a))
	copy(b, a)

	checkInclusion("ab", "eidbaooo")
}
