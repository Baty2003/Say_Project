package main

import (
	"fmt"
	"math"
	"strconv"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func Task2() {

	one := "104"
	two := 35
	i, _ := strconv.Atoi(one)
	j := strconv.Itoa(two)
	fmt.Println(i, j)

	var (
		speedy  AmericanVelocity = AmericanVelocity(math.Round(120.4 * 3.6))
		speedy1 EuropeanVelocity = EuropeanVelocity(math.Round(130 * 2.2))
	)
	fmt.Printf("%.2f\n%.2f", speedy, speedy1)

}
