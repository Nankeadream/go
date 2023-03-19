package main

import (
	"fmt"
	"reflect"
)

type Runner interface {
	run()
}
type person struct {
	Name string `json:"name"`
	age  int
}

func (p *person) run() {

}
func (p person) GetName(prefix, suffix string) string {
	return fmt.Sprintf("%s %s %s", prefix, p.Name, suffix)
}

func main() {
	var a int = 100
	t := reflect.TypeOf(a)                       // 提取类型信息
	v := reflect.New(t)                          // 创建一个该类型的新的零值返回指针的Value，相当于 new(int)
	fmt.Println(v, v.Elem(), v.Type(), v.Kind()) // 内存地址 0 *int ptr
	fmt.Println(t)

}
