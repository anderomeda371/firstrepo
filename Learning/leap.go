package main

import "fmt"

func main() {

	var year int32
	fmt.Println("type a year(exit=0)=")
	fmt.Scan(&year)

	for {
		if year == 0 {
			break
		}

		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {

			fmt.Printf("%d is Leap\n", year)

		} else {
			fmt.Printf("%d is Not Leap\n", year)
		}

		fmt.Println("type a year(exit=0)=")
		fmt.Scan(&year)
	}

}
