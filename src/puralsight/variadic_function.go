// variadic_function
package main

import (
	"fmt"
)

func main() {
	bestFinish := bestLeagueFinish(10, 30, 20, 15, 16)
	fmt.Println(bestFinish)
}

func bestLeagueFinish(finishesh ...int) int {

	best := finishesh[0]

	for _, i := range finishesh {
		if i < best {
			best = i
		}
	}

	return best
}
