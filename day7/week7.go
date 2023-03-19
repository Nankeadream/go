package main

type Point struct {
	x, y int
}

func (p Point) getX() int {
	return p.x
}

//func main() {
//	//var p1 = Point{1, 2}
//	//var p2 = Point{3, 4}
//	//fmt.Printf("%T,%[2]p,%[1]d\n", p1, &p1)
//	//fmt.Printf("%T,%[2]p,%[1]d\n", p2, &p2)
//	//fmt.Println(p2.getX(), p1.getX())
//	//fmt.Printf("%p,%p", p1.getX(), p2.getX())

//}
