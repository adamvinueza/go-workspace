package recursion

import "testing"

type Test struct {
	Name     string
	Array    []int
	Expected int
}

func TestGetLength(t *testing.T) {
	tests := []Test{
		{
			Name:     "primes",
			Array:    []int{2, 3, 5, 7, 11, 13, 17},
			Expected: 7,
		},
		{
			Name:     "negatives",
			Array:    []int{-1, -2, -3},
			Expected: 3,
		},
	}

	for _, j := range tests {
		actual := GetLength(j.Array)
		if actual != j.Expected {
			t.Errorf("Expected %d, got %d", j.Expected, actual)
		}
	}
}
