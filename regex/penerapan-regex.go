package main

import "fmt"
import "regexp"

func main() {
	/*				METHOD Compile()			*/
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`) // semua string alphabet yang hurufnya kecil, ekspresi tersebut di compile oleh *regexp.Compile() lalu disimpan ke variable regex bertipe *regexp.*Regexp*

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2) // method findAllString() berfungsi untuk mencari semua string yang sesuai dengan ekspetasi regex, dengan nilai kembalian berupa array, dan jumlahnya maksimal 2 data saja
	fmt.Printf("%#v \n", res1) // ["banana", "burger"]

	var res2 = regex.FindAllString(text, -1) // kita mengatur jumlah -1 yang berarti mengembalikan semua data
	fmt.Printf("%#v \n", res2) // ["banana","burger","soup"]
	fmt.Println("================================================")

	
	/*				METHOD MatchString()			*/
	var isMatch = regex.MatchString(text) 
	fmt.Println(isMatch) // akan bernilai true karna string "banana burger soup" memenuhi pola regex *[a-z]*
	fmt.Println("=================================================")

	/*				METHOD FindString()			*/
	var str = regex.FindString(text)
	fmt.Println(str) // "banana" (hanya mengembalikan 1 nilai pertama saja)
	fmt.Println("=================================================")

	/*				METHOD FingStringIndex()	*/
	var idx = regex.FindStringIndex(text)
	fmt.Println(idx) //[0, 6] (hanya mengembalikan 1 nilai pertama saja, dan hanya mendapatkan indexnya)
	fmt.Println("=================================================")

	var strIdx = text[0:6] 
	fmt.Println(strIdx) // ["banana"] // menampilkan index 0-6 dari 1 nilai pertama hasil pengembalian nilai
	fmt.Println("=================================================")

	/*				METHOD FindAllString()		*/
	var str1 = regex.FindAllString(text, -1)
	fmt.Println(str1) // ["banana","burger","soup"]

	var str2 = regex.FindAllString(text, 2)
	fmt.Println(str2) // ["banana"]
	fmt.Println("=================================================")

	/*				METHOD ReplaceAllString()	*/
	var strRAS = regex.ReplaceAllString(text,"potato")
	fmt.Println(strRAS) // ["potato","potato","potato"] (mengganti semua string dengan string "potato")
	fmt.Println("=================================================")

	/*				METHOD ReplaceAllStringFunc() */
	var strRASF = regex.ReplaceAllStringFunc(text, func(each string) string{ // dapatkan terlebih dahulu array stringnya
		if each == "burger" { // cek jika ada menemukan substring burger
			return "potato" // maka akan diganti dengan "potato", dan string yang lainnya tidak ada direplace
		}
		return each
	})

	fmt.Println(strRASF)
	fmt.Println("=================================================")

	/*				METHOD Split() 				*/
	var strSplit = regex.Split(text, -1)
	fmt.Printf("%#v \n", strSplit) // [""," "," ",""]


	


}