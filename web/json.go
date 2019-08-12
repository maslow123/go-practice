package main 

import(
	"encoding/json"
	"fmt"
)

// buat struct user yang nantinya digunakan untuk membuat variable baru penampung hasil decode json string. Proses decode nantinya akan dilakukan pada *json.Unmarshal*
type User struct {
	FullName	string	`json:"Name"`
						// ^ merupakan tag *json: "name", yang nantinya digunakan untuk mapping informasi json ke property yang bersangkutan
	Age			int
}

func main() {
	var jsonString = `{"Name" : "M Fadhly", "Age" : 17}` // property name akan ditampung pada struct FullName
	var jsonData = []byte(jsonString) // casting json menjadi type byte

	var data User

	var err = json.Unmarshal(jsonData, &data) // unmarshal hanya menerima data json dalam bentuk byte, dan param kedua harus diisi dengan pointer dari objek yang nantinya akan menampung hasilnya
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)
	fmt.Println("User:", data.FullName)
	fmt.Println("Age:", data.Age)
	fmt.Println("=========================")

	/* Decode JSON ke *map[string]interface{}* dan *interface{}* */
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("User data1:",data1["Name"])
	fmt.Println("Age data1:",data1["Age"])
	fmt.Println("=========================")

	// Variable bertipe *interface{}* juga dapat menampung hasil decode, dengan catatan nilai property harus di casting terlebih dahulu ke map[string]interface{}
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("User data2:", decodedData["Name"])
	fmt.Println("Age data2:", decodedData["Age"])
	fmt.Println("=========================")

	/*****			Decode Array JSON ke Array Object			*****/
	
	//siapkan saja variable penampung hasil decode dengan tipe slice struct
	var stringJson = `
	[
		{"Name": "M.Fadhly NR", "Age": 17},
		{"Name": "M.Rasya NF", "Age": 12}
	]`

	var datas []User

	var errs = json.Unmarshal([]byte(stringJson), &datas)
	if err != nil {
		fmt.Println(errs.Error())
		return
	}
	fmt.Println(datas) // [{M.Fadhly NR 17} {M.Rasya NF 12}]
	fmt.Println("User 1:", datas[0].FullName)
	fmt.Println("User 2:", datas[1].FullName)
	fmt.Println("=========================")

	/*****			Encode Objek ke JSON String			*****/

	// data slice struct dikonversi kedalam bentuk json string. Hasil konversi berupa *[]byte*, casting terlebih dahulu JSON-nya ke string agar bisa ditampilkan bentuk json stringya
	var object = []User{{"M.Fadhly NR", 17}, {"M.Rasya NF", 12}}
	var dataJson, errData = json.Marshal(object) // casting bentuk byte
	if errData != nil {
		fmt.Println(errData.Error())
		return
	}

	var jsonString1 = string(dataJson) //casting ke string
	fmt.Println(jsonString1) // [{"Name":"M.Fadhly NR", "Age":17}, {"Name":"M.Rasya NF","Age":12}]
}