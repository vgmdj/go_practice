package Decrypt_String_from_Alphabet_to_Integer_Mapping

func freqAlphabets(s string) string {
	result := make([]byte, 0, len(s))
	index := 0
	for index < len(s) {
		if index+2 < len(s) && s[index+2] == '#' {
			result = append(result, decode(s[index:index+2]))
			index += 3
			continue
		}

		result = append(result, decode(s[index:index+1]))
		index++

	}

	return string(result)
}

func decode(str string) byte {
	if len(str) == 1 {
		return str[0] - '1' + 'a'
	}

	add := (str[0] - '0') * 10
	add += str[1] - '0'

	return add - 1 + 'a'
}
