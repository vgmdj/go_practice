package Convert_a_Number_to_Hexadecimal

import "strings"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	var result string

	for i := 0; i*4 < 32; i++ {
		hex := (num >> (uint(i) * 4)) & 0xf
		if hex < 10 {
			result = string(hex+'0') + result
			continue
		}

		result = string(hex-10+'a') + result
	}

	return strings.TrimLeft(result, "0")

}
