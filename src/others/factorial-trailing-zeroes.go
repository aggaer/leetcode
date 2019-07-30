package others

func TrailingZeroes(n int) int {
	resp := 0
	for n > 0 {
		resp += n / 5
		n /= 5
	}
	return resp
}
