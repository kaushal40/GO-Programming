// pointers
package main

import (
	"fmt"
)

func main() {
	name := "kaushal"
	ptr := &name
	value := *ptr
	fmt.Println("address is=", ptr, " vlaue is=", value)
}
