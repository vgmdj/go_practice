package _022_Generate_Parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParentheses(t *testing.T) {
	ast := assert.New(t)

	ast.EqualValues([]string{"()"}, generateParenthesis(1))

	ast.EqualValues([]string{"(())", "()()"}, generateParenthesis(2))

	ast.EqualValues([]string{
		"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))

	ast.EqualValues([]string{
		"(((())))", "((()()))", "((())())",
		"((()))()", "(()(()))", "(()()())",
		"(()())()", "(())(())", "(())()()",
		"()((()))", "()(()())", "()(())()",
		"()()(())", "()()()()",
	}, generateParenthesis(4))

}
