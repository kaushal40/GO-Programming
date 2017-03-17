// nested_loop
package main

import (
	"fmt"
)

func main() {
	courseInProg := []string{"one", "two", "three"}
	courseCompleted := []string{"one", "four", "six"}

	for _, i := range courseInProg {
		fmt.Println(i)
		for _, j := range courseCompleted {
			if i == j {
				fmt.Println("course clash found")
			}
		}
	}

}
