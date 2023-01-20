package main

import (
	"fmt"
)

func main() {

	str := "ibabi"
	strAr := []int32{}
	//lenstr := len(str)
	//println(str[1])
	for _, v := range str {
		strAr = append(strAr, int32(v))

		//fmt.Print(i)

	}
	palindrome := true
	maxlen := len(strAr) - 1
	j := 0
	for j <= maxlen {
		//if strAr[i] != strAr[lenstr-i-1] {
		//	fmt.Printf("%s is not palindrom", str)
		//	palindrome = false
		//	break
		fmt.Printf("%d ,%d, %d \n", str[j], str[maxlen], maxlen)
		if str[j] != str[maxlen] {
			palindrome = false
			fmt.Printf("%s is not palindrom", str)
			break
		}
		j++
		maxlen--

	}
	if palindrome == true {
		fmt.Printf("%s is  palindrom", str)

	}
	//?fmt.Print(strAr)
}
