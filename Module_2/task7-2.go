
package main

import (
	"bufio"
	"fmt"
	"os"
)



func task7_2() {


	fileOpen, err := os.Open("Module_2\\Data\\in2.txt")
	if err != nil {
		panic(err)
	}

	fileCreate, err := os.Create("Module_2\\Data\\out2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lineCount := 0
	defer func() {
		fileOpen.Close()
		fileCreate.Close()
		if err := recover(); err != nil{
			fileCreated, err := os.Open("Module_2\\Data\\out2.txt")
			fmt.Println(err)
			strAnother := bufio.NewScanner(fileCreated)
			for strAnother.Scan() {
				fmt.Println(strAnother.Text(),"")
			}
			fileCreated.Close()
		}
	}()

	str := bufio.NewScanner(fileOpen)
	var Name,Address,City string
	var switchVariable int

	for str.Scan() {

		lineCount++
		switchVariable = 0

		for _,j := range str.Text(){

			if string(j) == "|"{
				switchVariable += 1
				continue
				}

			switch switchVariable {
				case 0:
					Name += string(j)
				case 1:
					Address += string(j)
				case 2:
					City += string(j)
				}
			}

		if Name == ""{
			panic(fmt.Sprintf("parse error: empty field 'Name' on string %d",lineCount))
		}else if Address == ""{
			panic(fmt.Sprintf("parse error: empty field 'Address' on string %d",lineCount))
		}else if City == ""{
			panic(fmt.Sprintf("parse error: empty field 'City' on string %d",lineCount))
		}

		strOut := fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n",lineCount,Name,Address,City)
		fmt.Fprint(fileCreate, strOut)
		Name,Address,City = "","",""
		}
		fmt.Println("Completed successfully!")
	}

