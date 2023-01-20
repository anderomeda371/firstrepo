package main

import (
	"time"
)

func main() {

	start := time.Now()
	curser := []string{"\\", "|", "/", "-"}
	i := 0
	print("please wait a moment ... ")
	for time.Since(start).Seconds() < 5 {

		print(curser[i])
		time.Sleep(100 * time.Millisecond)
		print("\b")
		i = (i + 1) % 4
	}
	print("\b\b")
	print("\n EXIT")
}
