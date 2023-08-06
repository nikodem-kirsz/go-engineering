package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

func main() {
	x := 10
	var y int = 5
	var z int = x + y
	fmt.Println(z)

	if x > 10 {
		fmt.Print("more than 10")
	}

	var a [5]int
	fmt.Println(a)

	a[2] = 7

	fmt.Println(a)
	//arrays
	b := []int{1, 2, 3, 4, 5}
	c := append(b, 5)
	fmt.Println(c)

	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "triangle")

	// alt + shift + up/down
	// command(windows) + shift + k
	numbers := []int{1, 2, 3}
	figures := append(make([]int, 1, 3), numbers...)
	// figures = append(figures, 2)
	// figures = append(figures, 2)
	// figures = append(figures, 2)
	fmt.Println("Figures:", figures)
	fmt.Println("Figures length", len(figures))

	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	//map range

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for index, value := range m {
		fmt.Println("index", index, "value", value)
	}

	n := []string{"a", "b", "c"}
	for index, value := range n {
		fmt.Println("index", index, "value", value)
	}
	result := sum(2, 3)
	fmt.Println(result)

	square, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(square)
	}

	p := person{name: "Niko", age: 16}
	fmt.Println(p)

	zz := 7
	inc(&zz)
	fmt.Println("zz:", zz)

	e := *getPointer()
	fmt.Println("getPointer", e)

	f := new(float64)
	*f = 234.0
	fmt.Println(f)
}

func getPointer() (myPointer *int) {
	a := 234
	return &a
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined or negative numbers")
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
