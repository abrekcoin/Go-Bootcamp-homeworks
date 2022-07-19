package main
import (
	"fmt"
	"os"
)
func main() {
	//ex.7.1
	fmt.Printf("my age is %d \n",30)
	//ex7.2
	fmt.Printf("my surname is %[2]s and my name is %[1]s \n","faik","sevim")
	mesaj := "my name is %s but every body calls me %s. \n"
	fmt.Printf(mesaj,"geovanni georgo","georco")
	//ex 7.3
	ft := false
	fmt.Printf("these are %v claims \n",ft)
	//ex 7.4
	temp := 29.5
	fmt.Printf("weather is %.1f celcius degree \n",temp)
	//ex 7.5
	fmt.Printf("\"hello world\"\n")
	fmt.Printf("%q\n", "hello world")
	//ex 7.6
	three := 3
	fmt.Printf("type of %d is %[1]T \n",three)
	//ex 7.7
	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)
	//ex 7.8
	fmt.Printf("Type of %s is %[1]T\n", "hello")
	//ex 7.9
	fmt.Printf("Type of %t is %[1]T\n", true)
	//ex 7.10
	name, lastname := os.Args[1], os.Args[2]
	msg := "Yor Name İs %s surname is %s \n"
	fmt.Printf(msg,name,lastname)
	//go run "c:\Users\faiks\OneDrive\Masaüstü\week-2\1-10excercise\excercise7\main.go" faik sevim
	//output = Yor Name İs faik surname is sevim
}
