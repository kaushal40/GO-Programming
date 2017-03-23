package main

import (
	"fmt"
	"math"
)

//class without struct type
type MyFloat float64

//use that as receiver to the method
//you can use class(receiver) only if it is declared in the same package
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
