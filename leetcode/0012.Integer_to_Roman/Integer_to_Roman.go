package Integer_to_Roman

var roman = make(map[int]string)
var keys []int

func init() {
	keys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	roman = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

}

func intToRoman(num int) string {
	var result string
	for _, v := range keys {
		t := num / v
		for i := 0; i < t; i++ {
			num -= v
			result += roman[v]
		}
	}

	return result

}
