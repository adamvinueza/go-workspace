package cryptogram

import (
	"fmt"
	"sort"
	"strings"
)

func isSubsequence(candidate string, target string) bool {
	if len(candidate) > 0 {
		idx := strings.Index(target, candidate[0:1])
		if idx > -1 {
			return isSubsequence(candidate[1:], target[idx:])
		} else {
			return false
		}
	} else {
		return true
	}
}

func sortWordsByLength(candidates []string) {
	sort.Slice(candidates, func(i, j int) bool {
		return len(candidates[i]) > len(candidates[j])
	})
}

func GetLongestSubsequence(target string, candidates []string) string {
	sortWordsByLength(candidates)
	fmt.Println("---")
	var longestSubsequence string
	for _, candidate := range candidates {
		if isSubsequence(candidate, target) {
			longestSubsequence = candidate
			break
		}
	}
	return longestSubsequence
}
