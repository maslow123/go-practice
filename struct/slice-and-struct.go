package main 
import "fmt"

type persons struct {
	name string
	age int
}

func main() {
	var allStudents = []persons {
		{name: "Fadhly", age: 17},
		{name: "Noor", age: 17},
		{name: "Rizqi", age: 17},
	}
	for _, student := range allStudents {
		fmt.Println(student.name,"age is",student.age)
	}
	// initialize slice anonymous struct
	var allStudent = []struct {
		persons
		grade int
	}{
		{persons: persons{"Rasya",14},grade: 2},
		{persons: persons{"Noor",14},grade: 2},
		{persons: persons{"Fajri",14},grade: 2},
	}
	fmt.Println("=====================")
	for _, student := range allStudent {
		fmt.Println(student.persons.name)
	}
	fmt.Println("=====================")
	// deklarate anonymous struct with var
	var studentt struct {
		persons
	}
	studentt.persons = persons{"Hilman",23}
	fmt.Println(studentt.persons.name)
	// deklarate and initialize
	var studentss = struct {
		grade int
	}{
		12,
	}
	fmt.Println(studentss.grade)
}