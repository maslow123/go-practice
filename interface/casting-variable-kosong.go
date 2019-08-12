package main 

import "fmt"
import "strings"

func main() {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10// casting interface (int)
	fmt.Println(secret, "multiplied by 10 is:", number)

	secret = []string{"Banana","Apple","Manggo"}
	var fruits = strings.Join(secret.([]string),", ") // casting interface ([]string])
	fmt.Println(fruits, "are my favorite fruits")
	fmt.Println("===============================")

	/* Casting variable interface kosong ke objek pointer */
// contoh penerapan interface untuk menampung data objek pointer

	type person struct {
		name string
		age int
	}
	
	var secrets interface{} = &person{name: "Fadhly", age: 17}
	var name = secrets.(*person).name
	fmt.Println(name)
	
}
