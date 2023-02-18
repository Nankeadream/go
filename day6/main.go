package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func (p Point) getX() int {
	fmt.Println("instance")
	return p.x
}

func (p *Point) getY() int {
	fmt.Println("pointer")
	return p.y
}

func main() {
	p := Point{4, 5}
	fmt.Println(p)
	fmt.Println(p.getX(), (&p).getX())
	fmt.Println("-----------")
	fmt.Println(p.getY(), (&p).getY())
}
