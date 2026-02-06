package utils

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
