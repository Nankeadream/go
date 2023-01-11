package main

import (
	"fmt"
	"sort"
)

// 类型转换
func tc() {
	//s1 := "abc" //3byte
	//s2 := "测试"  //6byte
	//fmt.Printf("%T %v %[3]v, %T %v %[6]v\n", s1, s1, len(s1), s2, s2, len(s2))
	////强制类型转换 string => []byte; string => []rune
	////注意[]byte表示字节序列；[]rune表示rune序列
	//fmt.Println([]byte(s1))
	//fmt.Println([]rune(s1))
	//fmt.Println([]byte(s2)) //utf-8 bytes,长度为6即6个字节
	//fmt.Println([]rune(s2)) //unicode切片，长度为2，每个元素4字
	//fmt.Printf("%x, %x", 27979, 35797)
	//fmt.Println("---------------------------")
	//
	////[]byte => string
	//fmt.Println(string([]byte{49, 65, 97}))
	////[]byte => string
	//fmt.Println(string([]rune{27979, 35797}))
	////rune使用unicode，但是字符串内部使用utf-8
	//fmt.Println("---------------------------")
	//s3 := []rune(s2)
	//fmt.Println(s2[0], s2[1], s2[2]) // 230 181 139
	//fmt.Println(s3[0], s3[1], len(s3), cap(s3), s3)
	//fmt.Printf("%T %[1]v, %[1]c\n", s3)
	//fmt.Println(string(s3)) //测 试
	//fmt.Println(string([]byte{230, 181, 139}))
	//
	//for i, i2 := range s2 {
	//	fmt.Println(i, i2)
	//}
	//for i := 0; i < len(s2); i++ {
	//	fmt.Println(i, s2[i])
	//}
}

