package main

import (
	movie "Learning/types"
	"fmt"
)

func main() {

	m1 := movie.Movie{"Gamer", "David Lynch", "1990"}
	//fmt.Printf("%v",m1.)
	fmt.Printf("%v", *m1.CreateUpperInfo())

	changeName(&m1)
	fmt.Printf("\n %v", *m1.CreateUpperInfo())
}

func changeName(m *movie.Movie) {

	m.Name = "khafan "

}
