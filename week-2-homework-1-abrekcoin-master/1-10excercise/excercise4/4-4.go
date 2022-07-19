package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.NumCPU() is a call expression
	fmt.Println("Bilgisayarınız" ,runtime.NumCPU()," Çekirdeklidir")
	fmt.Println(runtime.Version(),	"Go versyionu")
}
