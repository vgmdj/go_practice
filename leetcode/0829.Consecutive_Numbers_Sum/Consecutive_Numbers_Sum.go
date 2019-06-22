package Consecutive_Numbers_Sum

// Best
func consecutiveNumbersSum(N int) int {
	for N &1 == 0 {
		N >>= 1
	}
	ans := 1
	d := 3
	for d*d <= N {
		e := 0
		for N %d == 0 {
			N /= d
			e++
		}
		ans *= e+1
		d += 2
	}
	if N > 1 {
		ans <<=1
	}
	return ans
}

// AC
func consecutiveNumbersSum2(N int) int {
	count, k, half := 1, 2, N/2+1

	for k <= half {
		sk := N - k*(k-1)/2
		if sk <= 0 {
			return count
		}

		if sk%k == 0 {
			count++
		}

		k++

	}

	return count

}

//Time Limit Exceeded
func consecutiveNumbersSum3(N int) int {
	count, start, half := 1, 1, N/2+1
	k := half - start + 1

	for start < half && k > 1 {
		sum := (2*start + k - 1) * k

		if sum == 2*N {
			count++
			start++
			k--

		} else if sum > 2*N {
			k--

		} else {
			start++

		}

	}

	return count

}

//Time Limit Exceeded
func consecutiveNumbersSum4(N int) int {
	count, start, end, half := 1, 1, 2, N/2+1
	sum := 3

	for start < end && start < half && end <= half {
		if sum == N {
			count++
			sum -= start
			start++

			end++
			sum += end

		} else if sum > N {
			sum -= start
			start++

		} else {
			end++
			sum += end
		}

	}

	return count

}
