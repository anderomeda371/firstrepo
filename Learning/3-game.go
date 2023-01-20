package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int32
	s1 := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s1)

	//println(i)
	firstrnd := random.Intn(50)

	println("type a number = ")
	fmt.Scan(&num)

	s2 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s2)
	secondrnd := random.Intn(50)

	if firstrnd > secondrnd {
		firstrnd, secondrnd = secondrnd, firstrnd
	}
	//println(firstrnd, secondrnd)
	count := 0
	validCheck := checkRange(int32(firstrnd), int32(secondrnd), num)

	for {
		count++
		if validCheck == 0 {
			fmt.Printf("rand1=%d , rand2=%d ,guess=%d \n", firstrnd, secondrnd, num)
			fmt.Printf("you guess right in %d times", count)

			break
		} else if validCheck == 1 {
			fmt.Printf("type a  number bigger ")
			fmt.Scan(&num)
		} else if validCheck == 2 {
			println("type a lower number  ")
			fmt.Scan(&num)

		}

		validCheck = checkRange(int32(firstrnd), int32(secondrnd), num)
	}

	//fmt.Printf("first=%d,second=%d", firstrnd, secondrnd)

}

//less return 1 ,more 2, range 0
func checkRange(first, second, guess int32) int32 {

	if first <= guess && second >= guess {
		return 0
	} else if first > guess {
		return 1
	} else {
		return 2
	}

}
