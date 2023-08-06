package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	a := int64(2)

	defer func(d int64) {
		fmt.Printf("d: %v", d)
	}(a)

	a = 3

	var d = int64(0)
	defer func(d *int64) {
		fmt.Printf("& %v Unix sec\n", *d)
	}(&d)
	fmt.Println("Done")
	d = time.Now().Unix()

	v := &Vertex{1, 2}
	(*v).X = 3
	v.Abs()
	v.Scale(5)
	fmt.Println(v)

	var r Shape = Rectange{Length: 3, Width: 4}
	fmt.Printf("Type of r: %T, Area: %v, Perimeter: %v.", r, r.Area(), r.Perimeter())
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectange struct {
	Length, Width float64
}

func (r Rectange) Area() float64 {
	return r.Length * r.Width
}

func (r Rectange) Perimeter() float64 {
	return 2 * (r.Length * r.Width)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
