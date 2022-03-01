package main

import "time"

func main() {
	pinsResetAll()

	// displayCheckSegments()
	// time.Sleep(2 * time.Second)
	displayStartupAnimation()
	for i := 0; i < 16; i++ {
		displayDrawNumber(i)
		time.Sleep(666 * time.Millisecond)
	}
	pinsResetAll()
}
