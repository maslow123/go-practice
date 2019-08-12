package main
import "fmt"

type student struct {
	name string
	grade int
}

func main() {
	/*	### DEKLARASI VARIABLE OBJEK ### */
		//	1	//
	var s1 student
	s1.name = "M.Fadhly NR"
	s1.grade= 2
		//	2	//
	var s2 = student{"M.Rasya NF",12}
		//	3	//
	var s3 = student{name: "M.Hilman NM",grade: 3}

	fmt.Println("Name :", s1.name)
	fmt.Println("Grade:", s1.grade)

	
	fmt.Println("Name :", s2.name)
	fmt.Println("Grade:", s2.grade)

	
	fmt.Println("Name :", s3.name)
	fmt.Println("Grade:", s3.grade)
}