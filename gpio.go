package main

import (
	"time"

	"github.com/StefanBradaStankovic/dispcon"
)

func main() {

	dispcon.PinsResetIn(allPins[:])
	time.Sleep(1 * time.Second)

	go DisplayDrawNumberMultiple(gpioPins[:], gpioDigitQuad)
	time.Sleep(2 * time.Second)

	for i := 0; i < 10000; i += 9 {

		SplitDigits(i)
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(5 * time.Second)

	MainDone = 1

	dispcon.PinsResetLow(allPins[:])
	dispcon.PinsResetIn(allPins[:])
}
