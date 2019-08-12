package main

import "fmt"

func main() {
	var names[4] string
	names[0] = "Budi"
	names[1] = "Rani"
	names[2] = "Sadam"
	names[3] = "Mina"

	fmt.Println("Teman : ",names[0], names[1],names[2],names[3])
	fmt.Println("==============================================")

	/* 			Inisialisasi Nilai Awal array 			*/
		// cara horizontal
	var fruits [4]string
	// fruits = [4]string { "apple", "grape", "banana", "melon"}
		// cara vertikal
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t\t",len(fruits)) // menghitung banyaknya data yang ada pada array
	fmt.Println("Isi semua element \t",fruits)
	fmt.Println("================================")

	
	/*			Inisialisasi Nilai awal array tanpa jumlah elemen			*/
	var numbers = [...]int{2,3,4,5,6}

	fmt.Println("Data array \t:",numbers)
	fmt.Println("Jumlah elemen \t:",len(numbers))
	fmt.Println("================================")

	/*			Array Multidimensi			*/
	var numbers1 = [ 2 ] [ 3 ] int { [ 3 ] int{ 3, 2, 3 }, [ 3 ] int{ 3, 4, 5 }}
	var numbers2 = [ 2 ] [ 3 ] int { { 3, 2, 3 }, { 3, 4, 5 }} // [2][3] 2 array 3 values

	fmt.Println(numbers1)
	fmt.Println(numbers2)
	fmt.Println("================================")

	/*			For and Array			*/
	var fruit = [4]string  { "apple", "grape", "banana", "melon"}

	fmt.Printf("Didalam fruit ada : ")
	for i:=0; i < len(fruit); i++ {
		fmt.Printf("%s ",fruits[i])
	}
	fmt.Println("\n================================")

	/*			Looping array with for-range			*/
	for i, fruitss := range fruit {
		fmt.Printf("Elemen %d : %s\n",i, fruitss)
	}
	fmt.Println("\n================================")

	/*			Variable underscore (_) Dalam for-range			*/
		// fungsinya untuk menampung nilai variable yang tidak digunakan

	for _, fruitss := range fruit { //kita tampung pengulangannya diletakkan pada variable (_)
		fmt.Printf("Nama buah : %s\n",fruitss)
	}
	fmt.Println("\n================================")
		// sekarang kita ambil indexnya saja

	// for i,_ := range fruit {
	for i := range fruit {
		fmt.Printf("Ini index ke - [%d]\n",i)
	}
	fmt.Println("\n================================")

	/*			Alokasi	elemen array menggunakan keyword make			*/
		// gunakan 2 parameter, parameter pertama tipedata elemen array yang diinginkan, dan yang kedua jumlah elemen yang diinginkan
	var fruitt = make( [] string, 2)
	fruitt[0]	= "apple"
	fruitt[1]	= "manggo"

	fmt.Println(fruitt)
	fmt.Println("\n================================")

	/*=====================  SLICE  ===========================*/
	var fruitsA = [] string { "Apple", "Grape" } // slice
	var fruitsB = [2] string { "Banana", "Melon"} // array
	var fruitsC = [...] string { "Papaya", "Grape"} // array
	fmt.Printf("Ini slice %s\n",fruitsA)
	fmt.Printf("Ini array %s\n",fruitsB)
	fmt.Printf("Ini array juga %s",fruitsC)
	fmt.Println("\n================================")

	var sliceFruit = []string {"Apple","Grape","Banana","Melon"}
	var newFruit = sliceFruit[ 0:2 ] // kita ambil array mulai dari index ke-0 sampai sebelum index ke-2 
	fmt.Println(newFruit)
	fmt.Println("\n================================")

	/*===================== SLICE IS REFERENCE ======================*/
	var fruitz = [] string { "apple", "grape", "melon", "manggo" }

	var aFruitz = fruitz[ 0 : 3 ] // [ apple, grape, melon ]
	var bFruitz = fruitz[ 1 : 4 ] // [ grape, melon, manggo ]

	// get reference from old slice
	var aaFruitz = aFruitz[ 1 : 2 ] // [ grape ]
	var bbFruitz = bFruitz[ 0 : 1 ] // [ grape ]

	fmt.Println(fruitz)
	fmt.Println(aFruitz)
	fmt.Println(bFruitz)
	fmt.Println(aaFruitz)
	fmt.Println(bbFruitz)
	fmt.Println("================================")
	// change the element array
	
	bbFruitz[0] = "pineapple"

	fmt.Println(fruitz)
	fmt.Println(aFruitz)
	fmt.Println(bFruitz)
	fmt.Println(aaFruitz)
	fmt.Println(bbFruitz)
	fmt.Println("================================")
	
	/*===================== ARRAY CAP ======================*/
	fmt.Println(len(bFruitz)) // len : 3
	fmt.Println(cap(bFruitz)) // cap : 3
	// fmt.Println(cap(aFruitz)) // cap : 4
	fmt.Println("================================")
	/*===================== ARRAY APPEND ======================*/
	var cFruitzz = append(fruitz,"papaya")
	fmt.Println(cFruitzz)
	fmt.Println("================================")
	/*===================== ARRAY COPY ======================*/
		//fungsi copy() yaitu mengembalikan informasi angka, representasi dari jumlah element yang berhasil di copy
	dst := make([]string, 3)
	src := [] string {"watermelon","pineapple","apple","orange"}
	n := copy(dst,src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
	fmt.Println("================================")
	
	dst = []string {"potato","potato","potato","sambalado","ya"}
	src = []string {"watermelon","pineapple"}
	n = copy(dst,src)

	fmt.Println(dst) // [ watermelon, pineapple, potato, sambalado, ya]
	fmt.Println(src) // 
	fmt.Println(n) // 2 {karna hanya yang dicopy 2}
	fmt.Println("================================")

	/*===================== PENGAKSESAN ELEMEN SLICE 3 INDEKS ======================*/
	var aFruits = fruitz[0 : 2]
	// 3 index [mulainya, berenti sebelum index x, kapasitasnya]
	var bFruits = fruitz[0 : 2 : 4]

	fmt.Println(fruitz)
	fmt.Println(len(fruitz)) // len 3
	fmt.Println(cap(fruitz)) // cap 3

	fmt.Println(aFruits)// ["apple","grape"]
	fmt.Println(len(aFruits)) // len 2
	fmt.Println(cap(aFruits)) // cap 3

	fmt.Println(bFruits) // ["apple","grape"]
	fmt.Println(len(bFruits)) // len 2
	fmt.Println(cap(bFruits)) // cap 2
	fmt.Println("================================")

	/*===================== ARRAY MAP ======================*/
	var chicken map[string] int
	chicken = map[string] int {}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("Januari",chicken["januari"])
	fmt.Println("Mei",chicken["mei"]) // 0 karna belum ada item yang tersimpan menggunakan key mei, dan 0 merupakan nilai default dari int
	fmt.Println("============================")
	// definisi map diawal
	var chicken2 = map[string] int {
		"januari": 50,
		"februari": 40,
	}

	fmt.Println(chicken2["januari"])
	fmt.Println("================================")

	/*===================== Iterasi item Map dengan menggunakan for - range ======================*/
	var chicks = map[string]int{
		"januari": 50,
		"februari": 40,
		"maret":34,
		"april": 67,
	}
	for key,val := range chicks {
		fmt.Println(key, "\t:",val)
	}	
	fmt.Println("================================")

	/*===================== Menghapus Item Map ======================*/
		// delete() = menghapus item dengan key tertentu pada variable map
	delete(chicks,"januari")
	fmt.Println(len(chicks)) // 3
	fmt.Println(chicks)
	fmt.Println("================================")

	/*===================== Deteksi keberadaan item dengan key tertentu ======================*/
		// manfaatkan 2 variable sebagai penampung nilai kembalian pengaksesan item.
	var value,isExist = chicks["februari"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Data no exists")
	}
	fmt.Println("================================")

	/*===================== Kombinasi slice and map ======================*/
	var chick = []map[string]string{
		{ "name":"chicken blue", "gender": "male" },
		{ "name":"chicken red", "gender":"male"},
		{ "name":"chicken yellow","gender":"female"},
	}

	for _, chickk := range chick {
		fmt.Println(chickk["gender"], chickk["name"])
	}
	fmt.Println("================================")

	/* 
		file:///C:/Users/Maslow/Downloads/dasarpemrogramangolang.pdf (hal 52)
	*/

}
