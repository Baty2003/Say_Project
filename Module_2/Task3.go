package Module_2

import (
	"fmt"
	"math"
)

func Task3()  {
	//var a *int
	//b := 5
	//a = &b
	//fmt.Println(*a)
	//b = 10
	//fmt.Println(*a)
	radius := new(float64)
	*radius = math.Round((35 /(2 * math.Pi))*100)/100
	s := math.Round((math.Pi * math.Pow(*radius,2))*100)/100

	fmt.Println(*radius)
	fmt.Println(s)

}
