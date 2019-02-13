package Lemonade_Change

func lemonadeChange(bills []int) bool {
	bill5, bill10 := 0, 0

	for _, v := range bills {
		switch v {
		case 5:
			bill5++

		case 10:
			bill10++
			bill5--

		case 20:
			if bill10 > 0 {
				bill10--
				bill5--
			} else {
				bill5 -= 3
			}
		}

		if bill5 <0 || bill10 <0{
			return false
		}
	}

	return true

}
