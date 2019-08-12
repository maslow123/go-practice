/*
	Method pointer adalah method yang variable objek pemilik method tersebut berupa pointer
*/
package main 
import "fmt"

type student struct {
	name string
	grade int
}

func(s student) changeName1(name string){
	fmt.Println("----> on changeName1, name changed to", name)
}

func(s *student) changeName2(name string){
	fmt.Println("----> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	var s1 = student{"M.Fadhly NR",17}
	fmt.Println("s1 before", s1.name)
	// M.Fadhly NR

	s1.changeName1("M.Rasya NF")
	fmt.Println("s1 after changeName1",s1.name)
	// M.Fadhly NR

	s1.changeName2("M.Hilman NM")
	fmt.Println("s1 after changeName2",s1.name)
	// M.Hilman NM

}

