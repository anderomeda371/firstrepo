package main

import "fmt"

func main() {

	var bmi float64
	var weight float64
	var height float64
	fmt.Print("please type your Weight(KG)=")
	fmt.Scan(&weight)
	//fmt.Printf("%.2f", float32(weight))
	fmt.Print("Please type your Height(Cm)=")
	fmt.Scan(&height)
	height = height / 100
	bmi = weight / (height * height)
	//bmi =
	fmt.Printf("BMI is = %f", bmi)
	if bmi > 23 {
		fmt.Printf("your BMI is %f.\n you are %s", bmi, "OBESE")

	} else if bmi > 18 {
		fmt.Printf("your BMI is %f.\n you are %s", bmi, "Normal")
	} else {
		fmt.Printf("your BMI is %f.\n you are %s", bmi, "Thin")

	}

}
