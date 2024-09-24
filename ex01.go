package tp

import (
	"sort"
)

func Ft_coin(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	count := 0

	for _, coin := range coins {
		for amount >= coin {
			amount -= coin
			count++
		}
	}
	if amount > 0 {
		return -1
	}
	return count
}
