package main

import (
	"fmt"
	"math"
	"math/rand"
)


func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(29))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
}








