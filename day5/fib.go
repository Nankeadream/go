package main

func Fib0(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib0(n-1) + Fib0(n-2)
}

func Fib1(n int) int {
	switch {
	case n < 0:
		panic("n is negative")
	case n == 0:
		return 0
	case n == 1 || n == 2:
		return 1
	}
	a, b := 1, 1
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}
	return b
}

var M = make(map[uint]uint, 100)

func Fib2(n uint) uint {
	if n <= 1 {
		return n
	}
	M[0] = 1
	M[1] = 1
	if _, ok := M[n]; !ok {
		M[n] = Fib2(n-1) + Fib2(n-2)
	}
	return M[n]
}
func Fib3(a, b, n int) int {
	if n < 3 {
		return b
	}
	return Fib3(b, a+b, n-1)
}
