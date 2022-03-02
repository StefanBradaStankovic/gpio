package main

import (
	"log"
	"time"
)

//
//
//

// Defining gpio.High() parameters for digits 0 1 2 3 4 5 6 7 8 9 A B C D E F
var gpioDigit = [16][8]bool{
	{true, true, true, true, true, true, true, false},     //	0
	{false, false, true, true, true, false, false, false}, //	1
	{true, true, false, true, true, true, false, true},    //	2
	{false, true, true, true, true, true, false, true},    //	3
	{false, false, true, true, true, false, true, true},   //	4
	{false, true, true, true, false, true, true, true},    //	5
	{true, true, true, true, false, true, true, true},     //	6
	{false, false, true, true, true, true, false, false},  //	7
	{true, true, true, true, true, true, true, true},      //	8
	{false, true, true, true, true, true, true, true},     //	9
	{true, false, true, true, true, true, true, true},     //	A
	{true, true, true, true, false, false, true, true},    //	B
	{true, true, false, true, false, true, true, false},   //	C
	{true, true, true, true, true, false, false, true},    //	D
	{true, true, false, true, false, true, true, true},    //	E
	{true, false, false, true, false, true, true, true}}   //	F

// Draw an ordinal number from 0 to 9
func displayDrawNumber(sellectedPins [8]string, number int) {

	pinsResetAll(sellectedPins[:])

	if number >= 0 && number <= 15 {
		diodeOnCluster(sellectedPins, gpioDigit[number])
	} else {
		log.Println("Invalid number")
	}

}

//
//
// D I S P L A Y     A N I M A T I O N S
//
//

// Flash every segment to visually confirm they all work
func displayCheckSegments(sellectedPins [8]string) {

	for i := 0; i < len(sellectedPins); i++ {
		diodeFlash(sellectedPins[i], 100, 400)
	}

}

// Draw a running circle animation
func displayCircleRound(sellectedPins [8]string, speed int) {

	for i := 0; i < len(sellectedPins); i++ {
		diodeFlash(sellectedPins[i], time.Duration(speed), 0)
	}

}

// Draw a full circle animation
func displayCircleFlash(sellectedPins [8]string, speed int) {

	for i := 0; i < len(sellectedPins); i++ {
		diodeOn(sellectedPins[i])
	}

	time.Sleep(time.Duration(speed) * time.Millisecond)

	for i := 0; i < len(sellectedPins); i++ {
		diodeOff(sellectedPins[i])
	}

	time.Sleep(time.Duration(speed) * time.Millisecond)

}

// Draw a full circle segment by segment
func displayLoading(sellectedPins [8]string, speed int) {

	for i := 0; i < len(sellectedPins); i++ {
		diodeOn(sellectedPins[i])
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}

}

// Draw a startup animation
func displayStartupAnimation(sellectedPins [8]string) {

	displayLoading(sellectedPins, 150)
	displayCircleFlash(sellectedPins, 250)
	displayCircleFlash(sellectedPins, 250)
	displayCircleFlash(sellectedPins, 1500)
}
