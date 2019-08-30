package str

func DefangIPaddr(address string) string {
	resp := []byte("")
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			resp = append(append(append(resp, '['), address[i]), ']')
		} else {
			resp = append(resp, address[i])
		}
	}
	return string(resp)
}
