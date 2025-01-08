package main

import(
	 "fmt"
	"math"
	)

type shape interface {

	area() float64

}

type circle struct {
	radius float64
}

type recta struct {
	height float64
	width float64
}

func (c *circle) area() float64{
	return math.Pi * c.radius *c.radius 
}

func (r *recta) area() float64{

	return r.width * r.height 
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{4.5}
	r1:=recta{4,9}
	s1:=[]shape{&c1,&r1}
	for _, val := range s1{
		fmt.Println(getArea(val))
	}
}
