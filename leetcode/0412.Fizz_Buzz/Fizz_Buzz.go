package Fizz_Buzz

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 2; i < n; i += 3 {
		result[i] += "Fizz"
	}

	for i := 4; i < n; i += 5 {
		result[i] += "Buzz"
	}

	for i := 0; i < n; i++ {
		if result[i] == "" {
			result[i] = strconv.Itoa(i + 1)
		}
	}

	return result

}
