package main

import "fmt"

func main() {
	print("   ")
	for c := 1; c < 10; c++ {
		fmt.Printf(" %3d ", c)

	}
	print("\n")
	for c := 1; c < 18; c++ {
		fmt.Printf(" %s ", "-")

	}

	print("\n")
	for i := 1; i < 10; i++ {
		fmt.Printf("%d |", i)
		for j := 1; j < 10; j++ {

			fmt.Printf(" %3d ", i*j)
		}
		print("\n")
	}
}
