package nums

func Fib(N int) int {
	if N <= 1 {
		return N
	}
	a, b := 0, 1
	for i := 2; i <= N; i++ {
		b, a = b+a, b
	}
	return b
}
