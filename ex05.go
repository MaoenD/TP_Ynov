package tp

func Ft_max_substring(s string) int {
	charIndex := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, char := range s {
		if lastIndex, found := charIndex[char]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		charIndex[char] = i
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}
