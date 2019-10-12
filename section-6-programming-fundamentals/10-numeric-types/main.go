package main

import (
	"fmt"
	"math"
	"runtime"
)

// int and float64 are the most commonly used:
var w int64
var x int
var y float64
var z int8 = -128

func main() {
	x = 42
	// x = 2.3456 Doesn't wor, truncation from float to int
	y = 42.3456
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", z)
	// z = -129 Won't work, limit is -128
	// z - 128 Won't work, limit is 127
	fmt.Println("Some runtime related env variables:")
	fmt.Printf("runtime.GOOS: %v\n", runtime.GOOS)
	fmt.Printf("runtime.GOARCH: %v\n", runtime.GOARCH)
	fmt.Printf("math.MaxInt64: %v", math.MaxInt64)
	w = 9223372036854775807
	fmt.Printf("w: %v", w)
	fmt.Printf("math.MaxInt32 %v", math.MaxInt32)
}
