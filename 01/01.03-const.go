package main

import (
	"fmt"
	"time"
)

func main() {
	type Direction int

	const (
		North Direction = iota
		East
		South
		West
	)
	//var (
	//	age int
	//	name string
	//	balance []float32
	//)
	const pi = 3.14159
	//const (
	//	e = 2.7182818
	//	pi = 3.1415926
	//)
	const noTime time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noTime)      // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	//var x float32 = math.Pi
	//var y float64 = math.Pi
	//var z complex128 = math.Pi

	//const Pi64 float64 = math.Pi
	//var x float32 = float32(Pi64)
	//var y float64 = Pi64
	//var z complex128 = complex128(Pi64)
}
