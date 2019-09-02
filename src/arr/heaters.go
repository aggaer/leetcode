package arr

import "sort"

func FindRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	resp := 0
	for _, v := range houses {
		left := 0
		right := len(heaters) - 1
		for left < right {
			mid := (left + right) >> 1
			if heaters[mid] > v {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if heaters[left] == v {
			resp = max(resp, 0)
		} else if heaters[left] < v {
			resp = max(resp, v-heaters[left])
		} else {
			if left == 0 {
				resp = max(resp, heaters[left]-v)
			} else {
				resp = max(resp, min(heaters[left]-v, v-heaters[left-1]))
			}
		}
	}
	return resp
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
