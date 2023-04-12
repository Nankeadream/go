package main

import (
	"container/heap"
	"day12/heaptest"
	"fmt"
	"math/rand"
	"time"
)

type IntHeap []int

//实现排序接口方法，抄sort.IntSlice，在VSCode中敲sort试试

// Len 长度就是数组长度
func (h IntHeap) Len() int {
	return len(h)
}

// Less 索引i和索引j内容直接比较大小
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap 索引i，j互换
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Pop 实现堆方法
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// func (h IntHeap) Test() {
//
//	n := h.Len()
//	fmt.Println(n)
//
// }
func heapSort(a []int) {
	h := &IntHeap{}
	for _, v := range a {
		heap.Push(h, v)
	}
	for i := len(a) - 1; i >= 0; i-- {
		a[i] = heap.Pop(h).(int)
	}
}
func main() {
	//h := &IntHeap{80, 90, 70, 40, 50, 10, 60, 20, 30}
	//heap.Init(h)
	//fmt.Println(h)
	//heap.Push(h, 75)
	//fmt.Println(h)
	//heap.Push(h, 15)
	//fmt.Println(h)
	//(*h)[0] = 45
	//fmt.Println(h)
	//heap.Fix(h, 0)
	//fmt.Println(h)
	//s := *h
	//fmt.Println(s)
	//fmt.Printf("%T\n", s)
	////s[0], s[len(s)-1] = s[len(s)-1], s[0]
	////s[1], s[len(s)-2] = s[len(s)-2], s[1]
	////s[0], s[10] = s[10], s[0]
	////for i := 0; i < len(s)-1; i++ {
	////	s1 := len(s) - 1
	////	s[i], s[s1-i] = s[s1-i], s[i]
	////	fmt.Println(s)
	////}
	//heapSort(*h)
	//fmt.Println(h)
	a := make([]int, 0, 10)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		r := rand.Intn(100)
		a = append(a, r)
	}
	fmt.Println(a)
	heaptest.HeapInit(a)
	fmt.Println(a)
	fmt.Println("-------------")
	heaptest.HeapSort(a)

}
