package coinflip

import "math"

func expectedFlipsTail(n int, accum int) int {
	if n > 1 {
		return expectedFlipsTail(int(math.Ceil(float64(n)/2)), accum+1)
	} else {
		return accum
	}
}

func expectedFlipsWithBitShift(n int, accum int) int {
	if n > 1 {
		shifted := n >> 1
		next_number := shifted
		if n&1 > 0 {
			next_number = 1 + shifted
		}
		return expectedFlipsWithBitShift(next_number, accum+1)
	} else {
		return accum
	}
}

func ExpectedFlips(n int) int {
	if n > 1 {
		return 1 + ExpectedFlips(int(math.Ceil(float64(n)/2)))
	} else {
		return 0
	}
}

func ExpectedFlipsWithBitShiftOptimized(n int) int {
	return expectedFlipsWithBitShift(n, 0)
}

func ExpectedFlipsOptimized(n int) int {
	return expectedFlipsTail(n, 0)
}
