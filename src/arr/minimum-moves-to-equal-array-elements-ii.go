package arr

import "sort"

func MinMoves2(nums []int) int {
	resp := 0
	sort.Ints(nums)
	mid := len(nums) / 2
	midVal := nums[mid]
	i, j := 0, len(nums)-1
	for i <= mid && j >= mid {
		for nums[i] < midVal || nums[j] > midVal {
			if nums[i] < midVal {
				resp += midVal - nums[i]
				nums[i] = midVal
			}
			if nums[j] > midVal {
				resp += nums[j] - midVal
				nums[j] = midVal
			}
		}
		i++
		j--
	}
	return resp
}
