package others

import "math"

func ConstructRectangle(area int) []int {
	sqrt := int(math.Sqrt(float64(area)))
	for area%sqrt != 0 {
		sqrt--
	}
	return []int{area / sqrt, sqrt}
}
