package main 

import(
	"errors"
	"fmt"
	"strings"
)

func validate(input string)(bool, error) {
	if strings.TrimSpace(input) == "" { // untuk menghilangkan spasi sebelum dan sesudah string
		return false, errors.New("Cant be empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("Type your name : ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid { // mengembalikan 2 data, data pertama nilai bool yang menandakan inputan apakah valid atau tidak. Dan data kedua adalah pesan errornya(jika inputan tidak valid)
		fmt.Println("Halo",name)
	} else {
		// aktif ketika inputan user tidak valid
		panic(err.Error())
		fmt.Println("end")
	}
}