/*
Race Condition is the condition of an result is dependent on the sequence or timing of other uncontrollable events.

It occurs when scheduling algorithm scheduling two or more threads, the threads can swap between execution time.

So, In the code I can't make sure that the first printed number is 1 (after the function addOne executed).
*/

package main

import (
	"fmt"
	"time"
)

func addOne(x *int) {
	*x += 1
}

func main() {
	x := 0
	// can not make sure the printed number is 1
	go addOne(&x)
	go fmt.Printf("First time print x (race condition): %d\n", x)

	// waiting the go routines finish
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Second time print x (expected value): %d\n", x)
}
