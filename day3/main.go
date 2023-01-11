package main

import (
	"fmt"
	"unsafe"
)

//type Person struct {
//	name string
//}

func his() {
	//var a = 0x63
	//fmt.Printf("%T %[1]d %[1]c\n", a)
	//var b byte = 0x63
	//fmt.Printf("%T %[1]d %[1]c\n", b)
	//var c rune = '\x63'
	//fmt.Printf("%T %[1]d %[1]c\n", c)
	//var d = "\x63"
	//fmt.Printf("%T %[1]s\n", d)
	//fmt.Printf("%T %[1]s\n", string(a))

	//fmt.Println(1/2, 3/2, 5/2)
	//fmt.Println(-1/2, -3/2, -5/2)
	//fmt.Println("-------Ceil-------------")
	//fmt.Println(math.Ceil(2.01), math.Ceil(2.5), math.Ceil(2.8))
	//fmt.Println(math.Ceil(-2.01), math.Ceil(-2.5), math.Ceil(-2.8))
	//fmt.Println("--------Floor------------")
	//fmt.Println(math.Floor(2.01), math.Floor(2.5), math.Floor(2.8))
	//fmt.Println(math.Floor(-2.01), math.Floor(-2.5), math.Floor(-2.8))
	//fmt.Println("---------Round-----------")
	//fmt.Println(math.Round(2.01), math.Round(2.5), math.Round(2.8))
	//fmt.Println(math.Round(-2.01), math.Round(-2.5), math.Round(-2.8))
	//fmt.Println(math.Round(0.49999), math.Round(1.5), math.Round(2.5), math.Round(3.5))

	//fmt.Println(math.Abs(-2.7))                               //绝对值
	//fmt.Println(math.E, math.Pi)                              //常数
	//fmt.Println(math.MaxInt16, math.MinInt16)                 //常量，极值
	//fmt.Println(math.Log10(100), math.Log2(8))                //对数
	//fmt.Println(math.Max(1, 2), math.Min(-2, 3))              //最大值，最小值
	//fmt.Println(math.Pow(2, 4), math.Pow10(3))                //幂
	//fmt.Println(math.Mod(5, 2), 5%2)                          //取模
	//fmt.Println(math.Sqrt(3), math.Sqrt(6), math.Pow(2, 0.5)) //开方

	//var err error
	//var name string
	//var age int
	//var world1, world2 string
	//fmt.Print("Plz input words: ")
	//n, err = fmt.Scan(&world1, &world2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(n)
	//fmt.Printf("%T %s, %T %s\n", world1, world1, world2, world2)
	//fmt.Println("--------------------------")
	//var i1, i2 int
	//fmt.Print("Plz input two ins: ")
	//n, err = fmt.Scan(&i1, &i2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%T %[1]d, %T %[2]d", i1, i2)

	//fmt.Print("Plz input your name and age: ")
	//_, err = fmt.Scanf("%s %d\n", &name, &age)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(name, age)
	//
	//var weight, height int
	//fmt.Print("weight and height: ")
	//_, err = fmt.Scanf("%d %d", &weight, &height)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%T %[1]d,%T %[2]d", weight, height)

	//var a0 [3]int
	//var a1 = [3]int{}
	//var a2 [3]int = [3]int{1, 3, 5}
	//var a3 = [3]int{1, 3, 5}
	//fmt.Println(a0, a1)
	//const a = 3
	//a4 := [a]int{1, 3, 5}
	//fmt.Println(a2, a3, a4)
	//a5 := [...]int{10, 30, 50, 40, 60, 70} //...让编译器确定当前数组大小
	//a6 := [5]int{100, 200}                 //顺序初始化前面的，其余用零值填充
	//a7 := [5]int{1: 300, 3: 300}           //指定索引位置初始化，其余用零值填充
	////二维数组
	//a8 := [3][10]int{{100}} //两行三列[[100 0 0] [0 0 0]]
	////[[10 0 0] [11 12 0] [13 14 15] [16 0 0]]
	////多维数组，只有第一维才能用...推测
	////第一维有4个，第二维有3个，可以看做4行3列的表
	//a9 := [...][3]int{{10}, {11, 12}, {13, 14, 15}, {16}}
	//fmt.Printf("a5:%v len:%v cap:%v\n", a5, len(a5), cap(a5))
	//fmt.Printf("a6:%v len:%v cap:%v\n", a6, len(a6), cap(a6))
	//fmt.Printf("a7:%v len:%v cap:%v\n", a7, len(a7), cap(a7))
	//fmt.Printf("a8:%v len:%v cap:%v\n", a8, len(a8), cap(a8))
	//fmt.Printf("a9:%v len:%v cap:%v\n", a9, len(a9), cap(a9))

	//a5 := [...]int{10, 30, 50}
	//a5[0] += 100
	//for i := 0; i < len(a5); i++ {
	//	fmt.Println(i, a5[i])
	//}
	//for i, i2 := range a5 {
	//	fmt.Println(i, i2, a5[i])
	//}
	//var a = [...]int{100, 400, 300, 505, 289, 67878}
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(i, a[i], &a[i])
	//}
	//fmt.Printf("%p %p, %v\n", &a, &a[0], a)
	//a[0] = 1000
	//fmt.Printf("%p %p, %v\n", &a, &a[0], a)
	//fmt.Printf("%p %p, %v\n", &a, &a[1], a)
	//fmt.Printf("%p %p, %v\n", &a, &a[2], a)
	//fmt.Println("-----------------------")
	//for _, i2 := range a {
	//	fmt.Printf("%p %p, %v\n", &a, &i2, i2)
	//}
	//fmt.Println("-----------------------")
	//for i := 0; i < len(a); i++ {
	//	fmt.Printf("%p %p, %v\n", &a, &a[i], a[i])
	//}

	//arr := []Person{
	//	Person{"明月"},
	//	Person{"薇薇"},
	//}
	//var res []*Person
	//
	//for _, v := range arr {
	//	res = append(res, &v)
	//}
	//// 遍历查看结果集
	//for _, person := range res {
	//	fmt.Println("name-->:", person.name)
	//}
	//for _, v := range arr {
	//	fmt.Printf("v 指针 %p\n", &v)
	//	fmt.Println("v 的值", v)
	//	res = append(res, &v)
	//}
	////for range 的时候，v 只初始化了一次，之后的遍历都是在原来遍历的基础上赋值，所有v的指针（地址）并没有变。该指针指向的是最后一次遍历的v的值，所以最后结果集中，也就都成了最后遍历的v的值
	//fmt.Println(res)
	//
	//for i, _ := range arr {
	//	fmt.Printf("v 指针 %p\n", &arr[i])
	//	fmt.Println("v 的值", arr[i])
	//	res = append(res, &arr[i])
	//}

	//var a [3]int
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(i, a[i], &a[i])
	//}
	//fmt.Printf("%p %p,%v\n", &a, &a[0], a)
	//a[0] = 1000
	//fmt.Printf("%p %p,%v\n", &a, &a[0], a)

	//var a = [3]string{"abc", "def", "xyz"}
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(i, a[i], &a[i])
	//}
	//fmt.Printf("%p %p,%v\n", &a, &a[0], a)
	//a[0] = "oooooo"
	//fmt.Printf("%p %p,%v\n", &a, &a[0], a)
}
func showAddr(s []int) []int {
	fmt.Printf("s %p,%p,%d,%d,%v\n", &s, &s[0], len(s), cap(s), s)
	//if len(s) > 0 {
	//	s[0] = 123
	//}
	s = append(s, 100, 200)
	fmt.Printf("s %p,%p,%d,%d,%v\n", &s, &s[0], len(s), cap(s), s)
	return s
}

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func sli() {
	//var s1 []int               //长度，容量为0的切片，零值
	//var s2 = []int{}           //长度，容量为0的切片，字面量定义
	//var s3 = []int{1, 3, 5}    //字面量定义，长度，容量都是3
	//var s4 = make([]int, 0)    //长度，容量为0的切片，make([]T,length)
	//var s5 = make([]int, 3, 5) //长度为3，容量为5，底层数组长度为5，元素长度为3，所以显示[0 0 0]
	//fmt.Printf("s1: %v\n", s1)
	//fmt.Printf("s2: %v\n", s2)
	//fmt.Printf("s3: %v\n", s3)
	//fmt.Printf("s4: %v\n", s4)
	//fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	//a := []int{1, 3, 5, 7}
	//fmt.Printf("%v, %p,%p\n", a, &a, &a[0])
	//b := a
	//b = append(b, 100)
	//fmt.Printf("%v, %p,%p\n", b, &b, &b[0])
	//for i := 0; i < len(a); i++ {
	//	fmt.Printf("%p %p,%v\n", &a, &a[i], a[i])
	//}

	//s1 := make([]int, 3, 100)
	//fmt.Printf("s1: %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//fmt.Println("-----------------------")
	//s2 := append(s1, 1, 2)
	//fmt.Printf("s1: %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//fmt.Printf("s2: %p,%p, l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	//fmt.Println("-----------------------")
	//s3 := append(s1, -1)
	//fmt.Printf("s1: %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//fmt.Printf("s2: %p,%p, l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	//fmt.Printf("s3: %p,%p, l=%-2d,c=%-2d,%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	//fmt.Println("-----------------------")
	//s4 := append(s3, 3, 4, 5)
	//fmt.Printf("s1: %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//fmt.Printf("s2: %p,%p, l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	//fmt.Printf("s3: %p,%p, l=%-2d,c=%-2d,%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	//fmt.Printf("s4: %p,%p, l=%-2d,c=%-2d,%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	//fmt.Println("-----------------------")
	//s5 := append(s4, 6, 7, 8, 9)
	//fmt.Printf("s1: %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//fmt.Printf("s2: %p,%p, l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	//fmt.Printf("s3: %p,%p, l=%-2d,c=%-2d,%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	//fmt.Printf("s4: %p,%p, l=%-2d,c=%-2d,%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	//fmt.Printf("s5: %p,%p, l=%-2d,c=%-2d,%v\n", &s5, &s5[0], len(s5), cap(s5), s5)
	//fmt.Println("-----------------------")
	//s1 := []int{10, 20, 30}
	//fmt.Printf("s1 %p,%p,%d,%d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//s2 := s1
	//fmt.Printf("s2 %p,%p,%d,%d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	//fmt.Println("-----------------------")
	//s3 := showAddr(s1)
	//fmt.Printf("s1 %p,%p,%d,%d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//fmt.Printf("s2 %p,%p,%d,%d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	//fmt.Printf("s3 %p,%p,%d,%d,%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	//s1 := []int{10, 30, 50, 70, 90} //容量，长度为5，索引0，1，2，3，4
	//fmt.Printf("s1 %p, %p,%d ,%d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//s2 := s1 //和s1共用底层数组
	//fmt.Printf("s2 %p, %p,%d ,%d, %v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	//s3 := s1[:] //和s1共用底层数组，从头到尾元素都要
	//fmt.Printf("s3 %p,%p,%d,%d,%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	//fmt.Println("-----------------------")
	//s4 := s1[1:] //掐头，容量，长度都为4，收地址偏移1个元素，共用底层数组
	//fmt.Printf("s4 %p,%p,%d,%d,%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	//s5 := s1[1:4] //掐头去尾，容量为4，长度为3，首地址偏移1个元素，共用底层数组
	//fmt.Printf("s5 %p,%p,%d,%d,%v\n", &s5, &s5[0], len(s5), cap(s5), s5)
	//s6 := s1[:4] //去尾，容量为5，长度为4，首地址不变，共用底层数组
	//fmt.Printf("s6 %p, %p, %d, %d, %v\n", &s6, &s6[0], len(s6), cap(s6), s6)
	//fmt.Println("-----------------------")
	//s7 := s1[1:1] //首地址偏移一个元素，长度为0，容量为4，共用底层数组
	//fmt.Printf("s7 %p, %d, %d, %v\n", &s7, len(s7), cap(s7), s7)
	//s8 := s1[4:4] //首地址偏移4个元素，长度为0，容量为1，因为最后一个元素没在切片中，共用底层数组
	//fmt.Printf("s8 %p %d,%d,%v\n", &s8, len(s8), cap(s8), s8)
	//s9 := s1[5:5] //首地址偏移5个元素，长度为0，容量为0，共用底层数组
	//fmt.Printf("s9 %p, %d,%d,%v\n", &s9, len(s9), cap(s9), s9)
	//fmt.Println("-----------------------")
	//s9 = append(s9, 11)
	//fmt.Printf("s9 %p,%p,%d,%d,%v\n", &s9, &s9[0], len(s9), cap(s9), s9)
	//fmt.Printf("s1 %p, %p,%d ,%d, %v\n", &s1, &s1[0], len(s1), cap(s1), s1)

	//s1 := [5]int{10, 30, 50, 70, 90} //容量，长度为5，索引0，1，2，3，4
	//fmt.Printf("s1 %p,%p,%d,%d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//s4 := s1[1:] //掐头，容量，长度都为4，首地址偏移1个元素，以s1为底层数组
	//fmt.Printf("s4 %p,%p,%d,%d,%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	//s4[0] = 100 //s1,s4分别是什么？
	//fmt.Printf("s1 %p,%p,%d,%d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//fmt.Printf("s4 %p,%p,%d,%d,%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	//s4 = append(s4, 11, 22) //是否扩容？会怎样？
	//fmt.Printf("s1 %p %p,%d,%d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	//fmt.Printf("s4 %p,%p,%d,%d,%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	//a1 := [2][3]int{{1: 2, 2}, {3, 4}}
	//fmt.Println(a1)

}
func main() {
	//his()
	sli()
	//a1 := [...]int{10, 30, 50}
	//fmt.Printf("a1: %v,%p\n", a1, &a1)
	//a2 := a1
	//fmt.Printf("a2: %v,%p\n", a2, &a2)
	//fmt.Println("--------------------")
	//a3 := showAddr(a1)
	//fmt.Printf("a3_1: %v,%p\n", a3, &a3)
	var s0 = []int{1, 3, 5, 7, 9}
	fmt.Printf("%T %[1]v,%d %d\n", s0, len(s0), cap(s0))
	var s1 []int
	fmt.Println(s1, len(s1), cap(s1), &s1)
	fmt.Printf("%p %[1]T\n", &s1)
	print(s1)
}
