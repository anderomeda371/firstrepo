package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rank := []string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	Suit := []string{"Club", "Diamond", "Heart", "Spades"}
	total := make(map[string]int32)
	start := time.Now()
	var res string = ""
	rndRank, rndSuit := 0, 0
	count := 0
	for time.Since(start).Seconds() < 10 {
		if count%8 == 0 {
			print("\n")
		}
		s1 := rand.NewSource(time.Now().UnixNano())
		random := rand.New(s1)
		rndRank = random.Intn(13)
		rndSuit = rand.Intn(4)
		fmt.Printf("%s , %s  - ", rank[rndRank], Suit[rndSuit])
		res = rank[rndRank] + " , " + Suit[rndSuit]
		total[res]++
		time.Sleep(5 * time.Millisecond)
		count++
	}

	println("===========================================")
	count = 0
	var sum_ int32

	for v := range total {
		if count%6 == 0 {
			print("\n")
		}
		fmt.Printf("%v = %d | ", v, total[v])
		sum_ += total[v]
		count++
	}
	fmt.Printf("\n %d, %d ", count, sum_)
}
