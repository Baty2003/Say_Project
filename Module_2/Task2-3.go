package Module_2

import "fmt"

func Task4() {
	var s []string
	s = []string{
		"monday","Tuesday","Sreda","Chetverg","Friday","Saturday","Sunday",
	}
	b := make([]string, 10,10)
	b = s[:5]
	s = s[5:]
	fmt.Println(b)
	fmt.Println(b,s)
	b = append(b, s[:]...)
	fmt.Println(b)



}