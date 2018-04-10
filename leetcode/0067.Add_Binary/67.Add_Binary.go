package Add_Binary

func addBinary(a string, b string) string {
	if len(a) < 1 || len(b) < 1 {
		return "-1"
	}

	count, length, carry := 0, len(a)+len(b), byte(0)
	rst := make([]byte, length)

	for ai, bi := len(a)-1, len(b)-1; ai >= 0 || bi >= 0 || carry != 0; ai, bi = ai-1, bi-1 {
		sum := carry
		if ai >= 0 {
			sum += a[ai] - '0'
		}
		if bi >= 0 {
			sum += b[bi] - '0'
		}

		rst[length-count-1] = sum%2 + '0'
		carry = sum / 2
		count++
	}

	return string(rst[length-count:])

}
