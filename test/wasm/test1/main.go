package main

import (
	"fmt"
	//"syscall/js"
	"time"
)

func main() {
	//v := js.ProbeWASM()
	now := time.Now()
	fmt.Println(now)
	// fmt.Println("Hello world")
	// v := run()
	// fmt.Println(v)
}

func run() int {
	v := 1
	return v + 1
}
