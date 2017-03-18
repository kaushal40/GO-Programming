// concurrency
package main

import (
	"fmt"
	"time"
	//sync function for waiting to routines to finish
	"sync"
)

func main() {

	// to add two routines to wait group
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {

		defer waitGrp.Done()
		//as this routine goes to sleep the other routine will be executed on the thread
		time.Sleep(5 * time.Second)
		fmt.Println("hello")
	}()

	go func() {

		defer waitGrp.Done()
		fmt.Println("second hello")
	}()

	waitGrp.Wait()
}
