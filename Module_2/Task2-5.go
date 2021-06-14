package main

import "fmt"

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
	max = 0
	for _, v := range numbers{
		if v > max{
			max = v
		}
	}
	return
}

func Task2_5(){
	someSliceString := []string{"Book","Laptop","City","Lake","Work"}
	someString := "laptop"
	contains(someSliceString,someString)
	fmt.Println(getMax(1,2,883,44,5,6,7,100,999))
}
