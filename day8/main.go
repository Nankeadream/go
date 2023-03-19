package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "/Users/zxl/magego/golang11/test.txt"
	//Open为快捷的只读打开
	//flag := os.O_RDWR | os.O_CREATE | os.O_TRUNC //可读可写，文件不存在就创建文件，文件存在就清空文件从头开始写
	f, err := os.Open(filename)
	if err != nil {
		defer f.Close()
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("%T %+[1]v\n", f)
	//读
	//buffer := make([]byte, 2) //每次读取两个字节
	//for true {
	//	n, err := f.Read(buffer)
	//	if err != nil {
	//		fmt.Println(err, n, "$$$")
	//	} else {
	//		fmt.Println(string(buffer), buffer, n, "@@@")
	//	}
	//	if n == 0 {
	//		break
	//	}
	//
	//}
	reader := bufio.NewReader(f)
	b1 := make([]byte, 3)
	n, err := reader.Read(b1)
	fmt.Println("read:", n, err, b1, b1[:n])
	b2, err := reader.ReadByte()
	fmt.Println("ReadByte:", b2, err)
	b3, err := reader.ReadBytes('\n')
	fmt.Println("ReadBytes:", b3, err, string(b3))
	fmt.Println("------------------------")
	r, size, err2 := reader.ReadRune()
	fmt.Println("ReadRune:", r, size, err2, string(r))
	line, err3 := reader.ReadSlice('\n')
	fmt.Println("ReadSlice:", line, err3, string(line))
	fmt.Println("------------------------")
	f.Seek(0, 0)
	s, err4 := reader.ReadString('\n')
	fmt.Println("RuneString s1:", s, []byte(s), err4)
	s1, err5 := reader.ReadString('\n')
	fmt.Println("RuneString s2:", s1, []byte(s1), err5)
	s2, err6 := reader.ReadString('\n')
	fmt.Println("RuneString s3:", s2, []byte(s2), err6)
	err = os.MkdirAll("a/b/c/d", os.ModePerm)
	if err != nil {
		return
	}
}
