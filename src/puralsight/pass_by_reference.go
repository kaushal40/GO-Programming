// pass_by_reference
package main

import (
	"fmt"
)

func main() {

	name := "kaushal"

	changeName(&name)

	fmt.Println("name is still", name)

}

func changeName(adressOfName *string) string {
	fmt.Println("addressOfName is", adressOfName)

	*adressOfName = "kamesh"

	fmt.Println("name is changed to", *adressOfName)

	return *adressOfName
}
