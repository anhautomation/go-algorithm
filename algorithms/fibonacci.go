package algorithms

func Fibonacci(n int) (int) {
	if n <= 1 {
		return n
	}
	f := make([]int, n+1)
	f[0], f[1] = 0, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func FibonacciRecursion(n int) (int) {
    if n <= 1 {
        return n
    }
    return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
