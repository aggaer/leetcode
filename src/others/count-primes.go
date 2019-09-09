package others

func countPrimes(n int) int {
	a := make([]bool, n)
	cnt := 0
	for i := 2; i < n; i++ {
		if a[i] {
			continue
		}
		for j := i + i; j < n; j += i {
			a[j] = true
		}
		cnt++
	}
	return cnt
}
