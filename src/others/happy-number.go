package others

func IsHappy(n int) bool {
	m := make(map[int]bool)
	for {
		m[n] = true
		cur := 0
		for n != 0 {
			cur += pow(n % 10)
			n /= 10
		}
		n = cur
		if n == 1 {
			return true
		}
		if _, ok := m[n]; ok {
			return false
		}
	}
}

func pow(a int) int {
	return a * a
}
