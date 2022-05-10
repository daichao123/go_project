package main

func getSum(n int) (sum int) {
	if n < 1 {
		return 0
	}
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func GetSumRecursive(n int) (sum int) {
	if n <= 1 {
		return n
	}
	return n + GetSumRecursive(n-1)
}

func FibonacciNumbersRecursive(n int) (sum int) {
	if n == 1 || n == 0 {
		return 1
	}
	return FibonacciNumbersRecursive(n-1) + FibonacciNumbersRecursive(n-2)
}
