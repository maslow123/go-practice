package main

import "fmt"
import "strings"

func main() {
	var isExists = strings.Contains("M Fadhly","Fadhly") // cek apakah string merupakan bagian string lain(parameter pertama)
	fmt.Println(isExists)
	fmt.Println("====================================")

	// strings.HasPrefix()
	var isPrefix1 = strings.HasPrefix("M Fadhly","M") // cek string dari parameter pertama diawali huruf M atau tidak
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("M Fadhly","Fa")
	fmt.Println(isPrefix2) // false
	fmt.Println("====================================")

	// strings.HasSuffix()
	var isSufix1 = strings.HasSuffix("M Fadhly","ly") // cek string dari paramter pertama diakhiri huruf ly atau tidak
	fmt.Println(isSufix1)

	var isSufix2 = strings.HasSuffix("M Fadhly","lu")
	fmt.Println(isSufix2)
	fmt.Println("====================================")

	// strings.Count()
	var howMany = strings.Count("How many string from this sentence ?", "s")
	fmt.Println(howMany)
	fmt.Println("====================================")

	// strings.Index()
	var index1 = strings.Index("M Fadhly","y")
	fmt.Println(index1) // 7 {karena huruf Y berada pada index ke 7 *spasi tidak dihitung}

	var index2 = strings.Index("M Fadhly","F")
	fmt.Println(index2) // 2
	fmt.Println("====================================")

	// strings.Replace()
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1) // kita hanya mengganti 1 huruf "a" menjadi huruf "o"
	fmt.Println(newText1)

	var newText2 = strings.Replace(text, find, replaceWith, 2) // kita hanya mengganti 2 huruf "a" menjadi huruf "o"
	fmt.Println(newText2)

	var newText3 = strings.Replace(text, find, replaceWith, 3) // kita hanya mengganti 3 huruf "a" menjadi huruf "o"
	fmt.Println(newText3)
	fmt.Println("====================================")

	// strings.Repeat()
	var str = strings.Repeat("na", 4)
	fmt.Println(str)
	fmt.Println("====================================")

	// strings.Split()
	var string1 = strings.Split("The dark night"," ") 
	fmt.Println(string1)// [The dark night]

	var string2 = strings.Split("Fadhly","")
	fmt.Println(string2) // [F a d h l y]
	fmt.Println("====================================")

	// strings.Join()
	var data = []string{"banana","pisang","cau"}
	var strr = strings.Join(data,"-") // digabungkan dan dipisahkan oleh (dash)
	fmt.Println(strr) // banana-pisang-cau
	fmt.Println("====================================")

	// strings.ToUpper()
	var strUpper = strings.ToUpper("eat !")
	fmt.Println(strUpper)
	fmt.Println("====================================")
	
	// strings.ToLower()
	var strLower = strings.ToLower("EAT !")
	fmt.Println(strLower)
	fmt.Println("====================================")
}