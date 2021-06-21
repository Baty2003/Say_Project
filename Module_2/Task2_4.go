package main

import "fmt"

func task2_4() {
	var history = map[string]map[string][]string{}
	history["Sergey"] = make(map[string][]string)
	history["Sergey"]["Public"] = []string{"1C study","1C advanced","1C BUh"}
	history["Oleg"] = make(map[string][]string)
	history["Oleg"]["private"] = []string{"Warcraft"}
	history["Kostan"] = make(map[string][]string)
	history["Kostan"]["Public"] = []string{"war and world"}
	history["None"] = make(map[string][]string)

	var countBook, readers = 0, 0
	for key, value := range history {
		countBook = 0
		for _,value1 := range value{
			countBook += len(value1)
			readers ++
		}
		fmt.Println(key,":",countBook)
	}



	fmt.Println(fmt.Sprintf("Total readers with books: %d", readers ))

}
