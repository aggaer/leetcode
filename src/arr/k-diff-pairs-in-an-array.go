package arr

func FindPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	records := make(map[int]int)
	diff := make(map[int]int)
	for i, v := range nums {
		if _, ok := records[v-k]; ok {
			diff[v] = i
		}
		if _, ok := records[v+k]; ok {
			diff[v+k] = i
		}
		records[v] = i
	}
	return len(diff)
}
