package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

type shape interface{
	area() float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r rect) area() float64{
	return r.width*r.height
}

func main(){

	c1 := circle{radius: 10}
	r1 := rect{width: 10, height: 10}
    shapes := []shape{c1,r1}
	fmt.Println(shapes[0].area())
}