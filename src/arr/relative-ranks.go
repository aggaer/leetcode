package arr

import (
	"sort"
	"strconv"
)

func FindRelativeRanks(nums []int) []string {
	mp := make(map[int]int)
	for k, v := range nums {
		mp[v] = k
	}
	sort.Ints(nums)
	resp := make([]string, len(nums))
	order := 1
	for i := len(nums) - 1; i >= 0; i-- {
		key := nums[i]
		switch order {
		case 1:
			resp[mp[key]] = "Gold Medal"
		case 2:
			resp[mp[key]] = "Silver Medal"
		case 3:
			resp[mp[key]] = "Bronze Medal"
		default:
			resp[mp[key]] = strconv.Itoa(order)
		}
		order++
	}
	return resp
}
