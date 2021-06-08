package main

import (
	"fmt"
	"math"
	"time"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func (s AmericanVelocity) convert() AmericanVelocity {
	return s * 3.6
}
func (s EuropeanVelocity) convert() EuropeanVelocity {
	return s * 2.22
}

func Task1() {
	tTime := time.Date(2021, 06, 7, 20, 57, 2, 0, time.Local)
	fmt.Println(tTime.Format("02.01.2006 15:04"))
	var (
		speedy  AmericanVelocity = 120.4
		speedy1 EuropeanVelocity = 130
	)
	fmt.Println(math.Round(float64(speedy.convert())*100) / 100)
	fmt.Println(math.Round(float64(speedy1.convert())*100) / 100)
}
