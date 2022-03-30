package main

import (
	"time"

	"github.com/StefanBradaStankovic/dispcon"
)

func main() {

	dispcon.PinsResetIn(allPins[:])
	time.Sleep(1 * time.Second)

	go dispcon.DisplayDrawNumberMultiple(gpioPins[:], sellectorPins[:], gpioDigitQuad, &digits, &mainDone)
	time.Sleep(2 * time.Second)

	for i := 0; i < 10000; i += 9 {

		digits = dispcon.SplitDigits(i)
		time.Sleep(100 * time.Millisecond)
		if i == 1500 {
			i = 0
		}
	}

	time.Sleep(5 * time.Second)

	mainDone = 1

	dispcon.PinsResetLow(allPins[:])
	dispcon.PinsResetIn(allPins[:])
}
