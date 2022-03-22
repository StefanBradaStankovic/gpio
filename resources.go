package main

import (
	"log"
	"time"

	"github.com/StefanBradaStankovic/dispcon"
)

var allPins = [...]string{Pin01, Pin02, Pin03, Pin04, Pin05, Pin06, Pin07, Pin08, Pin09, Pin10, Pin11, Pin12}
var gpioPins = [...]string{Pin01, Pin02, Pin03, Pin04, Pin05, Pin07, Pin10, Pin11}
var sellectorPins = [...]string{Pin06, Pin09, Pin08, Pin12}

var MainDone = 0

// var circlePins = [...]string{Pin01, Pin02, Pin04, Pin06, Pin07, Pin09}

const (
	Pin01 = "6"
	Pin02 = "13"
	Pin03 = "19"
	Pin04 = "26"
	Pin05 = "5"
	Pin06 = "24" // GND
	Pin07 = "16"
	Pin08 = "18" // GND
	Pin09 = "23" // GND
	Pin10 = "20"
	Pin11 = "21"
	Pin12 = "22" // GND
)

var Digits = []int{0, 0, 0, 0}

func DisplayDrawNumberMultiple(sellectedPins []string, sellectedDigits [][]bool) {

	dispcon.PinsResetLow(gpioPins[:])
	dispcon.PinsResetIn(sellectorPins[:])

	for MainDone != 1 {

		for i := 0; i < 4; i++ {
			dispcon.DiodeOff(sellectorPins[i])
			dispcon.DisplayDrawNumber(sellectedPins, sellectedDigits, Digits[i])
			time.Sleep(1 * time.Millisecond)
			dispcon.DiodeIn(sellectorPins[i])
		}
	}
}

func SplitDigits(number int) {

	if number >= 0 && number <= 9999 {

		if number < 10 {
			Digits[3] = number
			Digits[2] = 0
			Digits[1] = 0
			Digits[0] = 0
		}

		if number >= 10 && number < 100 {
			Digits[3] = number % 10
			Digits[2] = number / 10
			Digits[1] = 0
			Digits[0] = 0
		}

		if number >= 100 && number < 1000 {
			Digits[3] = number % 10
			Digits[2] = (number / 10) % 10
			Digits[1] = number / 100
			Digits[0] = 0
		}

		if number >= 1000 {
			Digits[3] = number % 10
			Digits[2] = (number / 10) % 10
			Digits[1] = (number / 100) % 10
			Digits[0] = number / 1000
		}
	} else {
		log.Println("Number out of bounds !!!")
	}
}
