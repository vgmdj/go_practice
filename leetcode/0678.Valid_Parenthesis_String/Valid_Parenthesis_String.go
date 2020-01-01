package Valid_Parenthesis_String

/*
( *
* )
( * )
( ( ( * ) * ) * )

( * (
( ) ) *

*/
func checkValidString(s string) bool {
	left, right := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '*' {
			left++

		} else {
			left--

		}

		if left < 0 {
			return false
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || s[i] == '*' {
			right++

		} else {
			right--

		}

		if right < 0 {
			return false
		}

	}

	return true

}
