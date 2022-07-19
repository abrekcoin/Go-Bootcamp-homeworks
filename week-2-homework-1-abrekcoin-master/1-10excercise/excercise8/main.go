package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	//8.1
	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * 0.5)
	fmt.Println(50 / 0.5)
	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))
	//8.2
	x := float64(5) / 2
	fmt.Println(x)
	fmt.Println(10 + 5 - (5 - 10))
	fmt.Println(-10 + 0.5 - (1 + 5.5))

	// 8.3
	fmt.Println(5 + 10*(2-5))
	fmt.Println(0.5 * (2 - 1))
	fmt.Println((3+1)/2*10 + 4)
	fmt.Println(10 / 2 * (10 % 7))
	fmt.Println(100 / (5. / 2))

	//8.4

	counter, factor := 45, 0.5

	counter++
	counter++
	counter++
	counter++
	counter++
	factor--
	factor--

	fmt.Println(float64(counter) * factor)
	//8.5
	var counter2 int

	counter2++
	counter2--
	counter2 += 5
	counter2 *= 10
	counter2 /= 5
	fmt.Println(counter2)

	//8.6
	width, height := 10, 2

	/* width = width + 1
	width = width + height
	width = width - 1
	width = width - height
	width = width * 20
	width = width / 25
	width = width % 5 */

	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)

	//8.7

	var (
		radius = 10.
		area   float64
	)
	area = math.Pi * radius * radius

	fmt.Printf("radius: %g -> area: %g\n", radius, area)

	//8.8

	var (
		spradius, sparea float64
	)

	spradius, _ = strconv.ParseFloat(os.Args[1], 64)

	sparea = 4 * math.Pi * spradius * spradius
	// sparea = 4 * math.Pi * math.Pow(radius, 2)

	fmt.Printf("Radius is %1.f Area is %1.f", spradius, sparea)
	/////////////example output = Radius is 34 Area is 14527/////

	//8.9

	var spvradius, spvvol float64

	spvradius, _ = strconv.ParseFloat(os.Args[1], 64)

	spvvol = (4 * math.Pi * math.Pow(spvradius, 3)) / 3

	fmt.Printf("radius: %g -> volume: %.2f\n", spvradius, spvvol)
	//output radius: 34 -> volume: 164636
	//8.2.1
	json := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`
	fmt.Println(json)

	// 8.2.4 //it print thr lengt of given string (even if include turkish characters)
	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)

	//8.2.6
	fmt.Println(strings.ToLower(os.Args[1]))

	//8.2.7 sPACE TRİMMER

	fmt.Println(strings.TrimSpace(json))
	//
}
