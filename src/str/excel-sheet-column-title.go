package str

func ConvertToTitle(n int) string {
	resp := ""
	for n != 0 {
		k := n % 26
		if k == 0 {
			resp = "Z"+resp
			n -= 26
		} else {
			pre := 'A' + k - 1
			resp = string(pre) + resp
		}
		n /= 26
	}
	return resp
}
