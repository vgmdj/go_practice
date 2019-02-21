package Fibonacci_Number

func fib(N int) int {
	if N < 2 {
		return N
	}
	p, q, i := 0, 1, 1
	for i < N {
		q, p = q+p, q
		i++
	}
	return q
}
