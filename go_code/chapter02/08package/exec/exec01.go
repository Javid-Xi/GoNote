package exec

func FibonacciNum(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return FibonacciNum(n-1) + FibonacciNum(n-2)
	}
}
