// go_maps
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	//map is unordered
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}

	for key, value := range testMap {
		fmt.Println("key is = ", key, "value is = ", value)
	}

	//manipulating map

	//update existing value
	testMap["A"] = 5

	fmt.Println(testMap)

	//add new value

	testMap["F"] = 6

	fmt.Println(testMap)

	//delete a value

	delete(testMap, "F")

	fmt.Println(testMap)

}
