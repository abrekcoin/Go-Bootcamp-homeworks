package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// creating random number
	var randNum int
	CreateNum(&randNum)
	Randslice := make([]int, 4) // keep every digit of four digits
	GetNum(Randslice, randNum)  // slice is passed by reference
	OnGame(Randslice)

}

func CreateNum(p *int) {
	//Set seed
	rand.Seed(time.Now().UnixNano())

	var num int
	for {
		num = rand.Intn(10000) // four digit
		if num >= 1000 {
			break
		}
	}
	//fmt.Println("num=",num)
	*p = num
}
func GetNum(s []int, num int) {
	s[0] = num / 1000       // thousand bits
	s[1] = num % 1000 / 100 // take hundreds
	s[2] = num % 100 / 10   // take ten
	s[3] = num % 10         // take bits
}
func OnGame(randSlice []int) {
	var num int
	keySlice := make([]int, 4)
	for {
		for {
			fmt.Printf("please enter a four digit number:")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("the number you entered does not meet the requirements... please enter number between 999 and 100000")
		}
		// slice input number
		GetNum(keySlice, num)
		n := 0
		//check the bits
		for i := 0; i < 4; i++ {
			if keySlice[i] == randSlice[i] {
				fmt.Printf("+", i+1)
			} else if keySlice[i] != randSlice[i] {
				fmt.Printf("-", i+1)
			} else {
				fmt.Printf("Guessed correctly", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("Congratulations! All right! ")
			// jump out of the outermost loop
		}
	}
}
