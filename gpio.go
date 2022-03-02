package main

import (
	"time"
)

func main() {

	pinsResetAll(gpioPins[:])
	// displayCheckSegments()
	// time.Sleep(2 * time.Second)

	for i := 0; i < 16; i++ {
		displayDrawNumber(gpioPins, i)
		time.Sleep(666 * time.Millisecond)
	}

	pinsResetAll(gpioPins[:])

}
