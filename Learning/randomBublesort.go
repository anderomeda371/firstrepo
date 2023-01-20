package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ar := []int{}
	i := 0
	for i < 10 {
		s1 := rand.NewSource(time.Now().UnixNano())
		time.Sleep(time.Millisecond)
		random := rand.New(s1)
		ar = append(ar, random.Intn(100))
		i++
	}
	fmt.Println(ar)

	j := 0
	count := 0
	tmpar := ar
	for i = 0; i < len(tmpar); i++ {

		for j = 0; j < len(tmpar)-1-i; j++ {
			count++
			if tmpar[j] > tmpar[j+1] {
				tmpar[j], tmpar[j+1] = tmpar[j+1], tmpar[j]
			}
		}
	}

	fmt.Println(tmpar)
	fmt.Println(count)
	count = 0
	for i = len(ar) - 1; i >= 0; i-- {

		for j = len(ar) - 1; j > 0+len(ar)-i-1; j-- {
			count++
			if ar[j] > ar[j-1] {
				ar[j], ar[j-1] = ar[j-1], ar[j]
			}
		}
	}

	fmt.Println(ar)
	fmt.Println(count)

}
