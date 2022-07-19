package main

import (
	"errors"
	"fmt"
)

// Define a new data type "Rectangle"
type Rectangle struct {
	Length, Width float32
}

// A area method for type "Rectangle"
func (r Rectangle) Circumference() float32 {
	return (r.Length * 2) + (2 * r.Width)
}

// A  circumference method for type "Rectangle"
func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

// Define an interface as achieve abstraction
type Area interface {
	Area() float32
}
type Circumference interface {
}

//Function controls value 0 or negative

func control(Length int, Width int) (int, error) {
	if Length*Width <= 0 {
		return -1, errors.New("Any value can not be 0 or Negative!")
	}
	return Length * Width, nil
}

func main() {

	var Length float32
	fmt.Print(" Enter the Length ")
	fmt.Scan(&Length)
	fmt.Print(" Enter the Width ")
	var Width float32
	fmt.Scan(&Width)
	// Declare and assign values to varaibles
	r := Rectangle{Length, Width}

	_, err := control(int(Width), int(r.Length))
	if err != nil {
		// Handle the error!
		fmt.Println(err)
	} else {
		// No errors!
		// Define a variable of type interface
		var a Area
		a = r
		fmt.Println("Area of Rectangle", a.Area())
		// circumfence
		fmt.Println("Circumference of Rectangle", r.Circumference())

	}

}
