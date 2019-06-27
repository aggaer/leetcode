package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{2, 7, 4, 6, 9, 1}
	bubbleSort(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}
