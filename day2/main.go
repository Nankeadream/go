package main

import "fmt"

func main() {
	//var d int64 = 50
	//fmt.Printf("%T,%d\n", d, d)
	//fmt.Println("------分割线---------")
	//fmt.Printf("%T, %s; %T, %d; %T, %f\n",
	//	string(d), string(d), rune(d), rune(d), float32(d), float32(d))

	//type rune = int32                   //rune是int32的别名，4个字节，可以是unicode字符
	//type byte = uint8                   //byte是uint8的别名，1个字符
	//var c rune = '中'                    //字符用单引号
	//fmt.Printf("%T, %c, %d\n", c, c, c) //int32 ,中， 20013
	//c = 'a'
	//fmt.Printf("%T, %c, %d\n", c, c, c) //int32， a , 97
	////var d byte = '中'  //错误，超出byte范围
	//var d byte = '\x61'
	//fmt.Printf("%T, %c, %d\n", d, d, d) //byte, b, 98
	//var e rune = 20013
	//fmt.Printf("%T, %c, %d\n", e, e, e) //int32,中,20013
	//fmt.Println("------分割线---------")
	//fmt.Printf("%3.2f\n", math.MaxFloat32)
	////fmt的格式化，参考包帮助https://pkg.go.dev/fmt
	//f := 12.15
	//fmt.Printf("%T,%f\n", f, f) //默认精度6
	//fmt.Printf("%.3f\n", f)     //小数点后3位
	//fmt.Printf("[%3.2f]\n", f)  //宽度撑爆了,中括号加上没有特殊含义，只是位了看清楚占的打印宽度
	////打印宽度
	//fmt.Printf("[%8.2f]\n", f)  //宽度为6
	//fmt.Printf("[%-8.2f]\n", f) //左对齐
	//var a rune = 0o664
	//fmt.Printf("%T %c %d ", a, a, a)

	//a, b, c, d := 100, 200, 300, 400
	//fmt.Printf("%d,%[2]v,%[1]d,%d", a, b, c, d)

	//var a = 1 * 2.3 // 不报错
	//fmt.Printf("%T %v\n", a, a)
	//
	//fmt.Println(2&1, 2&^1, 3&1, 3&^1, 15&5, 15&^5)
	//fmt.Println(2|1, 3^3, 1<<3, 16>>3, 2^1)

	//a := 123
	//b := &a
	//c := *b
	//fmt.Printf("%d,%p,%d\n", a, b, c)
	//fmt.Println(a == c, b == &c, &c, &b)
	//var d = a
	//fmt.Println(a == d, &a, &d)

	//	for i := 0; i < 20; i++ {
	//		rand.Seed(time.Now().UnixNano())
	//		s := rand.Intn(100)
	//		if i%2 == 0 {
	//			continue
	//		}
	//		fmt.Println(i, s)
	//		if i > 10 {
	//			goto condition
	//		}
	//	}
	//condition:
	//	fmt.Println("done")

	//a := "\xe6\xb5\x8b\xe8\xaf\x95"
	//fmt.Println(a)
	//for i, v := range "abc测试" {
	//	fmt.Printf("%d,%[2]d,%[2]c, %#[2]x\n", i, v)
	//}
	//fmt.Println("\xe6\xb5\x8b\xe8\xaf\x95")
	//arr := [5]int{1, 3, 5, 7, 9}
	//for i, v := range arr {
	//	fmt.Println(i, v, arr[i])
	//}
	//fmt.Println("----------------")
	//for i := range arr {
	//	fmt.Println(i, arr[i])
	//}
	//fmt.Println("----------------")
	//for _, i := range arr {
	//	fmt.Println(i)
	//}
	//src := rand.NewSource(10)
	//r10 := rand.New(src)
	//r1 := rand.New(rand.NewSource(1))
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%d,%d,%d\n", rand.Intn(5), r1.Intn(5), r10.Intn(10))
	//}
	fmt.Println(15&5, 15&^5)

}
