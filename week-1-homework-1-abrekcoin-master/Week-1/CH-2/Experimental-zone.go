// this file include the experimental part of book. This includes the example parts. not the excersize parts.

package main

import "fmt"

func main() {
	p := new(int)   // p, of type *int, points to an unnamed int variable
	fmt.Println(*p) // "0"
	*p = 2          // sets the unnamed int to 2
	fmt.Println(*p) // "2"

	/////////////2.3.3////
	e := new(int)
	d := new(int)
	fmt.Println(e == d) // "false"
	fmt.Println(*e, *d) // "0-0"
	fmt.Println(e, d)   // "address"
	//////////////2.4.1//////
	fmt.Println(gcd(4, 16))
	fmt.Println(fib(7))

}

//////////////2.4.1//////
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

//////////////2.4.1//////
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
