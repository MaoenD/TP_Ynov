package tp

import (
	"sort"
)

func Ft_missing(nums []int) int {
	n := len(nums)
	
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}

	return n
}
