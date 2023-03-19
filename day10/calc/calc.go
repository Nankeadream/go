package calc

import "fmt"

func Add(x, y int) int {
	fmt.Printf("x and y in calc.Add: %d, %d\n", x, y)
	return x + y
}
