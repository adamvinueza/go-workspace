package cryptogram

import "testing"

type Test struct {
	Name               string
	Words              []string
	Target             string
	LongestSubsequence string
}

func TestGetLongestSubsequence(t *testing.T) {
	tests := []Test{
		{
			Name: "original test set",
			Words: []string{
				"apple",
				"able",
				"ale",
				"bale",
				"aple",
				"kangaroo",
			},
			Target:             "abppplee",
			LongestSubsequence: "apple",
		},
		{
			Name:               "no words in set",
			Words:              []string{},
			Target:             "apple",
			LongestSubsequence: "",
		},
		{
			Name: "Longest word in set is longest, include empty string",
			Words: []string{
				"abab",
				"babble",
				"abracadabra",
				"",
			},
			Target:             "abppapbracadablraee",
			LongestSubsequence: "abracadabra",
		},
		{
			Name:               "All words are empty",
			Words:              []string{"", "", "", ""},
			Target:             "apple",
			LongestSubsequence: "",
		},
	}

	for _, j := range tests {
		got := GetLongestSubsequence(j.Target, j.Words)
		if got != j.LongestSubsequence {
			t.Errorf("Expected '%s', got '%s'",
				j.LongestSubsequence, got)
		}
	}
}
