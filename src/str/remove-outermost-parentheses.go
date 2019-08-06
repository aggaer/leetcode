package str

func RemoveOuterParentheses(S string) string {
	L, R := 1, 0
	resp := []byte("")
	for i := 1; i < len(S); i++ {
		cur := S[i]
		if cur == '(' {
			L++
		} else {
			R++
		}
		if L != R {
			resp = append(resp, cur)
		} else {
			i++
			L, R = 1, 0
		}
	}
	return string(resp)
}
