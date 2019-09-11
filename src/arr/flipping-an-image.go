package arr

func flipAndInvertImage(A [][]int) [][]int {
	for _, v := range A {
		for i := 0; i < len(v)/2; i++ {
			j := len(v) - 1 - i
			v[i], v[j] = v[j], v[i]
			v[i] ^= 1
			v[j] ^= 1
		}
		if isodd := len(v) % 2; isodd == 0 {
			v[len(v)/2] ^= 1
		}
	}
	return A
}
