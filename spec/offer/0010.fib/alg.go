package _010_fib

func F() {

}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	pre := 0
	curr := 1
	for i := 2; i <= n; i++ {
		sum := pre + curr
		pre = curr % 1000000007
		curr = sum % 1000000007
	}
	return curr
}
