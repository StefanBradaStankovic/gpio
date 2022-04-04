package main

import "log"

var allPins = [...]string{Pin01, Pin02, Pin03, Pin04, Pin05, Pin06, Pin07, Pin08, Pin09, Pin10, Pin11, Pin12}
var displaySegmentPins = [...]string{Pin01, Pin02, Pin03, Pin04, Pin05, Pin07, Pin10, Pin11}
var displayGroundPins = [...]string{Pin06, Pin09, Pin08, Pin12}

var digits = []int{0, 0, 0, 0}
var mainFunctionIsFinished = 0

//  var circlePins = [...]string{Pin01, Pin02, Pin04, Pin06, Pin07, Pin09}

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

//  Split an integer into an array with 4 digits
func splitDigits(number int) []int {

	var digits = []int{0, 0, 0, 0}
	if number >= 0 && number <= 9999 {

		digits[3] = number % 10
		digits[2] = (number / 10) % 10
		digits[1] = (number / 100) % 10
		digits[0] = number / 1000
	} else {
		log.Println("Number out of bounds !!!")
	}

	return digits
}
