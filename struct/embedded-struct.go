/*
	Embeded struct adalah mekanisme untuk menempelkan sebuah struct sebagai properti struct lain, dan Embeded
	struct memiliki sifat mutable(dapat diubah-ubah).
*/
package main
import "fmt"

type person struct {
	name string
	age int
}

type student struct {
	age int
	grade int
	person
}

func main() {
	var s1 = student{}
	s1.name = "Fadhly"
	s1.age = 17
	s1.grade= 2
	s1.person.age = 18

	fmt.Println("Name : ", s1.name)
	fmt.Println("Age : ", s1.age) // age of student
	fmt.Println("Age : ", s1.person.age) // age of person
	fmt.Println("Grade: ", s1.grade)

	// pengisian nilai sub struct
	var p1 = person{name:"Marsha", age:18}
	var student1 = student{person:p1, grade:2}
	fmt.Println("=======================")
	
	fmt.Println("Name : ", student1.name)
	fmt.Println("Age : ", student1.age)
	fmt.Println("Grade : ", student1.grade)

}