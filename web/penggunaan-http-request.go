package main 

import(
	"bytes"
	"net/url"
	"fmt"
	"net/http"
	"encoding/json"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID		string
	Name	string
	Grade	int
}

/*			HTTP REQUEST			*/
// buat fungsi *fetchUsers()*, yang mana bertugas untuk melakukan request, menerima response dari req tersebut, dan menampilkannya
func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{} // menghasilkan instance *http.Client* objek ini nantinya digunakan untuk req
	var data []student

	request, err := http.NewRequest("POST", baseURL+"/users", nil) // digunakan untuk membuat request baru
								// (tipe request, url tujuan request, form data request(jika ada))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	
	if err != nil {
		return nil, err
	}
	defer response.Body.Close() // menghasilkan data respon

	err = json.NewDecoder(response.Body).Decode(&data) // mengkonversi response body yang tadinya berbentuk string menjadi json
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*			HTTP REQUEST FORM DATA			*/
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
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}
	
	fmt.Println("======================================================================")

	var user1, errs = fetchUser("E001")
	if errs != nil {
		fmt.Println("Error user not found !", "")
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\t", user1.ID, user1.Name, user1.Grade)
}