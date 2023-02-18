package main

import "fmt"

func fib(n int) int {
	switch {
	case n < 0:
		panic("n is not negative")
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

func fib1(n int) int {
	if n > 3 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}
func fib2(n, a, b int) int {
	if n > 3 {
		return b
	}
	return fib2(n-1, b, a+b)
}

func calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

type MyFunc = func(int, int) int

func calc2(a, b int, fn MyFunc) int {
	return fn(a, b)
}

func outer() func() {
	c := 99
	fmt.Printf("%p,\t%d\n", &c, c)
	var inner = func() {
		c = 100
		fmt.Printf("%p,\t%d\n", &c, c)
	}
	//inner()
	return inner
}

func outer2() {
	c := 99
	var inner = func() {
		c = 100
		fmt.Println("1 inner", c) //100
	}
	inner()
	fmt.Println("2 inner", c) //100
}

func outer3() func() {
	c := 99
	fmt.Printf("outher %d %p\n", c, &c)
	var inner = func() {
		fmt.Printf("inner %d %p\n", c, &c)

	}
	return inner
}

// PrintTrianglev1 倒三角
func PrintTrianglev1(n int) {
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
func main() {
	PrintTrianglev1(12)
	//fmt.Println(calc(4, 5, func(a int, b int) int {
	//	return a + b
	//}))
	//fmt.Println(calc(4, 5, func(a int, b int) int {
	//	return a * b
	//}))
	//fmt.Println(calc2(4, 5, func(a int, b int) int {
	//	return a + b
	//}))
	//fmt.Println(calc2(4, 5, func(a int, b int) int {
	//	return a * b
	//}))
	//
	//outer()
	//outer2()
	//var fn = outer3()
	//fn()
	//
	//var i int8 = -3
	//fmt.Println(i)
	//var j uint8 = uint8(i)
	//fmt.Println(j)
	//for i := 0; i < 45; i++ {
	//	fmt.Println(Fib0(i))
	//}
	//for i := 0; i < 100; i++ {
	//	Fib2(uint(i))
	//}
	//fmt.Println(M)
	//for i := 0; i < 12; i++ {
	//	fmt.Println(Fib3(1, 1, i))
	//}
	//fmt.Println(Fib3(1, 1, 10))
	//v := func(x, y int) int {
	//	return x + y
	//}(4, 5)
	//fmt.Println(v)
	//outer()
	//fn := outer()
	//fn()
}
