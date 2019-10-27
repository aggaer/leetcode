package arr

func Intersect(nums1 []int, nums2 []int) []int {
	var resp []int
	mp := make(map[int]int)
	for _, v := range nums1 {
		mp[v] += 1
	}
	for _, v := range nums2 {
		if i, ok := mp[v]; ok && i > 0 {
			mp[v] -= 1
			resp = append(resp, v)
		}
	}
	return resp
}
