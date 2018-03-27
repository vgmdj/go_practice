package Add_Binary

func addBinary(a string, b string) string {
	if len(a) < 1 || len(b) < 1 {
		return "-1"
	}

	as, bs := []byte(a), []byte(b)
	count, length, carry := 0, len(a)+len(b), byte('0')
	rst := make([]byte, length)

	for ai, bi := len(a)-1, len(b)-1; ai >= 0 || bi >= 0 || carry != byte('0'); ai, bi = ai-1, bi-1 {
		if ai < 0 {
			ai = 0
			as[0] = '0'
		}
		if bi < 0 {
			bi = 0
			bs[0] = '0'
		}

		sum := int(as[ai] - '0' + bs[bi] - '0' + carry - '0')

		rst[length-count-1] = byte(sum%2) + '0'
		carry = byte(sum/2) + '0'
		count++
	}

	return string(rst[length-count-1:])

}
