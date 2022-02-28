package main

import (
	"log"
	"time"
)

var gpioNumber0 = [...]bool{true, true, true, true, true, true, true, false}
var gpioNumber1 = [...]bool{false, false, true, true, true, false, false, false}
var gpioNumber2 = [...]bool{true, true, false, true, true, true, false, true}
var gpioNumber3 = [...]bool{false, true, true, true, true, true, false, true}
var gpioNumber4 = [...]bool{false, false, true, true, true, false, true, true}
var gpioNumber5 = [...]bool{false, true, true, true, false, true, true, true}
var gpioNumber6 = [...]bool{true, true, true, true, false, true, true, true}
var gpioNumber7 = [...]bool{false, false, true, true, true, true, false, false}
var gpioNumber8 = [...]bool{true, true, true, true, true, true, true, true}
var gpioNumber9 = [...]bool{false, true, true, true, true, true, true, true}

// Draw an ordinal number from 0 to 9
func displayDrawNumber(number int) {
	pinsResetAll()
	switch number {
	case 0:
		diodeOnCluster(gpioPins, gpioNumber0)
	case 1:
		diodeOnCluster(gpioPins, gpioNumber1)
	case 2:
		diodeOnCluster(gpioPins, gpioNumber2)
	case 3:
		diodeOnCluster(gpioPins, gpioNumber3)
	case 4:
		diodeOnCluster(gpioPins, gpioNumber4)
	case 5:
		diodeOnCluster(gpioPins, gpioNumber5)
	case 6:
		diodeOnCluster(gpioPins, gpioNumber6)
	case 7:
		diodeOnCluster(gpioPins, gpioNumber7)
	case 8:
		diodeOnCluster(gpioPins, gpioNumber8)
	case 9:
		diodeOnCluster(gpioPins, gpioNumber9)
	default:
		log.Print("Invalid number !")
	}
}

//
//
// +-----+  DISPLAY ANIMATION +-----+
//
//

func displayCheckSegments() {
	diodeFlash(Pin01, 100, 400)
	diodeFlash(Pin02, 100, 400)
	diodeFlash(Pin04, 100, 400)
	diodeFlash(Pin05, 100, 400)
	diodeFlash(Pin06, 100, 400)
	diodeFlash(Pin07, 100, 400)
	diodeFlash(Pin09, 100, 400)
	diodeFlash(Pin10, 100, 400)
}

// Draw a running circle animation
func displayCircleRound(speed int) {
	diodeFlash(Pin01, time.Duration(speed), 0)
	diodeFlash(Pin02, time.Duration(speed), 0)
	diodeFlash(Pin04, time.Duration(speed), 0)
	diodeFlash(Pin06, time.Duration(speed), 0)
	diodeFlash(Pin07, time.Duration(speed), 0)
	diodeFlash(Pin09, time.Duration(speed), 0)
}

// Draw a full circle animation
func displayCircleFlash(speed int) {
	diodeOn(Pin01)
	diodeOn(Pin02)
	diodeOn(Pin04)
	diodeOn(Pin06)
	diodeOn(Pin07)
	diodeOn(Pin09)
	time.Sleep(time.Duration(speed) * time.Millisecond)
	diodeOff(Pin01)
	diodeOff(Pin02)
	diodeOff(Pin04)
	diodeOff(Pin06)
	diodeOff(Pin07)
	diodeOff(Pin09)
	time.Sleep(time.Duration(speed) * time.Millisecond)
}

// Draw a full circle segment by segment
func displayLoading(speed int) {
	diodeOn(Pin01)
	time.Sleep(time.Duration(speed) * time.Millisecond)
	diodeOn(Pin02)
	time.Sleep(time.Duration(speed) * time.Millisecond)
	diodeOn(Pin04)
	time.Sleep(time.Duration(speed) * time.Millisecond)
	diodeOn(Pin06)
	time.Sleep(time.Duration(speed) * time.Millisecond)
	diodeOn(Pin07)
	time.Sleep(time.Duration(speed) * time.Millisecond)
	diodeOn(Pin09)
	time.Sleep(time.Duration(speed) * time.Millisecond)
}

// Draw a startup animation
func displayStartupAnimation() {
	displayLoading(150)
	displayCircleFlash(250)
	displayCircleFlash(250)
	displayCircleFlash(1500)
}
