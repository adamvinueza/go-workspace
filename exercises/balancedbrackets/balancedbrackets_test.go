package balancedbrackets

import "strings"
import "testing"

type Test struct {
	Name     string
	Brackets string
	Balanced bool
}

func TestBalancedBrackets(t *testing.T) {
	tests := []Test{
		{
			Name:     "test 0",
			Brackets: "()()()()",
			Balanced: true,
		},
		{
			Name:     "test 1",
			Brackets: "()())()",
			Balanced: false,
		},
		{
			Name:     "test 2",
			Brackets: "()((()()",
			Balanced: false,
		},
		{
			Name:     "test 3",
			Brackets: ")(",
			Balanced: false,
		},
	}

	for _, j := range tests {
		got, remains := IsBalanced(j.Brackets, true)
		if got != j.Balanced {
			t.Errorf("Test %s: expected '%v', got '%v'", j.Name, j.Balanced, got)
		} else {
			t.Logf("Test %v:\n%v\n", j.Name, strings.Join(remains, "\n"))
		}
	}
}
