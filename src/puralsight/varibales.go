package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// declare variabale outside must use var keyword

var (
	name string = "kaushal"
	//automatic type assignment
	number  = 10.7
	number1 = 10.0
	number2 = 11
)

func main() {

	//sorthand operator works within function only
	myName := "kaushal"

	fmt.Println("hello frm", runtime.GOOS)
	fmt.Println("name is", name, " varible type is", reflect.TypeOf(name))
	fmt.Println("number is", number, " varible type is", reflect.TypeOf(number))
	//number1 + number2 will give an error
	fmt.Println("sum is", int(number1)+number2)
	fmt.Println("my name is", myName)
}
