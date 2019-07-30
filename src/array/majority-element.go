package array

func MajorityElement(nums []int) int {
	count := 0
	resp := nums[0]
	for _, v := range nums {
		if count == 0 {
			resp = v
		}
		if v == resp {
			count++
		} else {
			count--
		}
	}
	return resp
}