// 字符串
func s() {
	//s := "www.magedu.edu.com马哥教育"
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%d, %T, %[2]d %[2]c\n", i, s[i])
	//}
	//for i, i2 := range s {
	//	fmt.Printf("%d, %T, %[2]d %[2]c %[2]c\n", i, i2, s[i])
	//}
	//s0 := "www.magedu.edu"
	//s1 := "马哥教育"
	//fmt.Printf("s0,s1 %s %p; %s %p\n", s0, &s0, s1, &s1)
	//s2 := s0 + "-" + s1
	//fmt.Printf("s2 %s %p\n", s2, &s2)
	//s3 := strings.Join([]string{s0, s1}, "-")
	//fmt.Printf("s3 %s %p\n", s3, &s3) //www.magedu.edu-马哥教育
	//fmt.Println("----------------------")
	//s4 := fmt.Sprintf("%s-%s", s0, s1)
	//fmt.Printf("s4 %s %p\n", s4, &s4)
	//var b strings.Builder
	//b.WriteString(s0)
	//b.WriteByte('-')
	//b.WriteString(s1)
	//for _, i2 := range s0 {
	//	fmt.Fprintf(&b, "%c-", i2)
	//}
	//for _, i2 := range s1 {
	//	fmt.Fprintf(&b, "%c-", i2)
	//}
	//b.WriteString(s1)
	//s5 := b.String()
	//fmt.Printf("s5 %s %p\n", s5, &s5)
	//fmt.Println(b.String())
	//for i := 0; i < len(s); i++ {
	//	if i <= 17 {
	//		fmt.Printf(
	//			//"%d %d %[2]c\n", i, s[i],
	//			strings.ToUpper(string(s[i])),
	//		)
	//	}
	//}
	//fmt.Println()
	//for i, i2 := range s {
	//	fmt.Println(i, i2)
	//}
	//for i, i2 := range s {
	//	if i <= 17 {
	//		fmt.Println(
	//
	//			strings.ToUpper(string(i2)),
	//		)
	//	}
	//}
	//fmt.Println('马', "马"[0])
	//fmt.Println(
	//	strings.Index(s, "马"),
	//	strings.IndexAny(s, "马m"),
	//	strings.IndexByte(s, 111),
	//	strings.IndexRune(s, 39532),
	//	strings.Contains(s, "ma"),
	//	strings.Count(s, "m"),
	//
	//	strings.ToLower(s),
	//	strings.ToUpper(s),
	//	strings.HasPrefix(s, "www"),
	//	strings.HasSuffix(s, "教育"),
	//)
	//s := "\v\b\r \tdddd\tccc\t \v\r\n"
	//TrimSpace: 去除字符串两端的空白字符
	//fmt.Println(strings.TrimSpace("\v\b\r \tabc\txyz\t \v\r\n"))
	////fmt.Println("\taaaa\t")
	////fmt.Println()
	////fmt.Println(strings.TrimSpace("\v\b\r aaaa"))
	////fmt.Println(strings.TrimSpace(s))
	//fmt.Println(strings.TrimSpace("\v\b   \tHello\tGophers\t \v\r\n"))
	////TrimPrefix:如果开头匹配，则去除，否则，返回原字符的副本
	//fmt.Println(strings.TrimPrefix("www.magedu.edu-马哥教育", "www."))
	////for i, i2 := range s {
	////	fmt.Printf("%d %d %c\n", i, i2, i2)
	////}
	////TrimSuffix: 如果开头匹配，则去除，否则，返回原字符的副本
	//fmt.Println(strings.TrimSuffix("www.magedu.edu-马哥教育", "教育"))
	////TrimRight: 字符串结尾的字符如果在字符集中，则全部一处，知道碰到第一个不在字符集中的字符为止
	//fmt.Println(strings.TrimLeft("abcdddeabeccc", "abcd"))
	////Trim: 字符串两头的字符如果在字符集中，则全部移除，知道左或有都碰到第一个不在字符集中的字符
	//fmt.Println(strings.Trim("abcdddeabeccc", "abcd"))

	//分割
	//s := "www.magedu.edu.com马哥教育"
	//fmt.Println(strings.Split(s, ","))
	//s1 := strings.SplitN(s, ".", 0)
	//fmt.Println(s1, len(s1))
	//fmt.Printf("%v,nil=%v\n", s1, s1 != nil)
	//s2 := strings.SplitAfter(s, "www")
	//fmt.Printf("%v %d\n", s2, len(s2))
	//s3 := strings.SplitAfterN(s, "w", 3)
	//fmt.Printf("%v %d\n", s3, len(s3))
	//s4 := strings.Replace(s, "w", "e", -1)
	//fmt.Printf("%v %d\n", s4, len(s4))

	//s := "www.magedu.edu-马哥教育"
	//for i, i2 := range s {
	//	fmt.Printf("%d %c;\n", i, i2)
	//}
	//fmt.Println(strings.Map(func(r rune) rune {
	//	if 'a' <= r && r <= 'z' {
	//		return r - 0x20
	//	}
	//	return r
	//}, s))
	//var i int8 = -1
	//var j uint8 = uint8(i)
	//fmt.Println(i, j) //-1, 255
	////fmt.Println(int(3.14)) //错误,不允许无类型float常量转到int
	//
	//var a = 3.14                               //定义有类型变量就没有问题
	//fmt.Printf("%T: %[1]v => %d\n", a, int(a)) //float64: 3.14 =>3
	////byte rune本质上就是整数和无类型常量可以直接计算，自动转换
	//b := 'a'
	//c := b + 1
	//fmt.Printf("%T %[1]d\n", c)
	//e := 1
	//f := b + e //rune和int类型不能家，必须转换。比如c:=int(b)+e或c

	//类型别名和类型定义
	//var a byte = 'C'
	//var b uint8 = 49
	//fmt.Println(a, b, a+b)
	//原因是在源码中定义了type byte = uint8, byte是uint8的别名.
	//go v1.9开始使用了类型别名，将byte等的定义修改成了类型别名的方式.
	//别名说明就是uint8的另外一个名字,和uint8是一回事。在看一段代码，正确吗？
	//type myByte uint8
	//
	//var a byte = 'C'
	//var b uint8 = 49
	//fmt.Println(a, b, a+b) //为什么类型不同，可以相加
	//var c myByte = 50
	//fmt.Println(a, c, a+c) //可以吗？为什么？
	//答案是不可以，原因就是Go原因不允许不通类型随便运算，就算我们眼睛看到可以，也不行
	//必须强制类型转换，转换是否成功，程序员自己判断
	//type myByte uint8 //类型定义
	//type byte = uint8 //类型别名
	//byte只是存在代码中，为了方便阅读或者使用，编译完成后，不会有byte类型.

	//字符串转换
	//在字符那一节，我们介绍过了字符串与[]byte,字符串与[]rune之间互转的例子，
	//s1 := "127"
	//fmt.Println(strconv.Atoi(s1))             //十进制理解
	//fmt.Println(strconv.ParseInt(s1, 16, 32)) //十六进制理解0x127
	//s2 := "115.6"
	//fmt.Println(strconv.ParseFloat(s2, 64))
	//fmt.Println(strconv.ParseBool("true"))
}

