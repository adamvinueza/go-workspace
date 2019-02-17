// Package recursion_exercise demonstrates recursive functions.
package recursion

func init() {}

// Standard recursion.
func GetLength(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + GetLength(arr[1:])
}

// Uses tail recursion. Note that the lower-case 'g' means this function won't
// be exported, so won't be visible from outside this package.
func getLengthTail(arr []int, accum int) int {
	if len(arr) == 0 {
		return accum
	}
	return getLengthTail(arr[1:], 1+accum)
}

// Optimized: uses tail recursion, wrapped for simplicity.
func GetLengthOptimized(arr []int) int {
	return getLengthTail(arr, 0)
}
