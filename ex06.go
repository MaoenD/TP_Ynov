package tp

import (
	"math"
)

func Ft_min_window(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}
	targetCount := make(map[rune]int)
	for _, char := range t {
		targetCount[char]++
	}
	left, right := 0, 0
	minLength := math.MaxInt32
	minLeft := 0
	formed := 0
	windowCount := make(map[rune]int)
	for right < len(s) {
		char := rune(s[right])
		windowCount[char]++
		if count, found := targetCount[char]; found && windowCount[char] == count {
			formed++
		}
		for left <= right && formed == len(targetCount) {
			if right-left+1 < minLength {
				minLength = right - left + 1
				minLeft = left
			}
			leftChar := rune(s[left])
			windowCount[leftChar]--
			if count, found := targetCount[leftChar]; found && windowCount[leftChar] < count {
				formed--
			}
			left++
		}
		right++
	}

	if minLength == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minLeft+minLength]
}
