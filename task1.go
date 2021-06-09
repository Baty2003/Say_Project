package main

import (
	"fmt"
	"time"
)

func Task1() {
	tTime := time.Date(2021, 06, 7, 20, 57, 2, 0, time.Local)
	fmt.Println(tTime.Format("02.01.2006 15:04"))

}
