package main

// SUPPORTED DIGITS WITH A DECIMAL POINT
// Single 8-segment screen support
var gpioDigitSingleDP = [][]bool{
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

// SUPPORTED DIGITS WITH A DECIMAL POINT
// Quad 8-segment screen support
var gpioDigitQuadDP = [][]bool{
	{true, true, true, true, false, true, true, true},     //	0
	{false, false, true, true, false, true, false, false}, //	1
	{true, true, true, false, true, true, false, true},    //	2
	{false, true, true, true, true, true, false, true},    //	3
	{false, false, true, true, true, true, true, false},   //	4
	{false, true, true, true, true, false, true, true},    //	5
	{true, true, true, true, true, false, true, true},     //	6
	{false, false, true, true, false, true, false, true},  //	7
	{true, true, true, true, true, true, true, true},      //	8
	{false, true, true, true, true, true, true, true},     //	9
	{true, false, true, true, true, true, true, true},     //	A
	{true, true, true, true, true, false, true, false},    //	B
	{true, true, true, false, false, false, true, true},   //	C
	{true, true, true, true, true, true, false, false},    //	D
	{true, true, true, false, true, false, true, true},    //	E
	{true, false, true, false, true, false, true, true}}   //	F

// SUPPORTED DIGITS
// Single 8-segment screen support
var gpioDigitSingle = [][]bool{
	{true, true, true, false, true, true, true, false},     //	0
	{false, false, true, false, true, false, false, false}, //	1
	{true, true, false, false, true, true, false, true},    //	2
	{false, true, true, false, true, true, false, true},    //	3
	{false, false, true, false, true, false, true, true},   //	4
	{false, true, true, false, false, true, true, true},    //	5
	{true, true, true, false, false, true, true, true},     //	6
	{false, false, true, false, true, true, false, false},  //	7
	{true, true, true, false, true, true, true, true},      //	8
	{false, true, true, false, true, true, true, true},     //	9
	{true, false, true, false, true, true, true, true},     //	A
	{true, true, true, false, false, false, true, true},    //	B
	{true, true, false, false, false, true, true, false},   //	C
	{true, true, true, false, true, false, false, true},    //	D
	{true, true, false, false, false, true, true, true},    //	E
	{true, false, false, false, false, true, true, true}}   //	F

// SUPPORTED DIGITS
// Quad 8-segment screen support
var gpioDigitQuad = [][]bool{
	{true, true, false, true, false, true, true, true},     //	0
	{false, false, false, true, false, true, false, false}, //	1
	{true, true, false, false, true, true, false, true},    //	2
	{false, true, false, true, true, true, false, true},    //	3
	{false, false, false, true, true, true, true, false},   //	4
	{false, true, false, true, true, false, true, true},    //	5
	{true, true, false, true, true, false, true, true},     //	6
	{false, false, false, true, false, true, false, true},  //	7
	{true, true, false, true, true, true, true, true},      //	8
	{false, true, false, true, true, true, true, true},     //	9
	{true, false, false, true, true, true, true, true},     //	A
	{true, true, false, true, true, false, true, false},    //	B
	{true, true, false, false, false, false, true, true},   //	C
	{true, true, false, true, true, true, false, false},    //	D
	{true, true, false, false, true, false, true, true},    //	E
	{true, false, false, false, true, false, true, true}}   //	F
