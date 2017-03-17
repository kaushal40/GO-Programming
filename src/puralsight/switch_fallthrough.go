// switch_fallthrough
package main

import (
	"fmt"
)

func main() {
	switch "docker" {
	case "linux":
		fmt.Println("learn linux courses")
	case "docker":
		fmt.Println("learn docker courses")
		//fallthrough will run next switch block as well
		fallthrough
	case "windows":
		fmt.Println("learn windows courses")
	default:
		fmt.Println("you can learn any course")
	}
}
