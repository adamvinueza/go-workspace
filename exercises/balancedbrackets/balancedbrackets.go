package balancedbrackets

import "container/list"
import "fmt"

type IndexedChar struct {
	Char  string
	Index int
}

func IsBalanced(s string, trace bool) (bool, []string) {
	matchingStack := list.New()
	var remains []string
	for i, r := range s {
		current := IndexedChar{string(r), i}
		lastItem := matchingStack.Front()
		if lastItem == nil {
			matchingStack.PushFront(current)
		} else {
			last := lastItem.Value.(IndexedChar)
			if isMatched(current, last) {
				matchingStack.Remove(lastItem)
			} else {
				if trace {
					matchingStack.PushFront(current)
				} else {
					return false, remains
				}
			}
		}
	}
	if matchingStack.Len() > 0 {
		remains = make([]string, 0)
		var item IndexedChar
		for e := matchingStack.Back(); e != nil; e = e.Prev() {
			item = e.Value.(IndexedChar)
			remains = append(remains, fmt.Sprintf("'%v' dangles '%v' at index %v", s, item.Char, item.Index))
		}
	}
	return matchingStack.Len() == 0, remains
}

func isMatched(first IndexedChar, second IndexedChar) bool {
	matched := false
	if second.Char == "}" {
		matched = first.Char == "{"
	} else if second.Char == "[" {
		matched = first.Char == "]"
	} else if second.Char == "(" {
		matched = first.Char == ")"
	}
	return matched
}
