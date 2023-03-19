package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

//	type Person struct {
//		Name string `json:"name" msgpack:"myname"`
//		Age  int    `json:"age" msgpack:"myage"`
//	}
type Person struct {
	Name string
	Age  int
}

func Json() {
	var data = Person{
		"Tom",
		32,
	}
	//var data = []any{
	//	100, 2.5, true, false, nil, "accc",
	//	[...]int{97, 98, 99},
	//}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	//var target = make([][]byte, 0, len(data))
	//for i, datum := range data {
	//	//fmt.Printf("%T %+[1]v; %T %+[2]v\n", data, b)
	//	b, err := json.Marshal(datum)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	target = append(target, b)
	//	fmt.Printf("%d {%[2]T-%[3]v}====>%[4]T-%[4]v %[5]s----%[4]q\n", i, datum, datum, b, b, string(b))
	//
	//}
	//fmt.Println(string(b))
	//fmt.Printf("%T %[1]q\n%s\n", b, string(b))
	//fmt.Println("----------------")
	//fmt.Println(data, len(data), cap(data))
	//fmt.Println(target, len(target), cap(target))
	//fmt.Printf("%T %[1]q\n %+[1]v\n", target)
	fmt.Println(b, string(b))
	fmt.Println("---------------")
	//for i, v := range target {
	//	var a any
	//	err := json.Unmarshal(v, &a)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Printf("%d %T %v  %v--%[3]q\n", i, v, v, string(v))
	//}
	var a Person
	err = json.Unmarshal(b, &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T,%+[1]v\n", a)
}
func base() {
	var a = "abc" //3字节，试试abcd，abcde，abcdef
	s := base64.StdEncoding.EncodeToString([]byte(a))
	fmt.Println(len(s), s)
	fmt.Println("------------------")
	d, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(d), d, string(d))
}
func main() {
	//File()
	//Json()
	base()
}
