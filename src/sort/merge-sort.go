package main

func mergeSort(arr []int, l int, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l int, m int, r int) {
	//拷贝两个小数组
	L := make([]int, m-l+1)
	R := make([]int, r-m)
	for i := 0; i < len(L); i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < len(R); j++ {
		R[j] = arr[m+1+j]
	}
	k := l
	//合并两个小数组
	l1, r1 := 0, 0
	for l1 < len(L) && r1 < len(R) {
		if L[l1] < R[r1] {
			arr[k] = L[l1]
			l1++
		} else {
			arr[k] = R[r1]
			r1++
		}
		k++
	}
	//L，R中剩余的元素加入到原数组
	for l1 < len(L) {
		arr[k] = L[l1]
		l1++
		k++
	}
	for r1 < len(R) {
		arr[k] = R[r1]
		r1++
		k++
	}
}

func main() {
	arr := []int{2, 7, 4, 6, 9, 1}
	mergeSort(arr, 0, len(arr)-1)
	for _, v := range arr {
		println(v)
	}
}
