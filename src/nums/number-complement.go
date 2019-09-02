package nums

func FindComplement(num int) int {
	shift := 1
	for shift < num {
		shift <<= 1
		shift++
	}
	return shift ^ num
}
