package main


import (
	"fmt"
)


func main() {

	fmt.Print(" Enter the Value for prime factorization ")
	var value int
	fmt.Scan(&value)

	// checks the input string or zero value.
	if value == 0 {
		fmt.Print("please enter valid number. not string or zero :) ")
	} else {
		fmt.Println(PrimeFactors(value))
	}
}

// Get all prime factors of a given number value
func PrimeFactors(value int) (pfs []int) {
	//Even number calculation
	for value%2 == 0 {
		pfs = append(pfs, 2)
		value = value / 2
	}

	//Odd number calculation. Number must be the odd in this part.
	for i := 3; i*i <= value; i = i + 2 {
		// while i divides value, append i and divide value
		for value%i == 0 {
			pfs = append(pfs, i)
			value = value / i
		}
	}

	// This condition is to handle the case when given value is a prime number
	if value > 2 {
		pfs = append(pfs, value)
	}

	return
}

