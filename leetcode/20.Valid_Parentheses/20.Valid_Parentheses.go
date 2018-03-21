package Valid_Parentheses

import (
	"container/list"
)

func isValid(s string) bool {
	reflect := make(map[rune]rune)
	reflect['('] = ')'
	reflect['['] = ']'
	reflect['{'] = '}'

	l := list.New()
	for _, v := range s {
		switch v {
		default:
			return false
		case '(', '[', '{':
			l.PushBack(v)
		case ')', ']', '}':
			e := l.Back()
			if e == nil || !isMatch(reflect, v, l.Remove(e).(rune)) {
				return false
			}
		}
	}

	return l.Len() == 0

}

func isMatch(reflect map[rune]rune, right, left rune) bool {
	if reflect[left] == right {
		return true
	}
	return false
}
