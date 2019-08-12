/*
	Anonymous Struct merupakan struct yang tidak bisa dideklarasikan di awal sebagai tipe data baru, melainkan
	langsung ketika pembuatan object. 
*/
package main
import "fmt"

type person struct {
	name string
	age int
}

func main() {
	// anonymous struct tanpa pengisian property
	var s1 = struct {
		person
		grade int
	}{}
	
	s1.person = person{"Fadhly",17}
	s1.grade = 2
	// anonymous struct dengan pengisian property
	var s2 = struct {
		person
		grade int
	}{
		person: person{"Fadhly",17},
		grade:2,
	}

	fmt.Println("Name : ", s1.person.name)
	fmt.Println("Age : ", s1.age)
	fmt.Println("Grade : ", s1.grade)
	/* ================================= */
	fmt.Println("Name : ", s2.person.name)
	fmt.Println("Age : ", s2.age)
	fmt.Println("Grade : ", s2.grade)
}

