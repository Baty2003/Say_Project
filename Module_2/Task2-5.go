package main

import (
	"fmt"
	"math"
)

func contains(a []string ,x string)  {
	for _,v := range a{
		if v == x{
			fmt.Println("Yes, string is find!")
		}
	}
	fmt.Println("Sorry, we really tried, but didn't find ")
	return
}

func getMax(numbers... int)  (max int){
	max = math.MinInt64
	for _, v := range numbers{
		if v > max{
			max = v
		}
	}
	return
}

func Task2_5(){
	someSliceString := []string{"Book","Laptop","City","Lake","Work","College"}
	someString := "laptop"
	contains(someSliceString,someString)
	fmt.Println(getMax(-2,-4))
}
