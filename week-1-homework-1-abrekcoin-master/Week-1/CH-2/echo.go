// Echo4 prints its commandline
package main
import (
"flag"
"fmt"
"strings"
)
var n = flag.Bool("n", true, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
	fmt.Println()
	}

// experimental zone of Section 2
	p := new(int) // p, of type *int, points to an unnamed int variable
	fmt.Println(*p) // "0"
	*p = 2 // sets the unnamed int to 2
	fmt.Println(*p) // "2"
	
	e := new(int)
	d := new(int)
	fmt.Println(*e, *d) // "false"
	fmt.Println(e, d) // "false"
	}