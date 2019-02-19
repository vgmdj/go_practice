package _022_Generate_Parentheses

func generateParenthesis(n int) []string {
	var result []string
	backTrack(&result, "", 0, 0, n)
	return result

}

func backTrack(result *[]string, parenthesis string, open, close, n int) {
	if open > n || close > n {
		return
	}

	if open == n && close == n {
		*result = append(*result, parenthesis)
		return
	}

	if open < n {
		backTrack(result, parenthesis+"(", open+1, close, n)

	}

	if close < open {
		backTrack(result, parenthesis+")", open, close+1, n)

	}

	return

}
