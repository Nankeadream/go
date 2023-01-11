package main

// 参数a的作用,
// 参数b的作用,
// 参数c的作用
func test(a string, b int, c bool) {

}

func main1() {
	// TODO 我还有一个功能没有完成

	//声明并赋值，初始化，数组不可以在go中定义常量，元素可变
	//常量，标识符，指向的这个值不可变，其他语言中值的是元素地址不变，go要求指向的内容不可变
	//const aa = 1 //只能是字面常量
	//
	//a := "abc"     //字面表达同时这个值不能修改，字符串不可变
	//var b = 1      //1是一个常量，字面常量是不可变的
	//var c = 1 + 1  //1+1=2，2是由1+1得出也是一个常量，不可修改
	//fmt.Println(a) //标识符在编译后就看不到了，被替换成了内存地址，地址里面放着那个值
	//fmt.Println(b) //打印什么？
	//fmt.Println(c)

	//const a = iota //总是从0开始
	//const b = iota
	//fmt.Println(a, b)
	//const (
	//	c = iota
	//	d //go语言可以重复上一行的公式，批量里面一定要注意同一个iota
	//	e
	//	_ //可以用来做标识符，但是不可以使用
	//	_ //空白标识符，匿名变量
	//	f
	//	g = iota + 10
	//	h
	//	i = 20
	//	j
	//	k = iota * 2
	//	l
	//)
	//fmt.Println(c, d, e, f, g)
	//const ( //枚举
	//	a = iota * 2
	//	b
	//	c
	//	d
	//	e
	//)

	//变量定义
	//var a int //此处声明已经被初始化为该类型的零值
	//a = 100
	//fmt.Println(a)
	//var b = 200 //初始化，声明+初始化，从右往左推导出类型
	////var a int,b int //批量，不能怎么写
	//var (
	//	c int = 300
	//	d int //零值
	//)
	//fmt.Println(c, d)
	//var e, f int = 1, 2 //批量同类合并
	//fmt.Println(b, e, f)
	////var m,n int ,t string == 100,200,"abc" //错误
	//var (
	//	m int
	//	n int
	//	t string = "abc"
	//)
	//m, n = 100, 200 //批量赋值
	//m, n = n, m     //交换
	//fmt.Println(m, n, t)
	//a = 1000
	////短格式，声明定义方式和const，var冲突
	//r := 200 //类型冲突 int
	//r = 3000 //修改
	//fmt.Println(r)
	//
	//var x1, y1, z1 int
	//x1, y1, z1 = func() (int, int, int) { return 100, 200, 300 }()
	//fmt.Println(x1, y1, z1)
	//x, y, z := func() (int, int, int) { return 1, 2, 3 }() //推荐使用这种
	//fmt.Println(x, y, z)

}
