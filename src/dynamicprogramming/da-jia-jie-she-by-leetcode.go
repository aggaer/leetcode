package dynamicprogramming

//动态规划，f(k) = max(f(k-2)+Ak,f(k-1))
func Rob(nums []int) int {
	preMax, curMax := 0, 0
	for _, v := range nums {
		temp := curMax
		curMax = max(preMax+v, curMax)
		preMax = temp
	}
	return curMax
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
