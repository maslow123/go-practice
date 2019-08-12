package main 	

import(
	"errors"
	"fmt"
	"strings"
)

func catch() {
	r:= recover();
	if  r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validate(input string)(bool, error) {
	if strings.TrimSpace(input) == "" { // untuk menghilangkan spasi sebelum dan sesudah string
		return false, errors.New("Cant be empty")
	}
	return true, nil
}

func main() {
	defer catch()

	var name string
	fmt.Print("Type your name : ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else{
		panic(err.Error())
		fmt.Println("end")
	}
}