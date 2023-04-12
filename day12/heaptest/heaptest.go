package heaptest

import "fmt"

func HeapSort(a []int) {
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		Heapify(a, i, 0)
		fmt.Println(a)
	}

}

func HeapInit(a []int) {
	n := len(a)
	for i := n / 2; i >= 0; i-- {
		Heapify(a, n, i)
	}
}

func Heapify(a []int, n, i int) {
	largest := i     //最后一个父节点的索引
	left := 2*i + 1  //左节点
	right := 2*i + 2 //右节点
	if left < n && a[left] > a[largest] {
		largest = left
	}
	if right < n && a[right] > a[largest] {
		largest = right
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		Heapify(a, n, largest)
	}
}
