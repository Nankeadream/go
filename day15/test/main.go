package main

import "fmt"

func main() {
	var (
		name     string
		age      int
		is_marry bool
	)

	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &is_marry)
	fmt.Printf("扫描结果 name:%s age:%d is_marry:%t \n", name, age, is_marry)
}
