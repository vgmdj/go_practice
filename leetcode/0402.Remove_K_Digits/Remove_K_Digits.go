package Remove_K_Digits

import "strings"

func removeKdigits(num string, k int) string {
	stack := make([]byte, 1)
	for i := range num {
		for k > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])

	}

	result := strings.TrimLeft(string(stack[1:len(stack)-k]), "0")
	if len(result) == 0 {
		result = "0"
	}

	return result
}
