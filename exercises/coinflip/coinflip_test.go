package coinflip

import "testing"

type Test struct {
	Coins    int
	Expected int
}

func TestExpectedFlips(t *testing.T) {
	tests := []Test{
		{
			Coins:    0,
			Expected: 0,
		},
		{
			Coins:    1,
			Expected: 0,
		},
		{
			Coins:    2,
			Expected: 1,
		},
		{
			Coins:    3,
			Expected: 2,
		},
		{
			Coins:    4,
			Expected: 2,
		},
		{
			Coins:    100,
			Expected: 7,
		},
		{
			Coins:    412,
			Expected: 9,
		},
	}

	for _, j := range tests {
		actual := ExpectedFlips(j.Coins)
		if actual != j.Expected {
			t.Errorf("Expected %d, actual %d", j.Expected, actual)
		}
	}
}
