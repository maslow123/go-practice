package main 

import "fmt"

func main() {
	var data = map[string]interface{} {
		"name"		: "M.Fadhly NR",
		"grade"		: 2,
		"height"	: 175.5,
		"isMale"	: true,
		"hobbies"	: []string{"eating", "sleeping"},
	}

	fmt.Println(data["name"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))
	fmt.Println("======================")
	// jika misal kita tidak tau tipe data asli dari data yang tersimpan dalam interface, kita bisa menggunakan teknik dibawah ini
	for _, val := range data {
		switch val.(type) {
		case string :
			fmt.Println(val.(string))
		case int :
			fmt.Println(val.(int))
		case float64 :
			fmt.Println(val.(float64))
		case bool :
			fmt.Println(val.(bool))
		case []string :
			fmt.Println(val.([]string))
		default :
			fmt.Println(val.(int))
		}
	}

}