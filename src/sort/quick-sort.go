package main

func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	cur := l
	for i := l + 1; i < r; i++ {
		if arr[i] < arr[l] {
			cur++
			arr[cur], arr[i] = arr[i], arr[cur]
		}
	}
	arr[l], arr[cur] = arr[cur], arr[l]
	quickSort(arr, l, cur-1)
	quickSort(arr, cur+1, r)
}

func main() {
	arr := []int{2, 7, 4, 6, 9, 1}
	quickSort(arr, 0, len(arr))
	for _, v := range arr {
		println(v)
	}
}
