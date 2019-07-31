package str

func ReverseString(s []byte) {
	for i, j := 0, len(s); i < len(s)/2; i++ {
		j--
		s[i], s[j] = s[j], s[i]
	}
}
