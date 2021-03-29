package dynamic

func fib(fibs []int, n int) int {
	if n < 2 {
		return n
	}
	if fibs[n] == 0 {
		fibs[n] = fibs[n-1] + fibs[n-2]
	}
	return fibs[n]
}

func Fibonacci(size int) (fibs []int) {
	fibs = make([]int, size)
	for i := range fibs {
		fibs[i] = fib(fibs, i)
	}
	return
}
