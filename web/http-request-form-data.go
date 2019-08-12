package main 

import(
	"bytes"
	"net/url"
	"fmt"
	"net/http"
	"encoding/json"
)

type student struct {
	ID		string
	Name	string
	Grade	int
}
// buat fungsi baru, isinya req ke [::1]:8080/user dengan data yang disisipkan adalah ID

var baseURL = "http://localhost:8080"

func fetchUser(ID string) (student, error) { // mengembalikan 2 nilai, dan param harus diisi string
	var err error
	var client = &http.Client{}
	var data student 

	var param = url.Values{} // menghasilkan objek yang nantinya digunakan sebagai form data request, pada objek ini perlu di set data apa saja yang ingin dikirimkan
	param.Set("id", ID) // objek yang ingin dikirimkan berupa ID
	var payload = bytes.NewBufferString(param.Encode()) // objek form data di-encode lalu diubah menjadi bentuk *bytes.Buffer()*

	request, err := http.NewRequest("POST", baseURL+"/user", payload) 
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded") // karena data yang dikirim di-encode, maka header perlu di set tipe konten requestnya menjadi form-urlencoded

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
	
	/*
		Response dari endpoint /user bukanlah slice, tetapi berupa objek. Maka pada saat decode perlu pastikan tipe variabel penampung hasil decode data response adalah *student* bukan *[]student*
	*/
}

func main() {
	var user1, err = fetchUser("E001")
	if err != nil {
		fmt.Println("Error !", err.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\t", user1.ID, user1.Name, user1.Grade)
}
