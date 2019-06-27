package main

//建立小顶堆
func heapify(arr []int, n int, i int) {
	//假设最大元素为i，将其设为堆顶元素
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int, n int) {
	//构造最大堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		//交换堆顶元素
		arr[0], arr[i] = arr[i], arr[0]
		//恢复最大堆
		heapify(arr, i, 0)
	}
}

func main() {
	arr := []int{2, 7, 4, 6, 9, 1}
	heapSort(arr, len(arr)-1)
	for _, v := range arr {
		println(v)
	}
}
