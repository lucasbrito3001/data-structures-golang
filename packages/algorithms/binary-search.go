package binarysearch

import (
	"math"
	"slices"
)

func BinarySearch(list []int, target int) int {
	slices.Sort(list)

	left := 0
	right := len(list) - 1
	mid := len(list) / 2

	for right >= 1 {
		mid = left + int(math.Floor(float64((right-left)/2)))

		if list[mid] == target {
			break
		} else if list[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return mid
}
