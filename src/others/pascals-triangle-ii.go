package others

func GetRow(rowIndex int) []int {
	var upper []int
	var cur []int
	if rowIndex == 0 {
		cur = append(cur, 1)
		return cur
	}
	upper = GetRow(rowIndex - 1)
	cur = append(cur, 1)
	for i := 0; i < rowIndex-1; i++ {
		cur = append(cur, upper[i]+upper[i+1])
	}
	cur = append(cur, 1)
	return cur
}
