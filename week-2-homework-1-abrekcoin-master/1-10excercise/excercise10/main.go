//on this file all excercise working together. İt may looks like complicated but i made this for shorten the hw

package main

import (
	"fmt"
	"time"
)

func main() {

	//10.1 and 10.2 and 10.3 and 10.4
	const (
		minsperday   = 60 * 24
		WeekDays     = 7
		hoursInDay   = 24
		hoursPerWeek = hoursInDay * WeekDays
		home         = "Milky Way Galaxy"
		lengt        = len(home)
		pi           = 3.14159265358979323846264
		tau          = pi * 2
		width        = 25
		height       = width * 2
		later        = 10               //10.6
		Nov          = 11 - (iota - 11) //10.7
		Oct
		Sep
		_ = iota - 14 //10.7
		Jan
		Feb
		Mar
		Spring = (iota - 17) * 3
		Summer
		Fall
		Winter
	)

	fmt.Printf("minutes in 2 weeks is %d \n", minsperday*WeekDays*2) //10.1
	fmt.Printf("There are %d hours in 2 weeks.\n", hoursPerWeek*5)   //10.2
	fmt.Printf("Ther are %d char in %q \n", lengt, home)             //10.3
	fmt.Printf("tau = %g\n", tau)                                    //10.4
	fmt.Printf("area = %d\n", width*height)                          // 10.5

	//10.6
	hours, _ := time.ParseDuration("1h")
	fmt.Printf("%s later...\n", hours*later)
	fmt.Println(Sep, Oct, Nov)                //10.7 it changes because of my code . but i got it :)
	fmt.Println(Jan, Feb, Mar)                //10.8
	fmt.Println(Winter, Spring, Summer, Fall) // 10.9

}
