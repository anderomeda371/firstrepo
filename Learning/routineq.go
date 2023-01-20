package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter = 0
var waitG sync.WaitGroup
var m sync.Mutex

func main() {

	waitG.Add(2)
	go number(1)
	go number(2)
	waitG.Wait()
	fmt.Printf("\n %d ", counter)
}

func number(num int) {

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		//tmpCounter := counter
		//tmpCounter++
		time.Sleep(200 * time.Millisecond)
		m.Lock()
		counter++
		fmt.Printf("\n(%d)  %d,  %d", num, rand.Intn(20)+20, counter)
		m.Unlock()

		//counter = tmpCounter

	}
	waitG.Done()
}
