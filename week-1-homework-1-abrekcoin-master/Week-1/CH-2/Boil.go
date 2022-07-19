// Boiling prints the boiling point of water.
package main
import "fmt"
const boilingF = 173.1
func main() {
var f = boilingF
var c = (f - 32) * 5 / 9
fmt.Printf("boiling point of alcohol = %g°F or %g°C\n", f, c)

// Output:
// boiling point = 212°F or 100°C
 

}