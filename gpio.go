package main

import "time"

func main() {
	pinsResetAll()

	// displayCheckSegments()
	// time.Sleep(2 * time.Second)
	displayStartupAnimation()
	for i := 0; i < 10; i++ {
		displayDrawNumber(i)
		time.Sleep(1 * time.Second)
	}
	pinsResetAll()
}
