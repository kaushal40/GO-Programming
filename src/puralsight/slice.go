// slice
package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mySlice[1] = 0

	sliceOfSlice := mySlice[2:5]

	fmt.Println(sliceOfSlice)

	//append if goes out of the underlying array lenght it will expand the backing array by double the size

	newSlice := make([]int, 1, 4)

	for i := 1; i < 17; i++ {
		newSlice = append(newSlice, i)
		fmt.Println("capacity is: %d", cap(newSlice))
	}

	//slice can be appended to slice using ellipses
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

}
