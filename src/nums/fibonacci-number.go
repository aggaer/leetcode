package nums

func Fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return Fib(N-1) + Fib(N-2)
}
