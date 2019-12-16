package Regular_Expression_Matching

import "testing"

func TestOK(t *testing.T) {
	t.Log(isMatch("aa", "a"))                    // → false
	t.Log(isMatch("aa", "aa"))                   // → true
	t.Log(isMatch("aaa", "aa"))                  // → false
	t.Log(isMatch("aa", "a*"))                   // → true
	t.Log(isMatch("aa", ".*"))                   // → true
	t.Log(isMatch("abbc", ".*c"))                // → true
	t.Log(isMatch("aab", "c*a*b"))               // → false
	t.Log(isMatch("mississippi", "mis*is*ip*.")) // → true
}
