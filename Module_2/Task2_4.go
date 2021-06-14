package main

import "fmt"

func task2_4() {
	var history = map[string]map[string][]string{}
	history["Sergey"] = make(map[string][]string)
	history["Sergey"]["Public"] = []string{"1C study"}
	history["Oleg"] = make(map[string][]string)
	history["Oleg"]["private"] = []string{"Warcraft"}
	history["Kostan"] = make(map[string][]string)
	history["Kostan"]["Public"] = []string{"war and world"}
	history["None"] = make(map[string][]string)

	var counter, readers = 0, 0
	for i, _ := range history {
		if len(history[i]) != 0 {
			counter++
		}
		readers++
	}
	fmt.Println(fmt.Sprintf("Total readers: %d \nTotal readers who have publications: %d ", readers, counter))

}
