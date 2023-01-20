package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s1)

	rnd := random.Intn(100)

	println(rnd)

	var num int32
	println("type a number= ")
	fmt.Scan(&num)

	if num == int32(rnd) {
		fmt.Printf("random=%d,guess=%d,WINN 1000$", rnd, num)
	} else if sum(num) == sum(int32(rnd)) {
		fmt.Printf("random=%d,guess=%d,WINN 500$", rnd, num)
	} else {
		println("YOU LOOSE!!")
	}
}

func sum(num int32) int32 {
	first := num / 10
	second := num % 10

	return first + second
}
