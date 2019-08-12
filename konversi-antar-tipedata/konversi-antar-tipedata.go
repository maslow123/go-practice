package main 

import "fmt"
import "strconv"

func main() {
	var str = "123"
	var num, err = strconv.Atoi(str) // Atoi mengembalikan 2 nilai, jika nilainya true maka error bernilai *nil*

	if err == nil {
		fmt.Println(num,"\n") // 124
	} 
	
	// strconv.Itoa() == kebalikan dari atoi
	var numItoa = 124
	var strItoa = strconv.Itoa(numItoa)
	fmt.Println(strItoa,"\n") // "124"
	
	// strconv.ParseInt()
	var strPI = "124"
	var numPI, errPI = strconv.ParseInt(strPI, 10, 10)// (formatString, eksponen(hexa/oktal/binner), lebardata(int10))

	if errPI == nil {
		fmt.Println(numPI,"\n") // 124
	}

	// strconv.FormatInt()
	var numFI = int64(24)
	var strFI = strconv.FormatInt(numFI, 8)
	fmt.Println(strFI,"\n") // 30

	// strconv.ParseFloat()
	var strPF = "24.12"
	var numPF, errPF = strconv.ParseFloat(strPF, 32) // dikonversi ke float32

	if errPF == nil {
		fmt.Println(numPF,"\n") // hasil konversi sesuai dengan	IEEE standard for floating-point arithmetic
	}

	// strconv.FormatFloat()
	var numFF =  float64(24.12)
	// dikonversi dengan format eksponen *f* / tanpa eksponen, lebar digit desimal 6, dan lebar tipedata 64
	var strFF = strconv.FormatFloat(numFF, 'f', 3, 64) 
	fmt.Println(strFF,"\n")

	// strconv.ParseBool()
	var strPB = "true"
	var bul, errPB = strconv.ParseBool(strPB)

	if errPB == nil {
		fmt.Println(bul,"\n")
	}

	// strconv.FormatBool()
	var bulFB = true
	var strFB = strconv.FormatBool(bulFB)

	fmt.Println(strFB,"\n")
}