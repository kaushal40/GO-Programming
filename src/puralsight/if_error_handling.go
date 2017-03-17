// if_error_handling
package main

import (
	"fmt"
	"os"
)

func main() {

	// _ means we do not care about the first return value so we can ignore it
	_, err := os.Open("c:\\temp\\test.txt")

	if err != nil {
		fmt.Println("Error Returned was:", err)
	}
}
