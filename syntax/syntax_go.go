package syntax

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

// func main() {
// 	//assingment
// 	x := 10
// 	var y int = 5
// 	var sum int = x + y
// 	fmt.Println(sum)

// 	if x > 10 {
// 		fmt.Print("more than 10")
// 	}

// 	var a [5]int
// 	fmt.Println(a)

// 	a[2] = 7

// 	fmt.Println(a)
// 	//arrays
// 	b := []int{1, 2, 3, 4, 5}
// 	c := append(b, 5)
// 	fmt.Println(c)

// 	vertices := make(map[string]int)
// 	vertices["triangle"] = 2
// 	vertices["square"] = 3
// 	vertices["dodecagon"] = 12

// 	delete(vertices, "triangle")

// 	fmt.Println(vertices)

// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}

// 	j := 0
// 	for j < 5 {
// 		fmt.Println(j)
// 		j++
// 	}
// 	//map range

// 	m := make(map[string]string)
// 	m["a"] = "alpha"
// 	m["b"] = "beta"

// 	for index, value := range m {
// 		fmt.Println("index", index, "value", value)
// 	}

// 	n := []string{"a", "b", "c"}
// 	for index, value := range n {
// 		fmt.Println("index", index, "value", value)
// 	}
// 	result := su2m(2, 3)
// 	fmt.Println(result)

// 	square, err := sqrt(-16)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(square)
// 	}

// 	p := person{name: "NIko", age: 16}
// 	fmt.Println(p)

// 	z := 7
// 	inc(&z)
// 	fmt.Println(z)

// 	e := *getPointer()
// 	fmt.Println(e)

// 	f := new(float64)
// 	*f = 234.0
// 	fmt.Println(f)

// }

func main() {
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

func getPointer() (myPointer *int) {
	a := 234
	return &a
}

func su2m(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined or negative numbers")
	}
	return math.Sqrt(x), nil
}

type person struct {
	name string
	age  int
}

func inc(x *int) {
	fmt.Println(reflect.TypeOf(x))
	*x++
}
