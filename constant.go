package main

import (
	"fmt"
	"math"
)

const s string = "Test"

func main() {
	fmt.Println(s)

	const n = 500000
	const d = 3e20

	fmt.Println(d)
	fmt.Println(math.Sin(n))

}
