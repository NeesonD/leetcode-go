package _739_dailyTemperatures

func dailyTemperatures(T []int) []int {
	var stack []int
	dict := make(map[int]int)
	for i := len(T) - 1; i > -1; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= T[i] {
			stack = stack[:len(stack)-1]
		}
		curr := T[i]

		if len(stack) == 0 {
			T[i] = 0
		} else {
			T[i] = dict[stack[len(stack)-1]] - i
		}
		stack = append(stack, curr)
		dict[curr] = i
	}
	return T
}
