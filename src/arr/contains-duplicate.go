package arr

func containsDuplicate(nums []int) bool {
	signs := make(map[int]bool)
	for _, v := range nums {
		if signs[v] {
			return true
		}
		signs[v] = true
	}
	return false
}
