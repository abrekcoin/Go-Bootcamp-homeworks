package main

import "fmt"

func main() {

	// defining integer array
	fmt.Println("Plese enter the size of array")
	var n int
	fmt.Scan(&n)
	set(n)
}

//defing array and calling rotate func
func set(n int) {
	fmt.Println("Plese enter array values")
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	fmt.Println("Your array is  ", numbers)

	// calling a function to rotate an array
	fmt.Println("Plese enter rotate number value")
	var rotate int
	fmt.Scan(&rotate)
	leftRotate(numbers, rotate)

	fmt.Println("After Rotate:", numbers)
}

// function rotate based on the count
func leftRotate(arr []int, count int) {
	for i := 0; i < count; i++ {
		leftRotatebyOne(arr) // calling function to perform rotation ont time
	}
}

// function perform left rotate only one time
func leftRotatebyOne(arr []int) {

	var i int = 0

	var temp int = arr[0]

	for ; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	arr[i] = temp
}
