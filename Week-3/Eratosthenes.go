package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter the upper limit for Sieve of Eratosthenes calculation: ")
	var value int
	fmt.Scan(&value)

	// checks the input string or zero value.
	if value == 0 {
		fmt.Print("please enter valid number. not string or zero :) ")
	} else {
		sieve(value)
	}
}

//Shows the pime numbers after calculation
func Primes(num int, prime []bool) {
	fmt.Printf("Primes less than %d : ", num)
	for i := 2; i <= num; i++ {
		if prime[i] == false {
			fmt.Printf("%d ", i)
		}
	}
}

//calculates eratoshenes
func sieve(num int) {

	prime := make([]bool, num+1)

	// initialize everything with false first(not crossed)
	for i := 0; i < num+1; i++ {
		prime[i] = false
	}
	//make prime numbers true
	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true // cross
			}
		}

	}

	Primes(num, prime)

}
