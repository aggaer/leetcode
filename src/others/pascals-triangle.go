package others

import "fmt"

func generate(numRows int) [][]int {
	var target [][]int
	for i := 0; i < numRows; i++ {
		if i == 0 {
			target = append(target, []int{1})
			continue
		}
		target = append(target, filling(target[i-1]))
	}
	return target
}

func filling(upper []int) []int {
	if len(upper) == 1 {
		return []int{1, 1}
	}
	var current []int
	current = append(current, 1)
	for i := 0; i < len(upper)-1; i++ {
		current = append(current, upper[i]+upper[i+1])
	}
	return append(current, 1)
}

func main() {
	target := generate(8)
	for _, v := range target {
		for _, v := range v {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
