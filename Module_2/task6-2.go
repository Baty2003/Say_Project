
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)


var T1 = time.Now()
func logTime()  {
	T2 := time.Now()
	fmt.Println(T2.Sub(T1))
}

func task6_2() {
	defer logTime()

	fileOpen, err := os.Open("Module_2\\Data\\in.txt")
	if err != nil {
		panic(err)
	}
	defer fileOpen.Close()

	fileCreate, err := os.Create("Module_2\\Data\\out.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lineCount := 0
	defer func() {
		stat, err := fileCreate.Stat()
		if err != nil {
		}
		fmt.Printf("The file is %d bytes long\n", stat.Size())
		fmt.Printf("The number of lines in the file %d\n", lineCount)
		fileCreate.Close()
	}()

	str := bufio.NewScanner(fileOpen)
	for str.Scan() {
		lineCount++
		newStr := fmt.Sprintf("%d. %s", lineCount, str.Text())
		fmt.Fprintln(fileCreate, newStr)
	}
}
