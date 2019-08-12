package main 

import "fmt"

func main() {
	var person = []map[string]interface{} {
		{"name" : "Fadhly", "Age" : 17},
		{"name" : "Noor", "Age" : 17},
		{"name" : "Rizqi", "Age" : 17},
	}

	for _,each := range person {
		fmt.Println(each["name"],"age is", each["Age"])
	}
	fmt.Println("====================================")
	/* 
		Manfaatkan slice dan interface{}, agar dapat membuat data array yang isinya adalah bisa apa aja
	*/
	var fruits = []interface{} {
		map[string]interface{}{"name":"strawberry","total":10},
		[]string{"manggo","pineapple","papaya"},
		"orange",
	}
	
	for _,each := range fruits {
		fmt.Println(each)
	}
}