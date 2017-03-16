// pass_by_value
package main

import (
	"fmt"
)

func main() {

	name := "kaushal"

	changeName(name)

	fmt.Println("name is still", name)

}

func changeName(name string) string {
	name = "kamesh"

	fmt.Println("name is changed to", name)

	return name
}
