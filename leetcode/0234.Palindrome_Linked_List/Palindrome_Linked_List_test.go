package Palindrome_Linked_List

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T){
	ast := assert.New(t)

	ast.Equal(true,isPalindrome(&ListNode{
		Val:0,
		Next:&ListNode{
			Val:0,
		},
	}))


}
