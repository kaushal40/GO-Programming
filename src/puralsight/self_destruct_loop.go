// self_destruct_loop
package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("detruct ship")
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}