// 数据结构
func ds() {
	//https://pkg.go.dev/crypto/sha256#example-New
	//h := sha256.New()
	//h.Write([]byte("abc"))
	//r := h.Sum(nil)
	//s := fmt.Sprintf("%x", r)
	//fmt.Printf("%T %s %d\n", r, s, len(s))
	//https:pkg.go.dev/croypto/md5#example-New
	//m := md5.New()
	//m.Write([]byte("abc"))
	//fmt.Printf("%T %[1]x\n", m.Sum(nil)) //[]uint8
	//
	//m.Reset()
	//m.Write([]byte("abd"))
	//fmt.Printf("%T %[1]x\n", m.Sum(nil)) //uint8
	//f, err := os.Open("/Users/zxl/go/src/study/magego11")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//h := sha256.New()
	//if _, err := io.Copy(h, f); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%x", h.Sum(nil))

	//var m1 map[string]int{} //int,很危险，map不是零值可用
	//fmt.Println(m1, m1 == nil)
	//m1["t"] = 200 //panic

	//var m0 = map[string]int{}
	//var m1 = map[string]int{
	//	"a": 11,
	//	"b": 22,
	//	"c": 33,
	//}
	//make
	m2 := make(map[int]string) //一个较小的起始空间

	m2[100] = "abc"
	//m3 := make(map[int]string, 100) //容量100，长度为0，为了减少扩瞳，提前给合适的容量
	var m = make(map[string]int)
	m["a"] = 11 //key不存在，则创建新的key和value对
	m["a"] = 22 //key已经存在，则覆盖value
	m["b"] = 33
	fmt.Println(m["a"]) //存在返回22
	fmt.Println(m["b"]) //不存在返回零值0，这样不能判断b这个key存在否，需要解析返回值
	fmt.Println(m)
	if _, ok := m["b"]; !ok {
		fmt.Println("key b不存在", m["b"])
	} else {
		fmt.Println("key b的值等于", m["b"])
	}
	//delete(m, "a") //存在，删除kv对
	//delete(m, "b") //不存在，删除操作也不会panic
	//fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("------------------")
	a := []int{-1, 23, 5, 9, 7}
	//sort.Sort(sort.Intslice(a)) //默认升序，有快捷写法Ints
	sort.Ints(a)   //就地修改原切片的底层数组
	fmt.Println(a) //默认升序
	//降序 sort.IntSlice(a)强制类型转换以施加接口方法
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
	fmt.Println("------------------")
	//二分查找
	sort.Ints(a)
	//二分查找，必须是升序
	//二分查找的前提是，有序
	i := sort.SearchInts(a, 7)
	fmt.Println(i)
}

func test() {
	arr := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	s := make([]int, len(arr)-1)
	fmt.Printf("%d,%v\n", len(arr), arr)
	fmt.Printf("%d,%v\n", len(s), s)
	for i := 0; i < len(s); i++ {
		s[i] = arr[i] + arr[i+1]
	}
	fmt.Println(s)
}
func main() {
	//tc()
	//s()
	//ds()
	test()
}
