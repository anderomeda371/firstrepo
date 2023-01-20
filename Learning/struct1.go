package main

import "fmt"

func main() {

	type myType float64

	var n myType = 14.9
	fmt.Printf("%f,%T", n, n)

	var m float64
	m = float64(n)
	fmt.Printf("\n%f,%T", m, m)

}
