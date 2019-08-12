package main 
import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	type People = Person
	
	people := People{"Wick",21}	
	fmt.Println(Person(people))

	var p1 = Person{"Wick",21}
	fmt.Println(p1)

	var p2 = People{"Wick",21}
	fmt.Println(p2)
}