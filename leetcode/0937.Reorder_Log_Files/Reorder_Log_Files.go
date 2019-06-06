package Reorder_Log_Files

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	tail := len(logs) - 1
	p := len(logs) - 1

	for p >= 0 {
		if !isDigitLogs(logs[p]) {
			p--
			continue
		}

		if p != tail {
			logs[p], logs[tail] = logs[tail], logs[p]
		}

		p--
		tail--
	}

	sort.Slice(logs[:tail+1], func(i, j int) bool {
		ihead, itail := split(logs[i])
		jhead, jtail := split(logs[j])

		if itail == jtail {
			return ihead < jhead
		}

		return itail < jtail
	})

	return logs

}

func isDigitLogs(s string) bool {
	idx := strings.Index(s, " ")
	return unicode.IsDigit(rune(s[idx+1]))

}

func split(s string) (head, tail string) {
	index := strings.Index(s, " ")
	return s[:index], s[index+1:]
}
