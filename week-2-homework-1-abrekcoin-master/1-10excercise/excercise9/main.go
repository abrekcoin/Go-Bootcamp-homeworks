package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: The Type Problem
//
//  Solve the data type problem in the program.
//
// EXPECTED OUTPUT
//  width: 265 height: 265
//  are they equal? true
// ---------------------------------------------------------

func main() {
	//9.2
	var (
		width  uint16
		height uint16
	)

	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)
	fmt.Println("are they equal?", width == height)
	//9.4
	t, _ := time.ParseDuration("1h30m")
	multiplier, _ := strconv.ParseInt(os.Args[1], 10, 64)
	t *= time.Duration(multiplier)
	fmt.Println(t)
	/* //İNPUT AND AUT PUT
	PS C:\Users\faiks\OneDrive\Masaüstü\learngo-master> go run
	"c:\Users\faiks\OneDrive\Masaüstü\learngo-master\09-go-type-system\exercises\04-time-multiplier\solution\main.go" 2
	3h0m0s	 */
	//9.5
	type (
		Feeet  float64
		Meters float64
	)

	var (
		feet   Feeet
		meters Meters
	)

	arg := os.Args[1]

	// feet is a Feet value now
	val, _ := strconv.ParseFloat(arg, 64)
	feet = Feeet(val)

	// meters is a Meters value now
	meters = Meters(feet * 0.3048)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)

	//9.6

	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	celsius *= Celsius(float64(celsiusDegree) * factor)
	fahr *= Fahrenheit(float64(fahrDegree) * factor)

	fmt.Println(celsius, fahr)

}
