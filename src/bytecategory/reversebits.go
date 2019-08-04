package bytecategory

func ReverseBits(num uint32) uint32 {
	var result uint32 = 0
	i := 32
	for i > 0 {
		result <<= 1
		result |= num & 1
		num = num >> 1
		i--
	}
	return result
}
