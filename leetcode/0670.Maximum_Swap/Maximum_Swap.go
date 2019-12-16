package Maximum_Swap

func maximumSwap(num int) int {
	if num <= 11 {
		return num
	}

	n1, n2, p1, p2, max, p := 0, 0, 0, 0, 0, 0

	for n, i := num, 1; n > 0; i *= 10 {
		d := n % 10
		if d > max {
			max, p = d, i
		} else if max > d {
			n1, n2, p1, p2 = d, max, i, p
		}
		n /= 10
	}

	return num + (n2-n1)*p1 + (n1-n2)*p2
}
