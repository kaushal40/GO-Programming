package main

import (
	"fmt"
	"math"
)

//class
type Vertex struct {
	X, Y float64
}

//class method: pass receiver(class) between function and method name
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//creating class instance and using it
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
