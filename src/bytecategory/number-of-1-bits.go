package bytecategory

func HammingWeight(num uint32) int {
	i := 32
	count := 0
	for i > 0 {
		if temp := num & 1; temp == 1 {
			count++
		}
		num = num >> 1
		i--
	}
	return count
}

func HammingWeight2(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}
