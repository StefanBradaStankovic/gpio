package main

import (
	"time"

	"github.com/StefanBradaStankovic/dispcon"
)

func main() {

	dispcon.DiodeArrayIn(allPins[:])
	time.Sleep(1 * time.Second)

	go dispcon.DisplayMultipleDigits(displaySegmentPins[:], displayGroundPins[:], diodeState_QuadDigit, &digits, &mainFunctionIsFinished)
	time.Sleep(2 * time.Second)

	for i := 0; i < 10000; i += 29 {

		digits = splitDigits(i)
		time.Sleep(100 * time.Millisecond)
		if i == 1500 {
			i = 0
		}
	}

	time.Sleep(5 * time.Second)

	mainFunctionIsFinished = 1

	dispcon.DiodeArrayIn(allPins[:])
}
