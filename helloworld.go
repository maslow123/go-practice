package main

import "fmt"

func main() {
	// fmt.Println("Hello","gan") 
	// fmt.Println("Hello gan")
	
	// deklarasi variable
	var firstName string = "M.Fadhly"
	// metode interface, tidak perlu menuliskan lagi var dan tipedatanya
	lastName := "NR"
	// jika ingin ada perubahan value dalam variable
	lastName = "Noor Rizqi"
	
	// deklarasi variable secara bersamaan
	alamat, umur := "Griya cipeucang indah","17 tahun"

	// deklarasi type interface secara bersamaan
	one, isWednesday, twoPointTwo, say := 1, true, 2.2, "Hello gan"

	fmt.Printf("Halo nama depan saya %s, dan nama akhir saya %s",firstName, lastName+" Salam kenal go !\n")
	fmt.Printf("Alamat saya %s, dan saya berumur %s", alamat, umur+"\n")
	fmt.Println("One",one,"isWednesday", isWednesday,"twoPointTwo",twoPointTwo,"Say",say+"\n\n")
	/*			Variable underscore			*/
	_ = "Belajar golang" // tidak perlu memakai := karna tidak digunakan
	_ = "Golang itu ga susah ko"
	name,_ := "M.Fadhly","NR"
	fmt.Println("Name ",name)
	/*			Deklarasi Variable menggunakan New			*/
	nama := new(string)
	// fmt.Println(tes) // 0xc00002e230
	fmt.Println(*nama)
	/*			Deklarasi Variable menggunakan Make			*/
	

}
